#!/usr/bin/env bash
# Injects failures into random NeuroOps pods (by calling /fail inside the pod)
# and measures, for each one, how long Kubernetes takes to restart the
# container and bring it back to Ready.
#
#   scripts/inject-failures.sh [COUNT]          default COUNT=10
#   INTERVAL=30 NAMESPACE=neuroops OUTPUT=loadtest/results/failures.csv scripts/inject-failures.sh 20
#
# Run loadtest/traffic.js at the same time to measure availability during the
# failures. Use hpa.minReplicas >= 2, otherwise every restart is downtime.
set -euo pipefail

COUNT=${1:-10}
NAMESPACE=${NAMESPACE:-neuroops}
SELECTOR=${SELECTOR:-app=neuroops-api}
INTERVAL=${INTERVAL:-30}
TIMEOUT=${TIMEOUT:-120}
OUTPUT=${OUTPUT:-}

now() { perl -MTime::HiRes=time -e 'printf "%.3f\n", time'; }

if [[ -n "$OUTPUT" ]]; then
  mkdir -p "$(dirname "$OUTPUT")"
  echo "injection,pod,injected_at,restart_seconds,ready_seconds" > "$OUTPUT"
fi

results=()

for i in $(seq 1 "$COUNT"); do
  pods=($(kubectl get pods -n "$NAMESPACE" -l "$SELECTOR" --field-selector=status.phase=Running \
    -o jsonpath='{.items[*].metadata.name}'))
  if [[ ${#pods[@]} -eq 0 ]]; then
    echo "no running pods match $SELECTOR in $NAMESPACE" >&2
    exit 1
  fi
  pod=${pods[RANDOM % ${#pods[@]}]}

  before=$(kubectl get pod "$pod" -n "$NAMESPACE" -o jsonpath='{.status.containerStatuses[0].restartCount}')
  injected_at=$(date -u +%Y-%m-%dT%H:%M:%SZ)
  start=$(now)
  kubectl exec -n "$NAMESPACE" "$pod" -- wget -q -O /dev/null http://localhost:8080/fail 2>/dev/null || true

  restart_s=""
  ready_s=""
  while :; do
    elapsed=$(echo "$(now) - $start" | bc)
    if (( $(echo "$elapsed > $TIMEOUT" | bc) )); then
      echo "[$i/$COUNT] $pod did not recover within ${TIMEOUT}s" >&2
      break
    fi

    read -r restarts ready < <(kubectl get pod "$pod" -n "$NAMESPACE" \
      -o jsonpath='{.status.containerStatuses[0].restartCount} {.status.containerStatuses[0].ready}')

    if [[ -z "$restart_s" && "$restarts" -gt "$before" ]]; then
      restart_s=$elapsed
    fi
    if [[ -n "$restart_s" && "$ready" == "true" ]]; then
      ready_s=$elapsed
      break
    fi
    sleep 0.5
  done

  printf "[%d/%d] %s  restarted after %ss, ready after %ss\n" "$i" "$COUNT" "$pod" "${restart_s:-n/a}" "${ready_s:-n/a}"
  [[ -n "$OUTPUT" ]] && echo "$i,$pod,$injected_at,$restart_s,$ready_s" >> "$OUTPUT"
  [[ -n "$ready_s" ]] && results+=("$ready_s")

  if [[ "$i" -lt "$COUNT" ]]; then
    sleep "$INTERVAL"
  fi
done

echo
echo "Recovered ${#results[@]}/$COUNT injected failures."
if [[ ${#results[@]} -gt 0 ]]; then
  printf "%s\n" "${results[@]}" | sort -n | awk '
    { v[NR] = $1; sum += $1 }
    END {
      p95 = v[int((NR - 1) * 0.95) + 1]
      printf "Time to Ready  avg %.1fs  min %.1fs  p95 %.1fs  max %.1fs\n", sum / NR, v[1], p95, v[NR]
    }'
fi

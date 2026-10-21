#!/usr/bin/env bash
# Measures how long a secret change in Vault takes to reach the Kubernetes
# Secret synced by the Vault Secrets Operator, with no redeploy or kubectl apply.
# It patches a rotated_at marker next to the real value, so nothing breaks.
#
#   ./vault/measure-rotation.sh
#   RUNS=5 ./vault/measure-rotation.sh
set -euo pipefail

cd "$(dirname "$0")"
NAMESPACE=${NAMESPACE:-neuroops}
SECRET=${SECRET:-slack-webhook}
VAULT_PATH=${VAULT_PATH:-kv/neuroops/slack}
RUNS=${RUNS:-1}
TIMEOUT=${TIMEOUT:-180}
INIT_FILE=.vault-init.json

vault() { kubectl exec -i -n vault vault-0 -- vault "$@"; }
now() { perl -MTime::HiRes=time -e 'printf "%.3f\n", time'; }

trap 'kubectl exec -n vault vault-0 -- sh -c "rm -f \$HOME/.vault-token" >/dev/null 2>&1 || true' EXIT
jq -r .root_token "$INIT_FILE" | vault login -no-print -

for run in $(seq 1 "$RUNS"); do
  marker="$(date -u +%Y-%m-%dT%H:%M:%SZ)-$RANDOM"
  start=$(now)
  vault kv patch "$VAULT_PATH" rotated_at="$marker" >/dev/null

  while :; do
    current=$(kubectl get secret "$SECRET" -n "$NAMESPACE" -o jsonpath='{.data.rotated_at}' | base64 --decode)
    elapsed=$(echo "$(now) - $start" | bc)
    if [[ "$current" == "$marker" ]]; then
      printf "[%d/%d] %s/%s updated %.1fs after the Vault write\n" "$run" "$RUNS" "$NAMESPACE" "$SECRET" "$elapsed"
      break
    fi
    if (( $(echo "$elapsed > $TIMEOUT" | bc) )); then
      echo "[$run/$RUNS] not propagated after ${TIMEOUT}s" >&2
      exit 1
    fi
    sleep 0.5
  done
done

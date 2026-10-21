#!/usr/bin/env bash
# Installs the NeuroOps observability stack into the "monitoring" namespace:
# Prometheus + Alertmanager + Grafana (metrics, alerts), Loki + Alloy (logs),
# Tempo (traces). Chart versions are pinned for reproducible installs.
set -euo pipefail

cd "$(dirname "$0")"
NAMESPACE=monitoring

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts >/dev/null
helm repo add grafana https://grafana.github.io/helm-charts >/dev/null
helm repo update >/dev/null

helm upgrade --install neuroops-monitoring prometheus-community/kube-prometheus-stack \
  --version 92.3.0 -n "$NAMESPACE" --create-namespace \
  -f kube-prometheus-stack.yaml --wait --timeout 10m

helm upgrade --install loki grafana/loki \
  --version 7.3.0 -n "$NAMESPACE" -f loki.yaml --wait --timeout 10m

helm upgrade --install tempo grafana/tempo \
  --version 1.24.4 -n "$NAMESPACE" -f tempo.yaml --wait --timeout 10m

helm upgrade --install alloy grafana/alloy \
  --version 1.13.1 -n "$NAMESPACE" -f alloy.yaml --wait --timeout 10m

echo
echo "Monitoring stack ready. Grafana:"
echo "  kubectl port-forward svc/neuroops-monitoring-grafana 3000:80 -n $NAMESPACE"

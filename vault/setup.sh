#!/usr/bin/env bash
# Installs Vault and the Vault Secrets Operator (VSO), then initializes, unseals,
# and configures Vault for NeuroOps:
#   - KV v2 secrets engine mounted at kv/
#   - Kubernetes auth: ServiceAccount neuroops/neuroops-vault can only read kv/neuroops/*
#   - audit log on stdout (shipped to Loki by Alloy)
# Safe to re-run, e.g. after a cluster restart: it only unseals an initialized Vault.
#
#   ./vault/setup.sh
#   SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXX ./vault/setup.sh
set -euo pipefail

cd "$(dirname "$0")"
NAMESPACE=vault
INIT_FILE=.vault-init.json

for bin in kubectl helm jq; do
  command -v "$bin" >/dev/null || { echo "missing required tool: $bin" >&2; exit 1; }
done

vault() { kubectl exec -i -n "$NAMESPACE" vault-0 -- vault "$@"; }
status() { vault status -format=json 2>/dev/null || true; }

helm repo add hashicorp https://helm.releases.hashicorp.com >/dev/null
helm repo update hashicorp >/dev/null

servicemonitor=false
if kubectl get crd servicemonitors.monitoring.coreos.com >/dev/null 2>&1; then
  servicemonitor=true
fi

helm upgrade --install vault hashicorp/vault --version 0.34.1 \
  -n "$NAMESPACE" --create-namespace -f values.yaml \
  --set serverTelemetry.serviceMonitor.enabled="$servicemonitor"

helm upgrade --install vault-secrets-operator hashicorp/vault-secrets-operator --version 1.6.0 \
  -n vault-secrets-operator-system --create-namespace --wait

echo "Waiting for vault-0..."
until kubectl get pod vault-0 -n "$NAMESPACE" >/dev/null 2>&1; do sleep 2; done
kubectl wait --for=jsonpath='{.status.phase}'=Running pod/vault-0 -n "$NAMESPACE" --timeout=300s >/dev/null
until [[ -n "$(status)" ]]; do sleep 2; done

if [[ "$(status | jq -r .initialized)" != "true" ]]; then
  echo "Initializing Vault (1 unseal key: lab setup)"
  vault operator init -key-shares=1 -key-threshold=1 -format=json > "$INIT_FILE"
  chmod 600 "$INIT_FILE"
fi

if [[ ! -f "$INIT_FILE" ]]; then
  echo "Vault is already initialized but $INIT_FILE is missing; cannot unseal." >&2
  exit 1
fi

if [[ "$(status | jq -r .sealed)" == "true" ]]; then
  echo "Unsealing Vault"
  vault operator unseal "$(jq -r '.unseal_keys_b64[0]' "$INIT_FILE")" >/dev/null
fi

# The root token is only used for this configuration run.
trap 'kubectl exec -n "$NAMESPACE" vault-0 -- sh -c "rm -f \$HOME/.vault-token" >/dev/null 2>&1 || true' EXIT
jq -r .root_token "$INIT_FILE" | vault login -no-print -

vault secrets list -format=json | jq -e '."kv/"' >/dev/null \
  || vault secrets enable -path=kv kv-v2

vault audit list -format=json 2>/dev/null | jq -e '."stdout/"' >/dev/null \
  || vault audit enable -path=stdout file file_path=stdout

vault auth list -format=json | jq -e '."kubernetes/"' >/dev/null \
  || vault auth enable kubernetes

vault write auth/kubernetes/config kubernetes_host=https://kubernetes.default.svc:443 >/dev/null

vault policy write neuroops-read - >/dev/null <<'EOF'
path "kv/data/neuroops/*" {
  capabilities = ["read"]
}
EOF

vault write auth/kubernetes/role/neuroops \
  bound_service_account_names=neuroops-vault \
  bound_service_account_namespaces=neuroops \
  audience=vault \
  policies=neuroops-read \
  token_ttl=1h >/dev/null

if [[ -n "${SLACK_WEBHOOK_URL:-}" ]]; then
  printf %s "$SLACK_WEBHOOK_URL" \
    | vault kv put kv/neuroops/slack url=- rotated_at="$(date -u +%Y-%m-%dT%H:%M:%SZ)" >/dev/null
  echo "Stored the Slack webhook at kv/neuroops/slack"
fi

echo
echo "Vault is unsealed and configured."
echo "  UI:   kubectl port-forward svc/vault 8200:8200 -n $NAMESPACE  ->  http://localhost:8200"
echo "  Keys: vault/$INIT_FILE (git-ignored; keep it private)"

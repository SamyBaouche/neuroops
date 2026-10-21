# NeuroOps 🚀

<div align="center">

## AI-powered Kubernetes observability and remediation platform

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)](#-technology-stack)
[![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?logo=docker&logoColor=white)](#-technology-stack)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Orchestrated-326CE5?logo=kubernetes&logoColor=white)](#-technology-stack)
[![Terraform](https://img.shields.io/badge/Terraform-IaC-7B42BC?logo=terraform&logoColor=white)](#-technology-stack)
[![AWS](https://img.shields.io/badge/AWS-Cloud%20Native-FF9900?logo=amazonaws&logoColor=white)](#-technology-stack)
[![GitHub Actions](https://img.shields.io/badge/GitHub%20Actions-CI%2FCD-2088FF?logo=githubactions&logoColor=white)](#-cicd-pipeline)
[![Prometheus](https://img.shields.io/badge/Prometheus-Metrics-E6522C?logo=prometheus&logoColor=white)](#-prometheus--grafana-observability)
[![Grafana](https://img.shields.io/badge/Grafana-Dashboards-F46800?logo=grafana&logoColor=white)](#-prometheus--grafana-observability)
[![Helm](https://img.shields.io/badge/Helm-K8s%20Packages-0F1689?logo=helm&logoColor=white)](#-prometheus--grafana-observability)

</div>

---

## 🌍 Vision

NeuroOps started as a Kubernetes self-healing lab and evolved into a production-oriented, AI-native platform engineering project.

The mission is simple and ambitious:

- ⚡ detect operational risk early,
- 🧠 explain incidents with AI-assisted reasoning,
- 🛠️ recommend (and later automate) safe remediation,
- ☁️ run cloud-native workloads with real SRE patterns.

This repository showcases platform engineering maturity across reliability controls, autoscaling, deployment automation, observability workflows, and AI-driven operations direction.

---

## 🎯 Project Goals

- Build a resilient Kubernetes-ready service with health and readiness semantics.
- Operate workload behavior under stress using controlled failure and load simulation.
- Apply cloud-native deployment practices with Docker, Kubernetes, and Infrastructure as Code.
- Provide an operator-first CLI/TUI workflow for fast diagnosis from the terminal.
- Implement delivery automation with CI/CD and secure secret handling.
- Evolve toward AI incident analysis and AI remediation loops.

---

## 📅 Project Progress

### Day 1

- Go API, operational endpoints, and local runtime baseline.

### Day 2

- Docker packaging and image lifecycle fundamentals.

### Day 3

- Kubernetes deployment, service exposure, and recovery debugging.

### Day 4

- Probe-first resiliency patterns (`livenessProbe`, `readinessProbe`).

### Day 5

- HPA autoscaling and operator UX via Cobra CLI + Bubble Tea TUI.

### CI/CD Automation

- GitHub Actions CI/CD with test/build/publish automation.

### Prometheus + Grafana Observability

- Prometheus + Grafana monitoring stack with Helm on Minikube.

---

## 🧱 Core Capabilities

### ✅ Kubernetes Self-Healing

NeuroOps uses Kubernetes-native resilience patterns:

- `livenessProbe` checks if the container should be restarted.
- `readinessProbe` checks if the pod is ready to receive traffic.
- automatic pod restart protects service continuity.
- self-healing behavior improves production resilience.

Why it matters:

- Prevents traffic from reaching unhealthy pods.
- Enables autonomous restart for failed app states.
- Reduces MTTR through platform-level recovery.

How Kubernetes reacts:

- Failed liveness probe: kubelet restarts container.
- Failed readiness probe: pod is removed from service endpoints.
- Recovery: pod rejoins traffic only when healthy and ready.

### ✅ HPA Auto-Scaling

NeuroOps demonstrates dynamic scaling with `metrics-server` + HPA:

- monitors CPU utilization,
- scales replica counts based on workload pressure,
- supports elastic workloads.

Example scaling behavior:

- normal traffic: `1 pod`
- CPU spike: `1 -> 5 pods`
- traffic normalization: scale down automatically

Why scaling matters:

- absorbs burst traffic safely,
- improves availability under variable load,
- optimizes cost/performance balance.

### ✅ Cobra CLI

Cobra powers production-style terminal operations:

- command-driven infrastructure workflow,
- repeatable operator actions,
- clean and scriptable UX for platform tasks.

### ✅ Bubble Tea TUI

Bubble Tea provides an interactive terminal dashboard for operations:

- service health checks,
- readiness checks,
- CPU load simulation,
- failure simulation.

Why it matters:

- improves developer/operator feedback loop,
- enables fast diagnostics without leaving terminal context,
- offers a modern CLI experience for SRE workflows.

### ✅ AWS Integration

NeuroOps is cloud-deployment ready with AWS patterns:

- Amazon ECR for container image registry,
- IAM for identity and least-privilege access,
- VPC and Security Groups for network boundary control,
- KMS for encryption key management and secrets posture.

Why it matters:

- enterprise-grade deployment path,
- secure supply chain and runtime foundation,
- production-friendly governance model.

### ✅ Terraform Infrastructure as Code

Terraform enables:

- reproducible infrastructure,
- versioned provisioning,
- automated cloud resource lifecycle.

Why it matters:

- eliminates manual configuration drift,
- improves auditability,
- standardizes environment creation.

### ✅ Docker

Docker provides:

- consistent runtime packaging,
- local-to-cloud portability,
- immutable deployable artifacts.

Why it matters:

- environment parity across dev/test/prod,
- faster release reliability,
- simpler Kubernetes deployment flow.

### ✅ GitHub Actions CI/CD

NeuroOps CI/CD supports:

- automated build validation,
- Go test execution,
- container build and ECR push,
- repeatable delivery gates.

Why it matters:

- faster and safer releases,
- reduced manual deployment risk,
- measurable engineering velocity.

### ✅ Observability Stack

NeuroOps includes a production-style observability baseline:

- Prometheus for metrics collection,
- Grafana dashboards for visualization,
- Kubernetes namespace and pod-level monitoring,
- CPU and memory telemetry for scaling analysis.

---

## 🧠 Platform Architecture

```text
Developer -> Cobra CLI / Bubble Tea TUI
            |
            v
        NeuroOps API (Go)
            |
            +--> Health/Ready/Fail/Load endpoints
            |
            v
      Docker Image Build
            |
            v
  GitHub Actions CI/CD Pipeline
            |
            +--> Terraform (IaC) -> AWS Resources (ECR/IAM/Network/Security)
            |
            v
       Kubernetes Cluster
            |
            +--> Probes (self-healing)
            +--> HPA (auto-scaling)
            +--> Prometheus metrics
            +--> Grafana dashboards
            +--> Future AI analysis/remediation loop
```

Architecture principles:

- reliability-first workload design,
- platform automation over manual ops,
- observable systems with AI-driven next steps,
- safe remediation before full autonomy.

---

## 🧰 Technology Stack

| Technology          | Purpose                                                     | Why It Matters                                                          |
| ------------------- | ----------------------------------------------------------- | ----------------------------------------------------------------------- |
| Go                  | Backend API and operational endpoints                       | High performance, simplicity, strong cloud-native ecosystem             |
| Docker              | Build and package immutable artifacts                       | Reliable portability across local and cloud runtimes                    |
| Kubernetes          | Workload orchestration and self-healing                     | Production-grade scheduling, recovery, and service resilience           |
| Terraform           | Infrastructure as Code provisioning                         | Reproducible and auditable cloud automation                             |
| AWS                 | Cloud infrastructure foundation                             | Enterprise-ready networking, security, and managed services             |
| GitHub Actions      | CI/CD automation                                            | Consistent build/test/push workflow with production-style repeatability |
| Cobra               | Operational CLI framework                                   | Structured command UX for platform tasks                                |
| Bubble Tea          | Interactive TUI framework                                   | Fast incident triage with modern terminal ergonomics                    |
| Helm                | Kubernetes package manager                                  | Standardized, fast installation of complex monitoring stacks            |
| Prometheus          | Metrics collection and scraping                             | SRE-grade telemetry for pods, workloads, and cluster behavior           |
| Grafana             | Metrics visualization dashboards                            | Real-time insight into CPU/memory trends and scaling behavior           |
| Observability Stack | Monitoring foundation (Prometheus + Grafana + kube metrics) | Makes platform behavior visible and measurable                          |
| Monitoring Stack    | Namespace-level and service-level monitoring workflows      | Supports troubleshooting and performance optimization                   |
| LLM APIs            | AI incident analysis and guidance (planned)                 | Faster root-cause understanding and remediation support                 |
| AWS Bedrock         | Managed enterprise LLM access (planned)                     | Secure, scalable AI integration for ops workflows                       |

---

## 🔁 CI/CD Pipeline

NeuroOps uses automated CI/CD with GitHub Actions to move delivery from manual steps to production-style automation.

What was implemented:

- GitHub Actions workflow for CI/CD lifecycle,
- automated Go testing,
- Docker image build automation,
- automated AWS ECR push,
- secure credential injection through GitHub Secrets,
- repeatable delivery pipeline behavior.

Workflow location:

- `.github/workflows/ci-cd.yml`

Why CI/CD matters:

- reliability: every push is validated by the same checks,
- repeatability: delivery process becomes deterministic,
- automation: reduces manual deployment errors,
- DevOps workflow maturity: integrates code quality and release flow.

GitHub Secrets used:

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_ACCOUNT_ID`

Security note:

- GitHub only shows secret names in repository settings and hides secret values to protect credentials.

PowerShell commands used to create workflow structure:

```powershell
cd C:\Users\Admin\Desktop\kubepulse-aws
mkdir .github
cd .github
mkdir workflows
cd workflows
notepad ci-cd.yml
```

Command explanations:

- `cd ...kubepulse-aws`: go to project root.
- `mkdir .github`: create GitHub configuration directory.
- `mkdir workflows`: create workflow directory recognized by GitHub Actions.
- `notepad ci-cd.yml`: create/edit CI/CD pipeline YAML file.

---

## 🔄 GitOps with Helm + Argo CD

Kubernetes resources are packaged as a Helm chart in `charts/neuroops` and deployed by Argo CD from Git.

Delivery flow:

```text
merge to main
  -> GitHub Actions: go test, go build, docker build, push to ECR (tag = commit SHA)
  -> gitops-update job: writes the ECR repository + SHA into charts/neuroops/values.yaml and commits it
  -> Argo CD detects the commit and syncs the chart into the neuroops namespace
```

Argo CD runs with `automated` sync, `prune`, and `selfHeal`: manual `kubectl edit` changes are reverted to what Git declares.

Install Argo CD and register the app:

```powershell
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl wait --for=condition=available deployment/argocd-server -n argocd --timeout=300s
kubectl apply -f argocd/application.yaml
```

Open the Argo CD UI:

```powershell
kubectl port-forward svc/argocd-server -n argocd 8081:443
kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | %{[System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($_))}
```

- URL: `https://localhost:8081`, username `admin`.

Pulling from private ECR on Minikube:

```powershell
minikube addons configure registry-creds
minikube addons enable registry-creds
```

- the addon creates an `awsecr-cred` Secret and refreshes the ECR token; set `imagePullSecrets: [{name: awsecr-cred}]` in `charts/neuroops/values.yaml`.

Chart values:

| Value                                | Default            | Purpose                                 |
| ------------------------------------ | ------------------ | --------------------------------------- |
| `image.repository` / `image.tag`     | set by CI          | Image deployed by Argo CD               |
| `resources.requests.cpu` / `limits`  | `100m` / `500m`    | Scheduling baseline and CPU cap         |
| `hpa.minReplicas` / `maxReplicas`    | `1` / `5`          | Autoscaling bounds                      |
| `hpa.targetCPUUtilizationPercentage` | `50`               | Scale-up threshold                      |
| `livenessProbe` / `readinessProbe`   | `/health`, `/ready`| Self-healing and traffic gating         |

---

## 📊 Observability: Metrics, Logs, Traces, Alerts, SLOs

NeuroOps covers the three observability signals plus SLO-based alerting.

| Signal  | Source in the API                                                   | Backend                             |
| ------- | ------------------------------------------------------------------- | ----------------------------------- |
| Metrics | `/metrics` (`client_golang`): requests, latency histogram, failure flag | Prometheus (ServiceMonitor)     |
| Logs    | JSON logs (`log/slog`) with `trace_id`                              | Loki, shipped by Grafana Alloy      |
| Traces  | OpenTelemetry (`otelhttp`, OTLP/gRPC)                               | Tempo                               |
| Events  | Kubernetes events (restarts, probe failures, scaling)               | Loki, via Alloy                     |
| Alerts  | `PrometheusRule` (SLO burn rate, latency, restarts, HPA saturation) | Alertmanager → Slack                |

API metrics:

- `neuroops_http_requests_total{route, method, code}`
- `neuroops_http_request_duration_seconds{route, method}` (histogram)
- `neuroops_failure_mode` (1 after `/fail`)

SLOs (configured in `charts/neuroops/values.yaml` under `slo`):

- availability: 99.5% of `/` and `/load` requests return non-5xx,
- latency: p99 on `/` under 200 ms.

Alerts (`charts/neuroops/templates/prometheusrule.yaml`):

| Alert                         | Fires when                                                 | Severity |
| ----------------------------- | ---------------------------------------------------------- | -------- |
| `NeuroOpsErrorBudgetFastBurn` | 5xx ratio > 14.4x budget over both 1h and 5m               | critical |
| `NeuroOpsErrorBudgetSlowBurn` | 5xx ratio > 6x budget over both 6h and 30m                 | warning  |
| `NeuroOpsHighLatencyP99`      | p99 on `/` above 200 ms for 5m                             | warning  |
| `NeuroOpsPodRestarted`        | a NeuroOps container restarted in the last 5m              | warning  |
| `NeuroOpsNoAvailablePods`     | 0 available replicas for 1m                                | critical |
| `NeuroOpsHPAAtMaxReplicas`    | HPA stuck at max replicas for 5m                           | warning  |

Install the stack (pinned chart versions: kube-prometheus-stack, Loki, Tempo, Alloy):

```bash
minikube start --driver=docker
minikube addons enable metrics-server
./monitoring/install.sh
```

The NeuroOps chart then adds its own `ServiceMonitor`, `PrometheusRule`, and Grafana dashboard ("NeuroOps — Service Overview"). They only render once the Prometheus Operator CRDs exist, so the chart still installs on a cluster without monitoring.

Enable Slack alerts (the webhook lives in Vault, see [Secrets with HashiCorp Vault](#-secrets-with-hashicorp-vault)):

```bash
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXX ./vault/setup.sh
# then set alerting.slack.enabled: true in charts/neuroops/values.yaml and push (Argo CD syncs it)
```

Open Grafana:

```bash
kubectl port-forward svc/neuroops-monitoring-grafana 3000:80 -n monitoring
```

Grafana access:

- URL: `http://localhost:3000`
- Username: `admin`
- Password: retrieve from Kubernetes secret.

Retrieve Grafana password command:

```powershell
kubectl get secret neuroops-monitoring-grafana -n monitoring -o jsonpath="{.data.admin-password}" | %{[System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($_))}
```

What this command does:

- reads the base64-encoded admin password from the Kubernetes secret,
- decodes it to UTF-8 text,
- prints the plain password in PowerShell.

### Measuring reliability (numbers for the README and CV)

Load tests use [k6](https://k6.io) (`brew install k6`). Get the service URL with `minikube service neuroops-service -n neuroops --url`.

Throughput, latency, and availability under 100 concurrent virtual users:

```bash
k6 run -e BASE_URL=$URL -e VUS=100 -e DURATION=10m \
  --summary-export=loadtest/results/traffic.json loadtest/traffic.js
```

- `http_reqs` rate = sustained requests/s, `http_req_duration` p99 = latency, `http_req_failed` = error rate (availability = 1 − rate).

Failure injection with recovery timing (run in a second terminal during the traffic test):

```bash
INTERVAL=30 OUTPUT=loadtest/results/failures.csv ./scripts/inject-failures.sh 20
```

- calls `/fail` inside a random pod, then reports seconds until the container restarts and is Ready again (avg / p95 / max).
- set `hpa.minReplicas: 2` first: with a single replica every restart is an outage by design.

Autoscaling under CPU pressure:

```bash
k6 run -e BASE_URL=$URL loadtest/autoscale.js
kubectl get hpa -n neuroops -w
```

Detection time: compare the injection timestamps in `failures.csv` with the `NeuroOpsPodRestarted` message time in Slack.

Observability note:

- Some Grafana panels can show `No data` on Minikube. This is normal for local clusters with low traffic or short scrape windows.

---

## 🔐 Secrets with HashiCorp Vault

Vault is the single source of truth for NeuroOps secrets. Nothing secret is stored in Git, Helm values, or created by hand with `kubectl create secret`.

```text
Vault KV v2 (kv/neuroops/*)
  -> Kubernetes auth: ServiceAccount neuroops/neuroops-vault, read-only policy
  -> Vault Secrets Operator syncs each path into a Kubernetes Secret (refresh every 30s)
  -> consumers (Alertmanager Slack receiver) read the Secret; changes apply without a redeploy
```

| Piece                       | Where                                   | Role                                                         |
| --------------------------- | --------------------------------------- | ------------------------------------------------------------ |
| Vault server                | `vault/values.yaml`                     | Standalone, file storage, Prometheus metrics                 |
| Setup script                | `vault/setup.sh`                        | Install, init, unseal, KV engine, Kubernetes auth, policy, audit log |
| Policy `neuroops-read`      | `vault/setup.sh`                        | `read` on `kv/data/neuroops/*` only (least privilege)        |
| VaultConnection / VaultAuth / VaultStaticSecret | `charts/neuroops/templates/vault.yaml` | Deployed by Argo CD with the app        |
| Rotation timing             | `vault/measure-rotation.sh`             | Seconds from a Vault write to the updated Kubernetes Secret  |

Setup (after `./monitoring/install.sh`, so Vault metrics are scraped too):

```bash
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXX ./vault/setup.sh
kubectl get vaultstaticsecret,secret slack-webhook -n neuroops
```

- re-run `./vault/setup.sh` after a cluster restart to unseal Vault.
- the unseal key and root token are written to `vault/.vault-init.json` (git-ignored). This is a lab setup: on EKS, use auto-unseal with AWS KMS, revoke the root token, and enable TLS.
- the Vault audit log goes to stdout, so Alloy ships it to Loki: `{namespace="vault"}` in Grafana.

Measure secret rotation:

```bash
RUNS=5 ./vault/measure-rotation.sh
```

---

## 🤖 Future AI Features (Strategic Direction)

NeuroOps is evolving into an end-to-end AI-powered Kubernetes observability and remediation platform.

### 🤖 AI Incident Analysis

Planned pipeline:

- collect Kubernetes events,
- collect pod logs,
- collect Prometheus metrics,
- detect anomalies,
- send incident context to an LLM API,
- generate root cause analysis,
- generate remediation suggestions.

### 🤖 AI Remediation Engine

Planned capabilities:

- suggest or trigger safe actions like `kubectl rollout restart`,
- adjust scaling with policy controls,
- orchestrate guided remediation playbooks,
- enforce approval gates for high-risk actions.

### 🤖 AI Observability Layer

Planned integrations:

- Prometheus metric analysis,
- Kubernetes log analysis,
- LLM-powered incident summaries,
- AWS Bedrock/OpenAI integration for analysis quality,
- anomaly detection and trend forecasting.

Example AI output:

```text
Incident detected:
CPU usage exceeded threshold.

AI analysis:
Potential stress-test activity detected.

Suggested remediation:
- optimize /load endpoint
- tune HPA thresholds
- increase CPU requests
```

---

## ⚙️ PowerShell Commands (End-to-End)

### 1) Current run commands (important)

```powershell
cd C:\Users\Admin\Desktop\kubepulse-aws\cli-go
go build -o neuroops.exe
.\neuroops.exe tui --url http://127.0.0.1:56227
```

```powershell
minikube service kubepulse-service -n kubepulse --url
```

Notes:

- the Minikube URL port is dynamic and can change between restarts,
- always use `minikube service ... --url` to get the current endpoint.

### 2) Build and run API locally

```powershell
cd C:\Users\Admin\Desktop\neuroops\app
go run main.go
```

- launches the API server for health/readiness/load/failure testing.

### 3) Docker build and test

```powershell
cd C:\Users\Admin\Desktop\neuroops\app
docker build -t neuroops:local .
docker run -p 8080:8080 neuroops:local
```

- builds a portable image and validates runtime behavior.

### 4) Kubernetes deploy and validate

```powershell
cd C:\Users\Admin\Desktop\neuroops
helm upgrade --install neuroops charts/neuroops -n neuroops --create-namespace -f charts/neuroops/values-local.yaml
kubectl get pods -n neuroops
kubectl get svc -n neuroops
```

- installs the Helm chart (Deployment, Service, HPA) with the local `neuroops:local` image and verifies workload/service health.
- for the GitOps flow with Argo CD, see [GitOps with Helm + Argo CD](#-gitops-with-helm--argo-cd).

### 5) HPA verification

```powershell
kubectl get hpa -n neuroops
kubectl top pods -n neuroops
```

- confirms autoscaling policy and current pod CPU behavior.

### 6) Terraform provisioning flow

```powershell
cd C:\Users\Admin\Desktop\neuroops\terraform
terraform init
terraform validate
terraform plan
terraform apply
```

- initializes and provisions cloud resources via IaC.

### 7) ECR login and push

```powershell
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <ECR_URL>
docker tag neuroops:local <ECR_URL>:latest
docker push <ECR_URL>:latest
```

- authenticates Docker to ECR and publishes an image artifact.

---

## ☸️ Kubernetes Configuration Examples

### Deployment probes

```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
  failureThreshold: 2

readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
  failureThreshold: 1
```

In failure mode, `/` and `/ready` return 503 and `/health` returns 500. Readiness fails after one check (≤5 s), which removes the pod from the Service; liveness fails after two (~10 s), which restarts it.

### Horizontal Pod Autoscaler

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: neuroops-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: neuroops-api
  minReplicas: 1
  maxReplicas: 5
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 50
```

---

## 🔄 CI/CD & Delivery Engineering

NeuroOps now uses an automated delivery workflow that validates code and prepares deployable artifacts on every change.

Pipeline responsibilities:

- pull request validation (lint/build/test),
- artifact consistency checks,
- Docker build and ECR push,
- deployment-ready handoff to Kubernetes environments.

Workflow file:

- `.github/workflows/ci-cd.yml`

Example workflow shape:

```yaml
name: neuroops-ci-cd
on:
  push:
    branches: ["main"]
  pull_request:

jobs:
  build-test-publish:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.22"
      - run: go test ./...
      - run: go build ./...
      - run: docker build -t neuroops:ci .
```

---

## 📈 Observability Strategy

Current observability model:

- Kubernetes events + pod status for platform signals,
- endpoint health/readiness for application-state truth,
- HPA metrics for demand-driven scaling insight,
- Prometheus scrape-based telemetry,
- Grafana dashboards for CPU/memory behavior and namespace monitoring,
- AI roadmap to convert telemetry into actionable remediation guidance.

Local environment reality:

- Grafana `No data` panels can appear in Minikube with low traffic and are expected during initial setup.

---

## 🧯 Troubleshooting Playbook

| Symptom                           | Likely Cause                                          | Fast Fix                                                      |
| --------------------------------- | ----------------------------------------------------- | ------------------------------------------------------------- | ----------------- |
| `go.mod file not found`           | Wrong working directory                               | `cd` into module folder containing `go.mod`                   |
| `ImagePullBackOff`                | Image not reachable by cluster                        | Push to registry or load local image into cluster runtime     |
| Probe failures                    | Endpoint path/port mismatch                           | Verify `/health` and `/ready` routes and container port       |
| HPA not scaling                   | Missing metrics-server or low load                    | Install metrics-server and generate CPU load                  |
| `no basic auth credentials`       | ECR login expired/missing                             | Re-run `aws ecr get-login-password ...                        | docker login ...` |
| Service unreachable               | Pod not ready or selector mismatch                    | Check `kubectl get pods`, labels, and service selectors       |
| `helm` not recognized             | PowerShell session did not refresh PATH after install | Restart PowerShell and run `helm version` again               |
| Grafana login failed              | Wrong/old admin password                              | Retrieve current password from Kubernetes secret and retry    |
| Docker commands fail              | Docker Desktop not running                            | Start Docker Desktop before Minikube, build, or push          |
| Minikube URL unavailable          | Cluster is stopped                                    | Run `minikube start --driver=docker` and re-check service URL |
| Grafana panels show `No data`     | Low traffic or fresh scrape cycle                     | Generate load and wait for scrape interval                    |
| GitHub Secrets values not visible | GitHub masks secret values by design                  | Confirm secret names exist and rerun workflow                 |

---

## 🖼️ Observability Screenshots

### Grafana dashboard

![Grafana dashboard](docs/screenshots/grafana-dashboard.png)

### Kubernetes metrics

![Kubernetes metrics](docs/screenshots/kubernetes-metrics.png)

### HPA scaling

![HPA scaling](docs/screenshots/hpa-scaling.png)

### NeuroOps TUI

![NeuroOps TUI](docs/screenshots/neuroops-tui.png)

### GitHub Actions pipeline

![GitHub Actions pipeline](docs/screenshots/github-actions-pipeline.png)

---

## 🧑‍💼 Recruiter Impact Highlights

- Designed and implemented a cloud-native platform project from API to orchestration.
- Demonstrated production reliability controls (probes, self-healing, autoscaling).
- Added CI/CD automation with secure secret-based cloud delivery.
- Added observability stack adoption with Prometheus/Grafana and dashboard workflows.
- Built operator-focused interfaces (CLI/TUI) for practical incident workflows.
- Applied DevOps and Platform Engineering patterns (Docker, Kubernetes, IaC, CI/CD, monitoring).
- Established AI-forward architecture for incident analysis and remediation automation.

---

## 🧠 Learning Outcomes

- How Kubernetes turns application signals into automated resilience behavior.
- How autoscaling policies improve stability under variable load.
- How IaC + CI/CD reduce drift and increase delivery confidence.
- How Prometheus and Grafana make workload behavior measurable.
- How local load generation can validate HPA and observability assumptions.
- How AI can augment SRE decision-making with context-aware recommendations.

---

## 🗺️ Roadmap

NeuroOps is evolving into a full AI-powered Kubernetes observability and remediation platform.

### Near-term

- Harden CI/CD with image scanning and release policies.
- Expand Grafana dashboards for service-level and namespace-level insights.
- Add richer CLI/TUI diagnostics for operators.

### Mid-term

- Incident analysis pipeline from K8s events, logs, and Prometheus signals.
- LLM-powered root cause analysis and structured incident summaries.
- OpenAI/AWS Bedrock integration options for enterprise-safe AI workflows.

### Long-term

- AI remediation suggestions with risk scoring and approval gates.
- Safe remediation workflows with rollback strategy.
- Autonomous incident response loops for low-risk events.

---

## 🤝 Contributing

Contributions are welcome.

- Open an issue for bugs or feature ideas.
- Propose improvements through pull requests.
- Keep reliability, security, and operator UX as first-class priorities.

---

## 📜 License

This project is intended for portfolio, learning, and engineering demonstration purposes.

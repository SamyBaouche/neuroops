# NeuroOps 🚀

<div align="center">

## AI-powered Kubernetes observability and remediation platform

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)](#-technology-stack)
[![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?logo=docker&logoColor=white)](#-technology-stack)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Orchestrated-326CE5?logo=kubernetes&logoColor=white)](#-technology-stack)
[![Terraform](https://img.shields.io/badge/Terraform-IaC-7B42BC?logo=terraform&logoColor=white)](#-technology-stack)
[![AWS](https://img.shields.io/badge/AWS-Cloud%20Native-FF9900?logo=amazonaws&logoColor=white)](#-technology-stack)
[![CI/CD](https://img.shields.io/badge/GitHub%20Actions-CI%2FCD-2088FF?logo=githubactions&logoColor=white)](#-cicd--delivery-engineering)

</div>

---

## 🌍 Vision

NeuroOps started as a Kubernetes self-healing lab and evolved into a production-oriented, AI-native platform engineering project.

The mission is simple and ambitious:

- ⚡ detect operational risk early,
- 🧠 explain incidents with AI-assisted reasoning,
- 🛠️ recommend (and later automate) safe remediation,
- ☁️ run cloud-native workloads with real SRE patterns.

This repository is designed to showcase platform engineering maturity: reliability controls, autoscaling, deployment automation, observability workflows, and the next generation of AI-assisted operations.

---

## 🎯 Project Goals

- Build a resilient Kubernetes-ready service with health and readiness semantics.
- Operate workload behavior under stress using controlled failure and load simulation.
- Apply cloud-native deployment practices with Docker, Kubernetes, and Infrastructure as Code.
- Provide an operator-first CLI/TUI workflow for fast diagnosis from the terminal.
- Evolve toward AI incident analysis and AI remediation loops.

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
- test automation,
- container build/push workflows,
- repeatable delivery gates.

Why it matters:

- faster and safer releases,
- reduced manual deployment risk,
- measurable engineering velocity.

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
            +--> Future: Prometheus + Grafana + AI analysis loop
```

Architecture principles:

- reliability-first workload design,
- platform automation over manual ops,
- observable systems with AI-driven next steps,
- safe remediation before full autonomy.

---

## 🧰 Technology Stack

| Technology     | Purpose                                        | Why It Matters                                                |
| -------------- | ---------------------------------------------- | ------------------------------------------------------------- |
| Go             | Backend API and operational endpoints          | High performance, simplicity, strong cloud-native ecosystem   |
| Docker         | Build and package immutable artifacts          | Reliable portability across local and cloud runtimes          |
| Kubernetes     | Workload orchestration and self-healing        | Production-grade scheduling, recovery, and service resilience |
| Terraform      | Infrastructure as Code provisioning            | Reproducible and auditable cloud automation                   |
| AWS            | Cloud infrastructure foundation                | Enterprise-ready networking, security, and managed services   |
| GitHub Actions | CI/CD automation                               | Consistent build/test/deploy lifecycle at scale               |
| Cobra          | Operational CLI framework                      | Structured command UX for platform tasks                      |
| Bubble Tea     | Interactive TUI framework                      | Fast incident triage with modern terminal ergonomics          |
| Prometheus     | Metrics collection (planned integration)       | Standardized telemetry for SRE-grade observability            |
| Grafana        | Visualization dashboards (planned integration) | Real-time operations insight and alerting UX                  |
| LLM APIs       | AI incident analysis and guidance (planned)    | Faster root-cause understanding and remediation support       |
| AWS Bedrock    | Managed enterprise LLM access (planned)        | Secure, scalable AI integration for ops workflows             |

---

## 🤖 Future AI Features (Strategic Direction)

NeuroOps is evolving into an end-to-end AI-powered Kubernetes observability and remediation platform.

### 🤖 AI Incident Analysis

Planned pipeline:

- collect Kubernetes events,
- collect pod logs,
- collect metrics,
- detect anomalies,
- send incident context to an LLM API,
- generate root cause analysis,
- generate remediation suggestions.

Example AI narrative:

- Incident: CPU spike detected
- Platform reaction: HPA scaled from `1 -> 5` pods
- AI recommendation:
  - optimize hotspot endpoint,
  - increase CPU requests/limits,
  - enable response caching.

### 🤖 AI Remediation Engine

Planned capabilities:

- suggest or trigger `kubectl rollout restart`,
- adjust scaling safely under policy controls,
- orchestrate guided remediation playbooks,
- enforce approval gates for high-risk actions.

### 🤖 AI Observability Layer

Planned integrations:

- Prometheus metrics ingestion,
- Grafana signal visualization,
- AI metric interpretation,
- anomaly detection and trend forecasting.

---

## ⚙️ PowerShell Commands (End-to-End)

### 1) Run NeuroOps CLI/TUI

```powershell
cd C:\Users\Admin\Desktop\neuroops\cli-go
go build -o neuroops.exe
.\neuroops.exe tui --url http://127.0.0.1:PORT
```

What each command does:

- `cd ...\cli-go`: moves into the CLI module.
- `go build -o neuroops.exe`: compiles the CLI binary.
- `.\neuroops.exe tui --url ...`: starts the interactive operations dashboard.

### 2) Build and Run API Locally

```powershell
cd C:\Users\Admin\Desktop\neuroops\app
go run main.go
```

- Launches the API server for health/readiness/load/failure testing.

### 3) Docker Build and Test

```powershell
cd C:\Users\Admin\Desktop\neuroops\app
docker build -t neuroops:local .
docker run -p 8080:8080 neuroops:local
```

- Builds a portable image and validates runtime behavior.

### 4) Kubernetes Deploy and Validate

```powershell
cd C:\Users\Admin\Desktop\neuroops\k8s
kubectl apply -f namespace.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl get pods -n neuroops
kubectl get svc -n neuroops
```

- Applies manifests and verifies workload/service health.

### 5) HPA Verification Example

```powershell
kubectl get hpa -n neuroops
kubectl top pods -n neuroops
```

- Confirms autoscaling policy and current pod CPU behavior.

### 6) Terraform Provisioning Flow

```powershell
cd C:\Users\Admin\Desktop\neuroops\terraform
terraform init
terraform validate
terraform plan
terraform apply
```

- Initializes and provisions cloud resources via IaC.

### 7) ECR Login and Push Example

```powershell
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <ECR_URL>
docker tag neuroops:local <ECR_URL>:latest
docker push <ECR_URL>:latest
```

- Authenticates Docker to ECR and publishes an image artifact.

---

## ☸️ Kubernetes Configuration Examples

### Deployment probes

```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 10
  periodSeconds: 10

readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
```

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
          averageUtilization: 60
```

---

## 🔄 CI/CD & Delivery Engineering

NeuroOps follows a modern CI/CD posture:

- pull request validation (lint/build/test),
- artifact consistency checks,
- container build and registry push,
- deployment-ready handoff to Kubernetes environments.

Example workflow shape:

```yaml
name: neuroops-ci
on:
  push:
    branches: ["main"]
  pull_request:

jobs:
  build-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.22"
      - run: go test ./...
      - run: go build ./...
```

---

## 📈 Observability Strategy

Current and planned observability model:

- Kubernetes events + pod status for platform signals,
- endpoint health/readiness for application-state truth,
- HPA metrics for demand-driven scaling insight,
- future Prometheus/Grafana for advanced telemetry and dashboards,
- future AI layer to convert telemetry into actionable remediation guidance.

---

## 🧯 Troubleshooting Playbook

| Symptom                     | Likely Cause                       | Fast Fix                                                  |
| --------------------------- | ---------------------------------- | --------------------------------------------------------- | ------------- |
| `go.mod file not found`     | Wrong working directory            | `cd` into module folder containing `go.mod`               |
| `ImagePullBackOff`          | Image not reachable by cluster     | Push to registry or load local image into cluster runtime |
| Probe failures              | Endpoint path/port mismatch        | Verify `/health` and `/ready` routes and container port   |
| HPA not scaling             | Missing metrics-server or low load | Install metrics-server and generate CPU load              |
| `no basic auth credentials` | ECR login expired/missing          | Re-run `aws ecr get-login-password ...                    | docker login` |
| Service unreachable         | Pod not ready or selector mismatch | Check `kubectl get pods`, labels, and service selectors   |

---

## 🖥️ Screenshots (Placeholders)

- 📷 CLI/TUI overview: `docs/screenshots/tui-overview.png`
- 📷 Kubernetes pods + HPA scaling: `docs/screenshots/k8s-hpa.png`
- 📷 CI/CD pipeline run: `docs/screenshots/github-actions.png`
- 📷 Future observability dashboard: `docs/screenshots/grafana-neuroops.png`

---

## 🧑‍💼 Recruiter Impact Highlights

- Designed and implemented a cloud-native platform project from API to orchestration.
- Demonstrated production reliability controls (probes, self-healing, autoscaling).
- Built operator-focused interfaces (CLI/TUI) for practical incident workflows.
- Applied DevOps and Platform Engineering patterns (Docker, Kubernetes, IaC, CI/CD).
- Established AI-forward architecture for incident analysis and remediation automation.

---

## 🧠 Learning Outcomes

- How Kubernetes turns application signals into automated resilience behavior.
- How autoscaling policies improve stability under variable load.
- How IaC + CI/CD reduce drift and increase delivery confidence.
- How AI can augment SRE decision-making with context-aware recommendations.
- How to design a platform project that is both technically strong and portfolio-ready.

---

## 🗺️ Roadmap

### Near-term

- Add production-grade GitHub Actions pipelines (test, build, image scan, release).
- Add Prometheus metrics scraping and Grafana dashboards.
- Add richer CLI commands for diagnostics and environment operations.

### Mid-term

- Integrate LLM-powered incident summarization and RCA generation.
- Add policy-driven remediation recommendations with human approval gates.
- Add environment profiles for dev/stage/prod workflow parity.

### Long-term

- Autonomous remediation engine with safety guardrails and rollback strategy.
- SLO-aware decisioning with anomaly detection and trend prediction.
- Multi-cluster observability and fleet-wide AI operations assistant.

---

## 🤝 Contributing

Contributions are welcome.

- Open an issue for bugs or feature ideas.
- Propose improvements through pull requests.
- Keep reliability, security, and operator UX as first-class priorities.

---

## 📜 License

This project is intended for portfolio, learning, and engineering demonstration purposes.

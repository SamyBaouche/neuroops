# KubePulse AWS 🚀

KubePulse AWS is a cloud-native DevOps portfolio project that simulates how modern production systems behave under failure and load. It starts as a Go HTTP API, then evolves into a containerized and cloud-ready workload using Docker, Terraform, and AWS.

The project is designed to showcase practical Platform Engineering skills: observability-friendly endpoints, failure simulation, Infrastructure as Code, registry publishing, and operational troubleshooting.

## 🎯 Project Goal

Build a realistic self-healing system foundation that can:

- Expose liveness and readiness signals for orchestration platforms.
- Simulate failure and stress conditions in a controlled way.
- Run consistently across local and cloud environments.
- Provision cloud resources reproducibly with Terraform.
- Support an eventual CI/CD and Kubernetes deployment workflow.

## 🧱 Tech Stack

- Go 1.22
- Docker (multi-stage image build)
- AWS CLI
- Terraform
- AWS ECR (Elastic Container Registry)
- Git + GitHub
- PowerShell (Windows)

## 🏗️ Architecture Overview

Current architecture (Day 1 and Day 2 scope):

1. A Go API exposes operational endpoints:
   - `/` for service confirmation
   - `/health` for liveness-style checks
   - `/ready` for readiness-style checks
   - `/fail` to simulate app failure
   - `/load` to simulate CPU pressure
2. Docker packages the API into a lightweight runtime image.
3. Terraform provisions an ECR repository in AWS.
4. Docker image is tagged and pushed to ECR.

This creates a practical base for Kubernetes probes, autoscaling, and self-healing behavior in upcoming phases.

## 📅 Project Progress

### Day 1: Application + Container Baseline

- Built a Go HTTP API.
- Implemented endpoints: `/`, `/health`, `/ready`, `/fail`, `/load`.
- Simulated application failure and CPU load.
- Ran and tested service locally.
- Containerized the app with Docker.

### Day 2: AWS + Terraform + ECR

- Installed AWS CLI and Terraform.
- Created Terraform configuration files.
- Provisioned AWS ECR repository using Terraform.
- Built Docker image for Day 2.
- Authenticated Docker against AWS ECR.
- Tagged and pushed image to ECR.
- Verified image in AWS Console.
- Resolved real-world setup and environment issues.

## ⚙️ API Endpoints

| Endpoint  | Method | Purpose                        | Typical Response                            |
| --------- | ------ | ------------------------------ | ------------------------------------------- |
| `/`       | GET    | Basic service check            | `200 OK` - KubePulse AWS is running         |
| `/health` | GET    | Liveness-style health signal   | `200 OK` healthy or `500` service unhealthy |
| `/ready`  | GET    | Readiness-style traffic signal | `200 OK` ready                              |
| `/fail`   | GET    | Activates failure mode         | `500` failure mode activated                |
| `/load`   | GET    | Simulates CPU usage for ~5s    | `200 OK` CPU load simulated                 |

## 💻 Run Locally (Go)

From the `app` directory:

```powershell
go mod init kubepulse
go run main.go
```

Quick endpoint test examples:

```powershell
curl.exe http://localhost:8080/
curl.exe http://localhost:8080/health
curl.exe http://localhost:8080/ready
curl.exe http://localhost:8080/fail
curl.exe http://localhost:8080/health
curl.exe http://localhost:8080/load
```

## 🐳 Docker Usage

From the `app` directory:

```powershell
docker build -t kubepulse:day1 .
docker run -p 8080:8080 kubepulse:day1
```

What this does:

- Builds a multi-stage image.
- Runs the service on local port 8080.

## ☁️ Terraform + AWS ECR Setup

From the repository root:

```powershell
cd terraform
terraform init
terraform fmt
terraform validate
terraform plan
terraform apply
```

Terraform provisions an ECR repository named from `project_name`.

## 🔐 AWS Configuration

Install tooling:

```powershell
winget install Amazon.AWSCLI
winget install Hashicorp.Terraform
```

Configure AWS credentials:

```powershell
aws configure
```

Recommended minimum setup:

- Use an IAM user with permissions for ECR and Terraform-managed resources.
- Set the correct default region (`us-east-1` in this project).

## 📦 Push Docker Image to AWS ECR

Authenticate Docker to ECR:

```powershell
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <ECR_URL>
```

Tag and push image:

```powershell
docker tag kubepulse:day2 <ECR_URL>:day2
docker push <ECR_URL>:day2
```

After push, verify the image tag in AWS Console under ECR.

## 🗂️ Project Structure

```text
kubepulse-aws/
|-- README.md
|-- app/
|   |-- Dockerfile
|   |-- go.mod
|   `-- main.go
`-- terraform/
      |-- main.tf
      |-- outputs.tf
      |-- providers.tf
      `-- variables.tf
```

## 🧠 What I Learned

- How liveness and readiness endpoints support platform-level reliability.
- How to package Go workloads with efficient multi-stage Docker builds.
- How Terraform enables repeatable cloud provisioning.
- How ECR authentication and tagging work in real workflows.
- Why region consistency is critical across CLI, Terraform, and AWS Console.
- How to debug practical environment and tooling issues quickly.

## 🚧 Next Steps

1. Deploy KubePulse AWS to Kubernetes (EKS or local cluster first).
2. Add liveness and readiness probes in deployment manifests.
3. Add HPA/VPA experiments for load behavior.
4. Integrate CI/CD with GitHub Actions for build and push automation.
5. Add observability stack (metrics, logs, alerts).
6. Implement rollback-safe deployment strategies.

## 🛠️ Troubleshooting & Issues I Faced

- Go not recognized:
  Fixed by installing Go and updating system `PATH`.
- Docker daemon not running:
  Fixed by starting Docker Desktop before build/run.
- Terraform not recognized:
  Fixed by installing Terraform via `winget`.
- AWS credentials missing:
  Fixed with `aws configure` and valid IAM user keys.
- Wrong AWS region (`us-east-2` vs `us-east-1`):
  Fixed by switching region in AWS Console and aligning CLI commands.
- Nested `terraform` folder created accidentally:
  Fixed by moving files to the correct directory and removing duplicate folder.
- PowerShell does not support `touch`:
  Used `New-Item` for file creation.
- ECR image not visible:
  Root cause was wrong region selected in AWS Console.

## 🧾 Real PowerShell Commands Used

### Go

```powershell
go mod init kubepulse
go run main.go
```

### Docker

```powershell
docker build -t kubepulse:day1 .
docker run -p 8080:8080 kubepulse:day1
```

### Git

```powershell
git init
git add .
git commit -m "Day 1 - Go API + Docker"
git push
```

### Tool Installation

```powershell
winget install Amazon.AWSCLI
winget install Hashicorp.Terraform
```

### AWS

```powershell
aws configure
```

### Terraform

```powershell
terraform init
terraform fmt
terraform validate
terraform plan
terraform apply
```

### ECR Login and Push

```powershell
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <ECR_URL>
docker tag kubepulse:day2 <ECR_URL>:day2
docker push <ECR_URL>:day2
```

### Navigation and Cleanup

```powershell
cd ..
cd terraform
pwd
ls
Remove-Item -Recurse -Force .\terraform
```

---

KubePulse AWS is progressing from a local reliability test service into a cloud-native, production-style platform engineering project. ✅

# KubePulse AWS 🚀

KubePulse AWS is a cloud-native DevOps and Platform Engineering portfolio project. It starts with a Go API, then grows into a containerized, infrastructure-managed, and Kubernetes-deployed system that simulates real production workflows.

The project focuses on reliability and operations-first thinking: health checks, readiness checks, failure simulation, load simulation, image publishing, and cluster deployment.

## 🎯 Project Goal

Build a practical self-healing platform foundation that can:

- expose liveness and readiness signals for orchestration systems,
- simulate failures and CPU pressure safely,
- run the same service locally and in cloud-ready form,
- provision infrastructure with Terraform,
- evolve into CI/CD automation with GitHub Actions.

## 🧱 Tech Stack

- Go 1.22
- Docker
- Terraform
- AWS CLI
- AWS ECR
- Kubernetes
- Minikube
- PowerShell (Windows)
- Git + GitHub

## 🏗️ Architecture Overview

Current Day 1 to Day 3 architecture:

1. Go API exposes operational endpoints (`/`, `/health`, `/ready`, `/fail`, `/load`).
2. Docker packages the API as a portable image.
3. Terraform provisions an AWS ECR repository.
4. Docker pushes images to ECR for cloud-oriented workflows.
5. Kubernetes manifests deploy the app on Minikube locally.
6. A NodePort service exposes the app for browser and CLI testing.

## 📅 Project Progress

### Day 1 - Go API + Docker

- Built a Go HTTP API.
- Implemented endpoints: `/`, `/health`, `/ready`, `/fail`, `/load`.
- Simulated failure and CPU load.
- Ran the app locally.
- Containerized with Docker.

### Day 2 - Terraform + AWS ECR

- Installed AWS CLI and Terraform.
- Created `providers.tf`, `variables.tf`, `main.tf`, `outputs.tf`.
- Provisioned ECR using Terraform.
- Built Docker image.
- Logged in to ECR.
- Tagged and pushed image.
- Verified image in AWS Console (`us-east-1`).

### Day 3 - Kubernetes Local Deployment

- Installed `kubectl` and Minikube.
- Started local Kubernetes with Docker driver.
- Created `namespace.yaml`, `deployment.yaml`, `service.yaml`.
- Deployed app to Kubernetes.
- Fixed `ImagePullBackOff` using local image workflow.
- Loaded Docker image into Minikube.
- Exposed service using NodePort.
- Accessed service using `minikube service`.

## ⚙️ API Endpoints

| Endpoint  | Method | Purpose                           | Typical Response                    |
| --------- | ------ | --------------------------------- | ----------------------------------- |
| `/`       | GET    | Basic service check               | `200 OK` - KubePulse AWS is running |
| `/health` | GET    | Liveness-style check              | `200 OK` healthy or `500` unhealthy |
| `/ready`  | GET    | Readiness-style check             | `200 OK` ready                      |
| `/fail`   | GET    | Activates failure mode            | `500` failure mode activated        |
| `/load`   | GET    | Simulates CPU load for ~5 seconds | `200 OK` CPU load simulated         |

## 💻 Run Locally with Go

From the `app` directory:

```powershell
go mod init kubepulse
go run main.go
```

What these commands do:

- `go mod init kubepulse`: creates a Go module file (`go.mod`) for dependency tracking. Run this once when starting the project.
- `go run main.go`: compiles and starts the API in one step.

## 🐳 Docker Usage

From the `app` directory:

```powershell
docker build -t kubepulse:day1 .
docker run -p 8080:8080 kubepulse:day1
```

What these commands do:

- `docker build -t kubepulse:day1 .`: builds the image from `Dockerfile` and tags it as `kubepulse:day1`.
- `docker run -p 8080:8080 kubepulse:day1`: starts the container and maps host port `8080` to container port `8080`.

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

What these commands do:

- `cd terraform`: move into Terraform configuration directory.
- `terraform init`: download providers and initialize Terraform state.
- `terraform fmt`: format `.tf` files to standard style.
- `terraform validate`: check configuration syntax and internal consistency.
- `terraform plan`: preview changes Terraform will make.
- `terraform apply`: create/update resources in AWS.

ECR login, tag, and push:

```powershell
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <ECR_URL>
docker tag kubepulse:day2 <ECR_URL>:day2
docker push <ECR_URL>:day2
```

What these commands do:

- `aws ecr get-login-password ... | docker login ...`: authenticates Docker to your ECR registry.
- `docker tag kubepulse:day2 <ECR_URL>:day2`: creates an ECR-ready tag.
- `docker push <ECR_URL>:day2`: uploads image to ECR.

## ☸️ Kubernetes Local Deployment

Install and verify tooling:

```powershell
winget install -e --id Kubernetes.kubectl
winget install -e --id Kubernetes.minikube
kubectl version --client
minikube version
```

Start cluster and verify:

```powershell
minikube start --driver=docker
kubectl get nodes
```

Create manifest files:

```powershell
mkdir k8s
cd k8s
New-Item namespace.yaml
New-Item deployment.yaml
New-Item service.yaml
```

Deploy resources:

```powershell
kubectl apply -f namespace.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl get pods -n kubepulse
kubectl describe pods -n kubepulse
```

Use local image to fix pull issues:

```powershell
docker build -t kubepulse:local .
minikube image load kubepulse:local
kubectl apply -f deployment.yaml
kubectl delete pod -n kubepulse --all
kubectl get deployment kubepulse-api -n kubepulse -o wide
```

Access service:

```powershell
minikube service kubepulse-service -n kubepulse
```

## 🧾 Real PowerShell Commands Used

```powershell
go mod init kubepulse
go run main.go

docker build -t kubepulse:day1 .
docker run -p 8080:8080 kubepulse:day1

git init
git add .
git commit -m "Day 1 - Go API + Docker"
git push

winget install Amazon.AWSCLI
winget install Hashicorp.Terraform

aws configure

terraform init
terraform fmt
terraform validate
terraform plan
terraform apply

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <ECR_URL>
docker tag kubepulse:day2 <ECR_URL>:day2
docker push <ECR_URL>:day2

winget install -e --id Kubernetes.kubectl
winget install -e --id Kubernetes.minikube

kubectl version --client
minikube version

minikube start --driver=docker
kubectl get nodes

mkdir k8s
cd k8s
New-Item namespace.yaml
New-Item deployment.yaml
New-Item service.yaml

kubectl apply -f namespace.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

kubectl get pods -n kubepulse
kubectl describe pods -n kubepulse

docker build -t kubepulse:local .
minikube image load kubepulse:local

kubectl apply -f deployment.yaml
kubectl delete pod -n kubepulse --all

kubectl get deployment kubepulse-api -n kubepulse -o wide

minikube service kubepulse-service -n kubepulse
```

### Command Explanations (Beginner Friendly)

- `git init`: start a new Git repository in the current folder.
- `git add .`: stage all current changes.
- `git commit -m "Day 1 - Go API + Docker"`: save staged snapshot with a message.
- `git push`: upload commits to remote GitHub repository.
- `winget install Amazon.AWSCLI`: install AWS CLI.
- `winget install Hashicorp.Terraform`: install Terraform CLI.
- `aws configure`: set AWS access key, secret key, region, and output format locally.
- `kubectl version --client`: verify `kubectl` installation.
- `minikube version`: verify Minikube installation.
- `kubectl get nodes`: confirm local cluster node is ready.
- `mkdir k8s`: create Kubernetes manifests folder.
- `cd k8s`: enter Kubernetes folder.
- `New-Item namespace.yaml`: create namespace manifest file in PowerShell.
- `New-Item deployment.yaml`: create deployment manifest file.
- `New-Item service.yaml`: create service manifest file.
- `kubectl apply -f ...`: create or update Kubernetes resource from file.
- `kubectl get pods -n kubepulse`: list pods in namespace.
- `kubectl describe pods -n kubepulse`: detailed pod diagnostics/events.
- `docker build -t kubepulse:local .`: build image intended for local Minikube use.
- `minikube image load kubepulse:local`: copy local image into Minikube image cache.
- `kubectl delete pod -n kubepulse --all`: force pods to restart with latest deployment/image settings.
- `kubectl get deployment kubepulse-api -n kubepulse -o wide`: detailed deployment status and image/node information.
- `minikube service kubepulse-service -n kubepulse`: open or print URL for the service endpoint.

## 🛠️ Troubleshooting & Issues I Faced

- Go not recognized:
  Installed Go and updated system `PATH`.
- Docker daemon not running:
  Started Docker Desktop before running `docker build`/`docker run`.
- Terraform not recognized:
  Installed Terraform with `winget install Hashicorp.Terraform`.
- AWS credentials missing:
  Created IAM user credentials and configured with `aws configure`.
- Wrong AWS region (`us-east-2` to `us-east-1`):
  Standardized region in CLI, Terraform variable, and AWS Console.
- Nested Terraform folder:
  Moved canonical files to top-level `terraform` folder for predictable execution.
- PowerShell does not support `touch`:
  Used `New-Item` to create files.
- `ImagePullBackOff`:
  Cluster could not pull image from registry path used in manifest.
- `no basic auth credentials`:
  Switched to local image strategy (`kubepulse:local` + `minikube image load`).
- YAML parsing error:
  Fixed indentation and key alignment in manifest files.
- `SVC_UNREACHABLE`:
  Service had no healthy running pods behind selector, resolved by fixing deployment and restarting pods.

## 🗂️ Project Structure

```text
kubepulse-aws/
|-- README.md
|-- app/
|   |-- Dockerfile
|   |-- go.mod
|   `-- main.go
|-- terraform/
|   |-- main.tf
|   |-- outputs.tf
|   |-- providers.tf
|   `-- variables.tf
`-- k8s/
    |-- namespace.yaml
    |-- deployment.yaml
    `-- service.yaml
```

## 🧠 What I Learned

- How health and readiness endpoints support platform reliability decisions.
- How Docker images bridge local development and deployment environments.
- How Terraform enforces repeatable, auditable infrastructure changes.
- How registry authentication and tagging affect deployment pipelines.
- How Kubernetes debugging depends on pods, events, selectors, and image strategy.
- Why clear structure and documentation matter for real DevOps delivery.

## 🚧 Next Steps

1. Day 4: Add liveness and readiness probes for self-healing behavior.
2. Day 5: Add HPA experiments and CLI terminal UI.
3. Day 6: Add GitHub Actions CI/CD for build, test, and push automation.
4. Day 7: Polish documentation, record demo, and publish project story.

---

KubePulse AWS now demonstrates a complete Day 1 to Day 3 progression from API to Kubernetes, with practical troubleshooting and strong portfolio-ready documentation.

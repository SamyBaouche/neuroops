# KubePulse AWS

KubePulse AWS is a cloud-native reliability engineering project focused on building a self-healing application platform using Go, Docker, Kubernetes, Terraform, AWS, and GitHub Actions.

Day 1 establishes the application foundation: a Go API that can intentionally simulate failure and CPU stress, packaged into a Docker container and executed locally.

## Project Goal

The long-term goal of KubePulse AWS is to create a production-style system that can:

- Run as cloud-native workloads
- Expose health and readiness signals
- Detect unhealthy behavior automatically
- Recover through orchestration policies (self-healing)
- Be provisioned and deployed with Infrastructure as Code and CI/CD

In short, this project is a practical journey from a local service to an automated, resilient cloud platform.

## Day 1 Scope

For Day 1, the following was implemented:

- Built a Go HTTP API with these endpoints:
  - /
  - /health
  - /ready
  - /fail
  - /load
- Added failure simulation logic to intentionally make health checks fail
- Added CPU load simulation logic to mimic high-resource pressure
- Containerized the application with Docker (multi-stage build)
- Built and ran the container locally on Windows using PowerShell

## Architecture Overview (Day 1)

Day 1 architecture is intentionally simple and focused:

- Go API:
  - Handles HTTP requests
  - Provides health/readiness signals
  - Includes controlled chaos endpoints (/fail and /load) for resilience testing
- Docker container:
  - Packages the Go binary into a lightweight runtime image
  - Provides environment consistency across local and future cloud runtimes

Flow summary:

1. You run the Go service directly or via Docker.
2. Clients call API endpoints.
3. /health reflects healthy or failing state.
4. /fail toggles the app into failure mode.
5. /load generates temporary CPU pressure for testing behavior under stress.

## API Endpoints

| Endpoint | Method | Purpose                                                                 | Typical Response                                                  |
| -------- | ------ | ----------------------------------------------------------------------- | ----------------------------------------------------------------- |
| /        | GET    | Home endpoint to confirm service is running                             | 200 OK, "KubePulse AWS is running"                                |
| /health  | GET    | Liveness-style check; returns unhealthy after failure mode is activated | 200 OK (healthy) or 500 Internal Server Error (service unhealthy) |
| /ready   | GET    | Readiness-style check indicating service can receive traffic            | 200 OK, "ready"                                                   |
| /fail    | GET    | Activates failure mode by setting an internal flag                      | 500 Internal Server Error, "failure mode activated"               |
| /load    | GET    | Simulates CPU load for 5 seconds via a busy loop                        | 200 OK, "CPU load simulated for 5 seconds"                        |

## Run Locally (PowerShell)

### Prerequisites

- Go 1.22+
- Docker Desktop (running)
- PowerShell

### 1. Initialize Go module (first time only)

```powershell
go mod init kubepulse
```

What it does:

- Creates go.mod
- Defines the module name
- Enables dependency and version management for the Go project

Note:

- This project already has go.mod, so this command is only needed if you start from scratch.

### 2. Run the API directly with Go

```powershell
go run .\app\main.go
```

What it does:

- Compiles and runs the application in one step
- Starts the HTTP server on port 8080
- Useful for rapid local development before containerization

### 3. Test API endpoints with curl

Open a second PowerShell terminal and run:

```powershell
curl.exe http://localhost:8080/
curl.exe http://localhost:8080/health
curl.exe http://localhost:8080/ready
curl.exe http://localhost:8080/fail
curl.exe http://localhost:8080/health
curl.exe http://localhost:8080/load
```

What each test does:

- GET / verifies the app is online
- First GET /health should return healthy
- GET /ready confirms readiness endpoint behavior
- GET /fail enables failure mode
- Second GET /health should now return unhealthy (500)
- GET /load triggers a 5-second CPU stress simulation

Why curl.exe on Windows:

- In PowerShell, curl may map to Invoke-WebRequest.
- curl.exe ensures you use the native curl binary syntax consistently.

### 4. Build Docker image

From the app directory:

```powershell
cd .\app
docker build -t kubepulse:day1 .
```

What it does:

- Uses the Dockerfile to build a multi-stage image
- Compiles the Go binary in a builder image
- Copies only the binary into a small Alpine runtime image
- Tags image as kubepulse:day1

### 5. Run Docker container

```powershell
docker run --name kubepulse-day1 -p 8080:8080 kubepulse:day1
```

What it does:

- Starts a container named kubepulse-day1
- Maps local port 8080 to container port 8080
- Makes the API accessible at http://localhost:8080

Optional cleanup commands:

```powershell
docker stop kubepulse-day1
docker rm kubepulse-day1
```

## Project Structure

```text
kubepulse-aws/
|-- README.md
`-- app/
	 |-- Dockerfile
	 |-- go.mod
	 `-- main.go
```

## What I Learned (Day 1)

- How to design and expose core operational endpoints (/health and /ready)
- How to simulate failure conditions in a controlled way for resilience testing
- How CPU stress behavior can be emulated to test service robustness
- How to package a Go service with a lightweight multi-stage Docker build
- How to run and validate a containerized API locally on Windows PowerShell

## Next Steps

Planned evolution for upcoming days:

1. AWS foundation:
   - Prepare target AWS environment for deployment
   - Define networking and compute strategy
2. Terraform (Infrastructure as Code):
   - Provision cloud resources declaratively
   - Make infrastructure reproducible and version-controlled
3. Kubernetes deployment:
   - Deploy KubePulse AWS as pods/services
   - Add liveness/readiness probes and self-healing behavior
4. CI/CD with GitHub Actions:
   - Automate build, test, and container publishing
   - Automate deployment workflow to cloud environments

---

Day 1 outcome: KubePulse AWS now has a working, testable, containerized service baseline ready for cloud and orchestration integration.

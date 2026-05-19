# ToggleMaster Targeting Service

Microsserviço responsável pelas regras de segmentação das Feature Flags.

## Responsabilidades

- Segmentação de usuários
- Regras de rollout
- Target por ambiente
- Target por perfil

## Stack

- Golang
- Docker
- Kubernetes
- GitHub Actions
- Amazon ECR
- Amazon EKS

## Execução local

```bash
go run cmd/main.go
Endpoint de Health Check
GET /health
Docker

Build:

docker build -t togglemaster-targeting .

Run:

docker run -p 8082:8082 togglemaster-targeting
CI/CD

Pipeline DevSecOps com:

Build
Unit Tests
GolangCI-Lint
Gosec
Trivy
Docker Build
Push para Amazon ECR
Deploy

Deploy automatizado via GitOps utilizando ArgoCD.




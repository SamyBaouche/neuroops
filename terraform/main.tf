# Creates an AWS Elastic Container Registry (ECR) repository.
# Docker images are pushed here before Kubernetes deploys them in cloud environments.
resource "aws_ecr_repository" "kubepulse" {
  # Repository name will be kubepulse-aws (from variables.tf by default).
  name = var.project_name
  # MUTABLE allows retagging existing tags (useful while learning and iterating quickly).
  image_tag_mutability = "MUTABLE"

  # Enables vulnerability scan when a new image is pushed.
  image_scanning_configuration {
    scan_on_push = true
  }

  # Helpful metadata for identification and cost/resource tracking.
  tags = {
    Project   = var.project_name
    ManagedBy = "Terraform"
  }
}

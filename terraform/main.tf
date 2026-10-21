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

# Keep storage bounded: expire the oldest images beyond 30 (multi-arch pushes create several artifacts each).
resource "aws_ecr_lifecycle_policy" "kubepulse" {
  repository = aws_ecr_repository.kubepulse.name
  policy = jsonencode({
    rules = [{
      rulePriority = 1
      description  = "Keep only the 30 most recent images"
      selection    = { tagStatus = "any", countType = "imageCountMoreThan", countNumber = 30 }
      action       = { type = "expire" }
    }]
  })
}

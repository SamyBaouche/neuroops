# Creates an AWS Elastic Container Registry (ECR) repository.
# Docker images are pushed here before Kubernetes deploys them in cloud environments.
resource "aws_ecr_repository" "kubepulse" {
  # Repository name comes from var.project_name (default: kubepulse-aws).
  name = var.project_name
  # MUTABLE allows replacing image tags while iterating quickly during development.
  image_tag_mutability = "MUTABLE"

  # Enables basic vulnerability scanning for every pushed image.
  image_scanning_configuration {
    scan_on_push = true
  }

  # Tags help identify ownership and automation source in AWS.
  tags = {
    Project   = var.project_name
    ManagedBy = "Terraform"
  }
}

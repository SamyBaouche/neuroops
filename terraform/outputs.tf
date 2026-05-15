# Prints the full ECR repository URL after 'terraform apply'.
# You can reuse this value for docker tag/push commands.
output "ecr_repository_url" {
  description = "URL of the ECR repository"
  value       = aws_ecr_repository.kubepulse.repository_url
}

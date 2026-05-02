# Input variable used by the AWS provider to choose deployment region.
variable "aws_region" {
  description = "AWS region used for the project"
  type        = string
  default     = "us-east-1"
}

# Input variable used for naming AWS resources consistently.
variable "project_name" {
  description = "Project name used for AWS resources"
  type        = string
  default     = "kubepulse-aws"
}

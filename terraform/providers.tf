# Terraform-level configuration: declares which providers this project needs.
terraform {
  required_providers {
    aws = {
      # Official AWS provider maintained by HashiCorp.
      source = "hashicorp/aws"
      # Compatible 6.x releases. Terraform picks the newest matching version.
      version = "~> 6.0"
    }
  }
}

# AWS provider configuration.
# Region comes from variables.tf so it can be changed in one place.
provider "aws" {
  region = var.aws_region
}

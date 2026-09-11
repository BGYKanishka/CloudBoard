output "alb_dns_name" {
  description = "The DNS name of the Application Load Balancer"
  value       = module.compute.alb_dns_name
}

output "ecr_repository_url" {
  description = "The URL of the ECR repository"
  value       = module.ecr.repository_url
}

output "db_endpoint" {
  description = "The endpoint of the RDS database"
  value       = module.database.db_endpoint
}

output "github_actions_role_arn" {
  description = "The ARN of the IAM role to be used by GitHub Actions"
  value       = module.iam.github_actions_role_arn
}

output "github_actions_role_arn" {
  description = "The ARN of the IAM role to be used by GitHub Actions"
  value       = aws_iam_role.github_actions.arn
}

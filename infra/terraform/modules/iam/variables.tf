variable "github_repo" {
  description = "The GitHub repository in the format username/repo (e.g. BGYKanishka/CloudBoard)"
  type        = string
}

variable "role_name" {
  description = "Name of the IAM role for GitHub Actions"
  type        = string
  default     = "github-actions-deploy-role"
}

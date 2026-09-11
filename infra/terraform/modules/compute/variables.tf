variable "vpc_id" {
  description = "The VPC ID"
  type        = string
}

variable "public_subnet_ids" {
  description = "The public subnet IDs for the ALB"
  type        = list(string)
}

variable "app_subnet_ids" {
  description = "The subnet IDs to deploy EC2 instances into"
  type        = list(string)
}

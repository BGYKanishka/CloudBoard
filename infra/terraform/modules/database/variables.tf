variable "vpc_id" {
  description = "The VPC ID"
  type        = string
}

variable "vpc_cidr" {
  description = "The CIDR block of the VPC to allow DB access from"
  type        = string
}

variable "db_subnet_ids" {
  description = "The private subnet IDs for the database"
  type        = list(string)
}

variable "db_name" {
  description = "The name of the database"
  type        = string
  default     = "cloudboard"
}

variable "db_username" {
  description = "The master username for the database"
  type        = string
  default     = "postgres"
}

variable "db_password" {
  description = "The master password for the database"
  type        = string
  sensitive   = true
}

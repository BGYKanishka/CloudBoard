variable "vpc_cidr" {
  description = "The CIDR block for the VPC"
  type        = string
}

variable "public_subnets" {
  description = "List of public subnet CIDR blocks"
  type        = list(string)
}

variable "private_app_subnets" {
  description = "List of private subnet CIDR blocks for applications"
  type        = list(string)
}

variable "private_db_subnets" {
  description = "List of private subnet CIDR blocks for databases"
  type        = list(string)
}

variable "azs" {
  description = "List of Availability Zones"
  type        = list(string)
}

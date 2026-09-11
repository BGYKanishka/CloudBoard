terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

module "vpc" {
  source = "../../modules/network"
  
  vpc_cidr = "10.0.0.0/16"
  public_subnets = ["10.0.1.0/24", "10.0.2.0/24"]
  private_app_subnets = ["10.0.11.0/24", "10.0.12.0/24"]
  private_db_subnets = ["10.0.21.0/24", "10.0.22.0/24"]
  azs = ["us-east-1a", "us-east-1b"]
}

resource "random_id" "bucket_id" {
  byte_length = 4
}

module "storage" {
  source = "../../modules/storage"
  bucket_name = "cloudboard-profile-pics-dev-${random_id.bucket_id.hex}"
}

module "ecr" {
  source = "../../modules/ecr"
  repository_name = "cloudboard-api-dev"
}

module "database" {
  source = "../../modules/database"
  vpc_id = module.vpc.vpc_id
  vpc_cidr = "10.0.0.0/16"
  db_subnet_ids = module.vpc.private_db_subnet_ids
  db_name = "cloudboard"
  db_username = "postgres"
  db_password = var.db_password
}

module "compute" {
  source = "../../modules/compute"
  vpc_id = module.vpc.vpc_id
  public_subnet_ids = module.vpc.public_subnet_ids
  app_subnet_ids = module.vpc.public_subnet_ids # Using public subnets to save NAT Gateway costs
}

module "iam" {
  source = "../../modules/iam"
  github_repo = "BGYKanishka/CloudBoard"
}

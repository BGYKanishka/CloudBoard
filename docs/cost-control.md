# Cost Control Guide

> **IMPORTANT**: This project is built for learning. Destroy AWS resources when you are not using them to avoid unexpected charges.

## LOW COST / GENERALLY SAFE
- **VPC & Subnets**: Free to create.
- **Security Groups & Route Tables**: Free.
- **IAM Roles & Policies**: Free.
- **Internet Gateway**: Free.

## POTENTIALLY COSTLY
- **S3**: Very cheap, but costs scale with storage and data transfer.
- **ECR**: Storage costs for Docker images.
- **CloudWatch**: Logs and custom metrics cost money. Avoid noisy alarms.

## HIGHER COST / CREATE ONLY TEMPORARILY
- **EC2 Instances**: You pay per hour. (e.g., t3.micro is cheap but still costs money if left running 24/7).
- **RDS (PostgreSQL)**: Managed databases are relatively expensive. Always delete when done learning.
- **Application Load Balancer (ALB)**: Charges an hourly rate and per LCU (Load Balancer Capacity Unit).
- **NAT Gateway**: Extremely expensive hourly rate and data processing fee. **Avoid using NAT Gateways in learning projects unless explicitly testing private outbound internet access.**
- **EKS (Elastic Kubernetes Service)**: The control plane costs ~$73/month flat, PLUS the cost of EC2 worker nodes. **DO NOT CREATE EKS UNLESS YOU ARE ACTIVELY TESTING IT, THEN DESTROY IMMEDIATELY.**

## Cleanup Procedure
1. `terraform destroy` (Destroys infrastructure).
2. Delete EKS cluster (if manually created).
3. Delete ECR images.
4. Empty S3 buckets (Terraform can't destroy buckets with items in them).

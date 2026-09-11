# AWS Learning Guide

This guide explains the AWS components used in the CloudBoard architecture.

## EC2 (Elastic Compute Cloud)
**What is it?** Virtual servers in the cloud.
**Why do we need it?** To run our Docker containers (backend API).
**What problem does it solve?** Provides scalable computing capacity without buying hardware.
**How does it connect?** Our Docker images from ECR run on EC2 instances.
**What breaks if removed?** The backend API won't run, breaking the application.

## VPC (Virtual Private Cloud)
**What is it?** A logically isolated section of the AWS Cloud.
**Why do we need it?** To securely host our resources in a private network.
**What problem does it solve?** Network security and isolation.
**How does it connect?** All EC2, RDS, and Load Balancers live inside the VPC.
**What breaks if removed?** Nothing can securely communicate; resources would be exposed to the public internet by default or lack routing.

## Subnet
**What is it?** A range of IP addresses in your VPC.
**Why do we need it?** To isolate resources (public vs private).
**What problem does it solve?** Allows placing the database in a private subnet (no internet access) and load balancer in a public subnet.
**How does it connect?** EC2 instances and RDS are assigned to specific subnets.
**What breaks if removed?** We couldn't logically separate public and private resources.

*(Expand on RDS, S3, ALB, Auto Scaling, IAM, CloudWatch, Route 53, etc. as you continue your learning journey!)*

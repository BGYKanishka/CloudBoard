# Architecture

CloudBoard is designed to be a learning tool for AWS and DevOps.

## Mermaid Diagram

```mermaid
flowchart TD

    User --> DNS
    DNS --> ALB
    ALB --> EC2A
    ALB --> EC2B
    EC2A --> RDS
    EC2B --> RDS
    EC2A --> S3
    EC2B --> S3

    GitHub --> GitHubActions
    GitHubActions --> ECR
    ECR --> EC2A
    ECR --> EC2B

    CloudWatch --> EC2A
    CloudWatch --> EC2B
    CloudWatch --> RDS
```

## Description
- **Frontend**: React SPA running in a user's browser, fetching data from the API.
- **API Backend**: A Go REST API running on EC2 instances or EKS pods.
- **Database**: Managed PostgreSQL (AWS RDS).
- **Storage**: AWS S3 for profile pictures.
- **Load Balancing**: AWS Application Load Balancer distributes traffic between instances.
- **Monitoring**: AWS CloudWatch.

Yes. And because **today is the expiration day for your AWS credits**, I would make this a **hands-on AWS/DevOps lab project**, not a polished production application. The goal is to touch as many real concepts as possible while keeping the architecture understandable and avoiding expensive AWS components that provide little learning value.

AWS currently says Free Tier credits expire according to the expiration date shown in the Billing/Credits page, and after the free period, remaining resources can incur normal charges. ([Amazon Web Services][1])

# Project: CloudBoard

We're going to build a small **Task Management API + web application**.

It doesn't need to be complicated.

Users can:

```text
Register/login
      ↓
Create task
      ↓
View tasks
      ↓
Update task
      ↓
Delete task
```

The application itself is deliberately simple.

The **infrastructure** is the difficult/educational part.

---

# 1. What we're going to build

Final architecture:

```text
                         INTERNET
                            |
                            ↓
                       Route 53
                    (optional domain)
                            |
                            ↓
                    Application Load
                       Balancer
                            |
              ┌─────────────┴─────────────┐
              ↓                           ↓
           EC2 #1                      EC2 #2
        Docker Container            Docker Container
              |                           |
              └─────────────┬─────────────┘
                            ↓
                       PostgreSQL
                          RDS
                            |
                            ↓
                      S3 Object Storage


       ┌─────────────────────────────────────────┐
       │                  AWS                    │
       │                                         │
       │                  VPC                    │
       │                                         │
       │     Public Subnets    Private Subnets  │
       │                                         │
       │      ALB + EC2          RDS             │
       │                                         │
       └─────────────────────────────────────────┘


GitHub
   |
   ↓
GitHub Actions
   |
   ├── Test
   ├── Build
   ├── Docker image
   └── Push
        ↓
       ECR
        ↓
     EC2 deploy
```

Then we'll take the **same Docker application** and deploy it to Kubernetes:

```text
Docker Image
     |
     ↓
    ECR
     |
     ↓
    EKS
     |
     ├── Pod
     ├── Service
     ├── Deployment
     └── LoadBalancer
```

Then Terraform:

```text
Terraform
    |
    ↓
AWS
    |
    ├── VPC
    ├── Subnets
    ├── EC2
    ├── RDS
    ├── ALB
    ├── Security Groups
    └── IAM
```

So this single project covers almost everything we've discussed.

---

# 2. What each piece teaches you

This is the important part.

| Technology            | What you learn            |
| --------------------- | ------------------------- |
| EC2                   | Virtual servers           |
| Linux                 | Server administration     |
| VPC                   | Cloud networking          |
| Subnet                | Network segmentation      |
| Internet Gateway      | Internet connectivity     |
| Route tables          | Traffic routing           |
| Security Groups       | Firewall/network security |
| IAM                   | Identity & permissions    |
| RDS                   | Managed database          |
| S3                    | Object storage            |
| ALB                   | Load balancing            |
| Auto Scaling          | Horizontal scaling        |
| ECR                   | Container registry        |
| Docker                | Containerization          |
| CloudWatch            | Monitoring/logs           |
| Route 53              | DNS                       |
| ACM                   | HTTPS/TLS                 |
| GitHub Actions        | CI/CD                     |
| Terraform             | Infrastructure as Code    |
| EKS                   | Kubernetes                |
| Kubernetes Deployment | Container orchestration   |

That's basically your **beginner → intermediate cloud/DevOps lab**.

---

# 3. The application itself

Don't waste time making a beautiful frontend.

Use something simple such as:

```text
Frontend
React

Backend
Go / Node.js

Database
PostgreSQL
```

Since you're learning Go, I'd actually use:

```text
React
   ↓
Go REST API
   ↓
PostgreSQL
```

Example API:

```text
GET    /api/tasks
POST   /api/tasks
GET    /api/tasks/:id
PUT    /api/tasks/:id
DELETE /api/tasks/:id
GET    /health
```

The `/health` endpoint is important because AWS load balancers and monitoring can use it.

Example:

```text
GET /health

{
    "status": "ok"
}
```

---

# 4. Start locally

Before touching AWS:

```text
React
  ↓
Go API
  ↓
PostgreSQL
```

Run everything locally.

Then Dockerize it.

Eventually:

```text
docker compose up
```

and:

```text
Frontend container
Backend container
PostgreSQL container
```

At this point you have learned:

**application + Docker + containers.**

---

# 5. Build the Docker image

Create:

```text
Dockerfile
```

Then:

```bash
docker build -t cloudboard-api .
```

Run it:

```bash
docker run ...
```

Now you understand:

```text
Source Code
     ↓
Dockerfile
     ↓
Docker Image
     ↓
Docker Container
```

---

# 6. Put the image in Amazon ECR

AWS ECR is your private container registry.

Architecture:

```text
Your Mac
   |
   ↓
Docker build
   |
   ↓
Docker image
   |
   ↓
ECR
```

This teaches the same general concept as:

```text
Docker Hub
GHCR
ECR
ACR
```

You previously asked about these, so this gives you practical experience with one of them.

---

# 7. Build the AWS network manually

This is one of the most important sections.

We'll create:

```text
VPC
│
├── Public Subnet A
│
├── Public Subnet B
│
├── Private Subnet A
│
└── Private Subnet B
```

Why two?

Because AWS Availability Zones are designed to let you spread resources across separate infrastructure locations within a region.

We'll use:

```text
ALB
 ↓
EC2
```

across multiple AZs.

And:

```text
RDS
```

in private subnets.

---

# 8. Internet Gateway

Create:

```text
Internet
    |
    ↓
Internet Gateway
    |
    ↓
VPC
```

The Internet Gateway gives your VPC a path to/from the internet for appropriately configured resources.

You'll then learn:

```text
Route Table
```

For example:

```text
0.0.0.0/0
      ↓
Internet Gateway
```

This is where your earlier networking knowledge starts connecting with AWS.

---

# 9. Security Groups

Create separate security groups.

### ALB

Allow:

```text
80
443
```

from the internet.

### EC2

Allow:

```text
Application port
```

**only from the ALB security group**.

### RDS

Allow:

```text
5432 PostgreSQL
```

**only from the EC2 security group**.

So:

```text
Internet
   |
   ↓
ALB
   |
   ↓
EC2
   |
   ↓
RDS
```

Not:

```text
Internet → RDS
```

This is one of the most important security concepts you'll learn.

---

# 10. Launch EC2

Now create your first cloud server.

Conceptually:

```text
AWS
 |
 └── EC2
      |
      ├── Linux
      ├── CPU
      ├── RAM
      └── Storage
```

SSH into it:

```bash
ssh ...
```

Now you're doing actual Linux administration on a cloud server.

Install/use:

```text
Docker
Git
AWS CLI
```

Then pull your image from ECR.

```text
ECR
 ↓
EC2
 ↓
Docker
 ↓
CloudBoard
```

---

# 11. Deploy the application manually first

Don't immediately automate everything.

Do it manually once.

For example:

```text
EC2
 ↓
docker login
 ↓
docker pull
 ↓
docker run
```

Then open your browser.

```text
http://EC2-IP
```

And see your application running from an AWS server.

That moment is important.

You'll have gone from:

```text
localhost
```

to:

```text
Internet
   ↓
AWS
   ↓
Linux server
   ↓
Docker
   ↓
Your application
```

---

# 12. Add RDS

Now remove the PostgreSQL container.

Instead:

```text
EC2
  |
  ↓
RDS PostgreSQL
```

Your application connects using something like:

```text
DB_HOST
DB_PORT
DB_NAME
DB_USER
DB_PASSWORD
```

Don't hardcode those into Git.

This teaches:

**managed cloud database + networking + credentials.**

AWS RDS pricing depends on region, instance type and usage; AWS also bills partial instance usage, so we'll create it only when needed and tear it down afterward. ([Amazon Web Services][2])

---

# 13. Add S3

Add one small feature to CloudBoard:

### User profile image

```text
User
 ↓
Application
 ↓
S3
 ↓
profile-picture.jpg
```

Now you understand why S3 exists.

You're no longer using S3 just because someone told you:

> "S3 is object storage."

You're using it for a real reason.

---

# 14. Add Application Load Balancer

Instead of:

```text
Internet
   ↓
EC2
```

change it to:

```text
Internet
   ↓
ALB
   ↓
EC2
```

Then create:

```text
EC2 #1
EC2 #2
```

Now:

```text
                 ALB
               /     \
              ↓       ↓
           EC2 #1   EC2 #2
```

The ALB distributes traffic between them.

This gives you practical **horizontal scaling**.

---

# 15. Add Auto Scaling

Now we tell AWS:

> Keep at least 2 instances.

and:

> Add more when the workload increases.

Conceptually:

```text
Normal traffic

EC2 #1
EC2 #2
```

High traffic:

```text
EC2 #1
EC2 #2
EC2 #3
EC2 #4
```

Low traffic again:

```text
EC2 #1
EC2 #2
```

This is the practical difference between:

**vertical scaling**

and

**horizontal scaling**.

---

# 16. Add CloudWatch

Now monitor the infrastructure.

Watch:

```text
CPU
Network
Instance health
Application logs
ALB metrics
```

And create an alarm such as:

```text
CPU > 70%
      ↓
CloudWatch Alarm
```

You are now entering the **operations** side of DevOps.

---

# 17. Add IAM properly

Don't run everything using the root account.

Create an IAM setup where:

```text
Administrator
    |
    └── infrastructure management

EC2 IAM Role
    |
    ├── Read ECR
    ├── Write CloudWatch logs
    └── Access required AWS services
```

The key lesson:

```text
Who?
  ↓
Can do what?
  ↓
On which resource?
```

Never use your AWS root credentials inside the application or GitHub Actions.

---

# 18. Add GitHub Actions

Now automate deployment.

Your workflow becomes:

```text
Developer
   |
   ↓
git push
   |
   ↓
GitHub
   |
   ↓
GitHub Actions
   |
   ├── Run tests
   ├── Build application
   ├── Build Docker image
   ├── Push image to ECR
   └── Deploy
```

Now you have a real CI/CD pipeline.

---

# 19. Use GitHub → AWS without storing long-lived AWS keys

This is where you'll learn a much better real-world pattern:

**GitHub Actions + AWS IAM + OIDC**

Conceptually:

```text
GitHub Actions
      |
      | temporary identity
      ↓
AWS IAM
      |
      ↓
ECR / AWS resources
```

Instead of putting a permanent:

```text
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
```

into GitHub.

This is a very worthwhile DevOps security lesson.

---

# 20. Add Route 53

Buy/use a domain only if you already have one.

Then:

```text
cloudboard.example.com
          |
          ↓
       Route 53
          |
          ↓
         ALB
```

Now your application isn't:

```text
http://some-random-address
```

It's:

```text
https://cloudboard.example.com
```

---

# 21. Add HTTPS

Use:

**AWS Certificate Manager**

Conceptually:

```text
User
  |
 HTTPS
  ↓
ALB
  |
  ↓
EC2
```

This teaches you:

```text
TLS
Certificates
HTTPS
Secure communication
```

and connects nicely to the TLS/certificate concepts we've already discussed.

---

# 22. Now introduce Terraform

At this point you'll have manually created:

```text
VPC
Subnets
Route Tables
Internet Gateway
Security Groups
EC2
RDS
ALB
Auto Scaling
IAM
```

Now we destroy the manual infrastructure.

Then recreate it with:

```text
Terraform
```

For example:

```text
terraform/
├── main.tf
├── variables.tf
├── outputs.tf
├── network.tf
├── compute.tf
├── database.tf
├── security.tf
└── iam.tf
```

The idea becomes:

```text
Terraform code
       ↓
terraform plan
       ↓
terraform apply
       ↓
AWS infrastructure
```

This is a **massive** DevOps lesson.

---

# 23. Finally: Kubernetes

Now take your existing Docker image:

```text
ECR
 |
 ↓
cloudboard-api
```

and deploy it to Amazon EKS.

You learn:

```text
Cluster
 ↓
Node
 ↓
Pod
 ↓
Deployment
 ↓
Service
```

For example:

```text
EKS
│
├── Node
│    ├── Pod
│    │    └── CloudBoard
│    │
│    └── Pod
│         └── CloudBoard
│
└── Service
```

The application remains basically the same.

Only the **deployment platform** changes.

That teaches you something important:

> Your application and your infrastructure/orchestration platform are separate concepts.

---

# 24. EKS is the expensive part

This is the one part I would **not leave running**.

AWS currently charges EKS standard clusters at **$0.10 per cluster-hour**, before the additional costs of worker nodes, storage, public IPv4 addresses and other resources. ([Amazon Web Services][3])

So:

```text
Create EKS
      ↓
Learn Kubernetes
      ↓
Deploy application
      ↓
Test
      ↓
DELETE CLUSTER
```

Don't create an EKS cluster and forget about it.

---

# 25. Your complete learning journey

I would do the project in this exact order:

```text
PHASE 1
Build application locally
        ↓
React + Go + PostgreSQL

PHASE 2
Docker
        ↓
Docker image
        ↓
Run locally

PHASE 3
AWS fundamentals
        ↓
IAM
        ↓
Region
        ↓
VPC
        ↓
Subnets
        ↓
Route tables
        ↓
Internet Gateway
        ↓
Security Groups

PHASE 4
EC2
        ↓
Linux server
        ↓
SSH
        ↓
Docker
        ↓
Run application

PHASE 5
ECR
        ↓
Push Docker image
        ↓
EC2 pulls image

PHASE 6
RDS
        ↓
PostgreSQL

PHASE 7
S3
        ↓
File/profile upload

PHASE 8
ALB
        ↓
EC2 #1
EC2 #2

PHASE 9
Auto Scaling
        ↓
Horizontal scaling

PHASE 10
CloudWatch
        ↓
Logs
        ↓
Metrics
        ↓
Alarms

PHASE 11
Route 53
        ↓
DNS

PHASE 12
ACM
        ↓
HTTPS

PHASE 13
GitHub Actions
        ↓
CI/CD

PHASE 14
Terraform
        ↓
Infrastructure as Code

PHASE 15
EKS
        ↓
Kubernetes
```

---

# 26. What the final project teaches

By the end, you will understand this picture:

```text
                         USER
                           |
                           ↓
                         DNS
                           |
                           ↓
                        HTTPS
                           |
                           ↓
                     Load Balancer
                           |
                ┌──────────┴──────────┐
                ↓                     ↓
             EC2 #1               EC2 #2
                |                     |
             Docker                Docker
                |                     |
                └──────────┬──────────┘
                           ↓
                     PostgreSQL
                         RDS
                           |
                           ↓
                         S3


     AWS VPC
     ├── Subnets
     ├── Route Tables
     ├── Internet Gateway
     └── Security Groups

     IAM
     └── Access control

     CloudWatch
     └── Monitoring

     ECR
     └── Container images

     Auto Scaling
     └── More/fewer EC2 instances

     GitHub Actions
     └── CI/CD

     Terraform
     └── Infrastructure as Code

     EKS
     └── Kubernetes
```

That is **far more useful than simply getting a website deployed to AWS**.

### One important change because your credits expire today

Don't start by building every component at once.

Start with:

```text
CloudBoard
   ↓
Docker
   ↓
ECR
   ↓
VPC
   ↓
EC2
   ↓
RDS
```

Get that working first. Then add ALB, Auto Scaling, S3, CloudWatch, CI/CD, Terraform, and finally EKS. That way, even if we hit the credit-expiration deadline, you still finish with a working AWS project rather than half-built infrastructure.

Also, **don't try to deliberately spend all $100**. The useful outcome is the knowledge and working infrastructure, not consuming the credits. AWS itself recommends deleting resources you don't need as the free-period expiration approaches because active resources can otherwise generate charges. ([AWS Documentation][4])

## Our starting point

We'll call the project **CloudBoard**, and I recommend:

```text
Frontend → React
Backend  → Go
Database → PostgreSQL
Container → Docker
Cloud    → AWS
CI/CD    → GitHub Actions
IaC      → Terraform
K8s      → EKS
```

**Next, we should build the application itself first**, with a deliberately tiny Go API and React frontend, then containerize it. That gives us a clean application to carry through the entire AWS journey.

[1]: https://aws.amazon.com/free/free-tier-faqs/?utm_source=chatgpt.com "Free Tier FAQs"
[2]: https://aws.amazon.com/rds/postgresql/pricing/?utm_source=chatgpt.com "Amazon RDS for PostgreSQL Pricing"
[3]: https://aws.amazon.com/eks/pricing/?utm_source=chatgpt.com "Amazon EKS Pricing"
[4]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/avoid-charges-after-free-tier.html?utm_source=chatgpt.com "Avoiding unexpected charges after Free Tier - AWS Billing"

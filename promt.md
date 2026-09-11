You are the lead software engineer, DevOps engineer, cloud architect, and technical mentor for this project.

I am a beginner learning AWS, cloud infrastructure, Linux, Docker, networking, CI/CD, Terraform, and Kubernetes.

Your job is to BUILD the complete application and its development/deployment infrastructure in this workspace, but also structure the project so I can use it as a learning project for AWS and DevOps.

IMPORTANT:
- Do not blindly over-engineer the application.
- The business application itself must remain simple.
- The infrastructure and DevOps side is the educational focus.
- Do not use unnecessary technologies just to make the project look complicated.
- Do not hide important configuration behind abstractions.
- Prefer explicit, understandable code and configuration.
- Every major infrastructure concept should be represented by real files/configuration where appropriate.
- Do not deploy anything to AWS automatically unless I explicitly ask you later.
- Do not create AWS resources, delete AWS resources, or spend AWS credits without my explicit instruction.
- Build everything locally first.
- Make the repository production-like but beginner-readable.
- Whenever there is a choice between a clever solution and an easy-to-understand solution, choose the easy-to-understand solution.

==================================================
PROJECT
==================================================

Project name:

CloudBoard

CloudBoard is a small task management application.

The purpose of the application is NOT to demonstrate advanced frontend development.

The purpose is to create a realistic application that can later be deployed through AWS infrastructure and used to demonstrate:

- Linux
- Networking
- AWS
- Docker
- Container Registry
- EC2
- VPC
- Subnets
- Route Tables
- Internet Gateway
- Security Groups
- IAM
- RDS
- S3
- Application Load Balancer
- Auto Scaling
- CloudWatch
- DNS
- HTTPS
- GitHub Actions
- CI/CD
- Terraform
- Kubernetes
- EKS

==================================================
HIGH LEVEL ARCHITECTURE
==================================================

The final target architecture should conceptually look like this:

Internet
    |
    v
Route 53
    |
    v
HTTPS / TLS
    |
    v
Application Load Balancer
    |
    +----------------------+
    |                      |
    v                      v
EC2 instance 1         EC2 instance 2
    |                      |
    +----------+-----------+
               |
               v
        Go backend
               |
               v
        PostgreSQL RDS
               
Application files / profile images
               |
               v
              S3

Monitoring:
CloudWatch

Container images:
ECR

CI/CD:
GitHub Actions

Infrastructure as Code:
Terraform

Container orchestration extension:
EKS / Kubernetes

Networking:
VPC
  |
  +-- Public subnets
  |
  +-- Private subnets
  |
  +-- Route tables
  |
  +-- Internet Gateway
  |
  +-- Security Groups

Identity:
IAM

IMPORTANT:
This architecture is the TARGET architecture.
Do not actually provision it automatically.

==================================================
APPLICATION STACK
==================================================

Use:

Frontend:
- React
- TypeScript
- Vite

Backend:
- Go
- Standard library where reasonable
- net/http for the HTTP server
- PostgreSQL

Database:
- PostgreSQL

Containers:
- Docker
- Docker Compose for local development

Infrastructure as Code:
- Terraform

CI/CD:
- GitHub Actions

Container registry target:
- AWS ECR

Cloud target:
- AWS

Kubernetes target:
- Amazon EKS

==================================================
APPLICATION FUNCTIONALITY
==================================================

Keep the application simple.

Implement:

1. User registration
2. User login
3. User logout
4. Create task
5. View tasks
6. View single task
7. Update task
8. Delete task
9. Mark task completed
10. Profile page
11. Profile image upload
12. Application health endpoint

Task fields:

- id
- user_id
- title
- description
- status
- priority
- created_at
- updated_at

Statuses:

- pending
- in_progress
- completed

Priorities:

- low
- medium
- high

User fields:

- id
- name
- email
- password_hash
- profile_image_url
- created_at
- updated_at

==================================================
AUTHENTICATION
==================================================

Implement a simple secure authentication system.

Use:

- password hashing
- HTTP-only secure cookies or another straightforward session/token approach
- proper password validation
- authorization checks
- users must only be able to access their own tasks

Do not implement a complicated OAuth provider for the first version.

The AWS/DevOps project should focus on infrastructure rather than authentication complexity.

Never:

- store plaintext passwords
- commit secrets
- hardcode credentials
- hardcode AWS credentials
- put secrets in Docker images
- put secrets in frontend source code

Use environment variables for local configuration.

Create:

.env.example

Never create a real .env file containing secrets.

==================================================
BACKEND API
==================================================

Use REST APIs.

Create endpoints similar to:

POST   /api/auth/register
POST   /api/auth/login
POST   /api/auth/logout
GET    /api/auth/me

GET    /api/tasks
POST   /api/tasks
GET    /api/tasks/:id
PUT    /api/tasks/:id
DELETE /api/tasks/:id

PATCH  /api/tasks/:id/status

POST   /api/profile/image

GET    /health
GET    /ready

Health endpoint:

GET /health

should return a simple successful response when the application process is alive.

Readiness endpoint:

GET /ready

should verify that the application can reach required dependencies such as PostgreSQL.

These two endpoints are important because they will later be useful for:

- load balancers
- container orchestration
- Kubernetes
- monitoring
- operations

==================================================
BACKEND STRUCTURE
==================================================

Use a clean but beginner-friendly Go structure.

Prefer something similar to:

backend/
  cmd/
    server/
      main.go

  internal/
    config/
    database/
    models/
    repository/
    handlers/
    middleware/
    auth/
    storage/
    health/

  migrations/

  go.mod

Do not create dozens of meaningless abstraction layers.

Keep the architecture easy to follow.

Use dependency injection where it genuinely improves testing or clarity.

==================================================
DATABASE
==================================================

Use PostgreSQL.

Create SQL migrations.

At minimum create:

users
tasks

Create proper:

- primary keys
- foreign keys
- indexes where appropriate
- timestamps
- constraints

Do not use an ORM unless it significantly improves the project.

Prefer SQL that I can read and understand.

Create seed/sample data support for local development if useful.

==================================================
PROFILE IMAGE STORAGE
==================================================

The application must support profile image uploads.

However, design the storage layer so that local development and AWS deployment can use different implementations.

Local:
- store files in a local uploads directory

AWS:
- store files in Amazon S3

Use an interface such as:

Storage interface

Then provide:

LocalStorage
S3Storage

Do not tightly couple the application to S3.

The application should choose the storage implementation using environment configuration.

==================================================
FRONTEND
==================================================

Build a clean, modern but simple dashboard.

Pages:

/login
/register
/dashboard
/tasks
/tasks/:id
/profile

Features:

- responsive UI
- navigation
- task list
- create task form
- edit task
- delete task
- filter tasks
- mark complete
- profile information
- profile image upload
- loading states
- error states
- empty states

Do not waste most of the project time on visual polish.

Prioritize functionality and clean structure.

Use TypeScript properly.

Avoid unnecessary state-management complexity.

==================================================
LOCAL DEVELOPMENT
==================================================

The entire application must run locally with Docker Compose.

Create:

docker-compose.yml

It should provide:

frontend
backend
postgres

The architecture should be:

Browser
   |
   v
Frontend
   |
   v
Backend
   |
   v
PostgreSQL

Create appropriate Dockerfiles.

Use multi-stage Docker builds where appropriate.

The Go production image should be small.

Do not run the production container as root unless there is a genuine reason.

Add Docker healthchecks where appropriate.

==================================================
DOCKER
==================================================

Create:

backend/Dockerfile
frontend/Dockerfile

Also create a production-oriented container configuration.

Requirements:

- small images
- multi-stage builds
- no secrets inside images
- non-root containers where practical
- deterministic dependency installation
- health checks where practical
- clear exposed ports
- sensible environment configuration

Create:

.dockerignore

==================================================
LOCAL DATABASE
==================================================

Docker Compose should start PostgreSQL automatically.

The backend should wait/retry until PostgreSQL becomes available rather than crashing permanently if PostgreSQL takes a few seconds to initialize.

Implement sensible database connection retry logic.

==================================================
REVERSE PROXY
==================================================

For local production-like testing, optionally use Nginx or another simple reverse proxy.

Do not make it mandatory if it creates unnecessary complexity.

The cloud target should eventually support:

Internet
   |
   v
Application Load Balancer
   |
   v
Backend containers

==================================================
PROJECT STRUCTURE
==================================================

Create a monorepo structure similar to:

cloudboard/
│
├── frontend/
│
├── backend/
│
├── infra/
│   ├── terraform/
│   │
│   └── kubernetes/
│
├── docker/
│
├── .github/
│   └── workflows/
│
├── docs/
│
├── docker-compose.yml
├── .env.example
├── .gitignore
├── .dockerignore
├── Makefile
├── README.md
└── LICENSE

Adjust the structure if you have a clearly better beginner-friendly organization.

==================================================
MAKEFILE
==================================================

Create a Makefile with commands such as:

make dev
make build
make test
make lint
make docker-build
make docker-up
make docker-down
make db-migrate
make clean

Commands should be simple and documented.

==================================================
TESTING
==================================================

Backend:

Write unit tests for important logic.

Test:

- authentication logic
- password validation
- task logic
- authorization
- handlers where practical
- health endpoint
- database-related behavior where practical

Frontend:

Add reasonable component tests if the setup remains simple.

At minimum, the CI pipeline must be able to:

- install dependencies
- run frontend tests
- run backend tests
- build frontend
- build backend
- build Docker images

Do not fake tests just to make CI green.

==================================================
LINTING / STATIC ANALYSIS
==================================================

Go:

Use standard Go tooling:

go fmt
go vet
go test

Add golangci-lint only if it is appropriate and does not create unnecessary complexity.

Frontend:

Use ESLint and TypeScript checking.

Add scripts for:

npm run lint
npm run build
npm run test

==================================================
GITHUB ACTIONS
==================================================

Create CI/CD workflows.

Workflow 1:

.github/workflows/ci.yml

Run on:

- pull_request
- push to main

Pipeline should:

1. Checkout
2. Setup Go
3. Run gofmt check
4. Run go vet
5. Run Go tests
6. Setup Node
7. Install frontend dependencies
8. Run frontend lint
9. Run frontend tests
10. Build frontend
11. Build backend
12. Build Docker images

Do not deploy from CI yet.

Workflow 2:

.github/workflows/deploy.yml

Prepare a production deployment workflow that eventually supports:

GitHub
  |
  v
GitHub Actions
  |
  v
Authenticate to AWS using OIDC
  |
  v
ECR
  |
  v
Deploy

IMPORTANT:

Do not use permanent AWS access keys in GitHub Actions.

Design the workflow around:

GitHub Actions OIDC
+
AWS IAM role

The workflow should be clearly documented but should not require actual AWS credentials to exist in the repository.

Use placeholders where needed.

==================================================
AWS ECR
==================================================

Create documentation and scripts for eventually:

docker build
docker tag
docker push

to Amazon ECR.

Do not put an actual AWS account ID into the code.

Use variables such as:

AWS_REGION
AWS_ACCOUNT_ID
ECR_REPOSITORY

The repository should be designed so that the image names can later become:

cloudboard/backend
cloudboard/frontend

or an equivalent clean naming scheme.

==================================================
AWS EC2 DEPLOYMENT
==================================================

Create deployment documentation for running the Docker application on a Linux EC2 server.

Document:

1. Create EC2
2. Configure security group
3. SSH into server
4. Install Docker
5. Authenticate to ECR
6. Pull image
7. Run container
8. Configure environment variables
9. Check health endpoint
10. View logs
11. Restart container
12. Stop container
13. Update container

Create deployment helper scripts where useful.

Do NOT actually create the EC2 instance.

==================================================
AWS NETWORKING DESIGN
==================================================

Create Terraform configuration for a beginner-friendly but realistic AWS network.

Target:

One AWS Region

VPC
CIDR:
10.0.0.0/16

Two Availability Zones

Public subnets:
10.0.1.0/24
10.0.2.0/24

Private application subnets:
10.0.11.0/24
10.0.12.0/24

Private database subnets:
10.0.21.0/24
10.0.22.0/24

Design the network so that:

Public:
- Application Load Balancer

Private:
- application instances

Private:
- RDS PostgreSQL

Create:

- VPC
- public subnets
- private application subnets
- private database subnets
- route tables
- Internet Gateway
- NAT Gateway only if actually required

IMPORTANT COST RULE:

Do not add expensive resources just because they are common in enterprise architecture.

Explain the cost implication of NAT Gateways.

For a learning environment, provide a lower-cost alternative where practical.

==================================================
AWS SECURITY GROUP DESIGN
==================================================

Create separate security groups.

ALB security group:

Allow:
80
443

from internet.

Application security group:

Allow application traffic only from ALB security group.

Database security group:

Allow PostgreSQL 5432 only from application security group.

SSH:

Do NOT expose SSH to the entire internet in the production design.

Use:

- restricted source IP
or
- Systems Manager Session Manager where practical.

Explain why.

==================================================
AWS EC2 TERRAFORM
==================================================

Create Terraform for EC2.

Use variables.

Do not hardcode:

- AMI IDs
- region
- account IDs
- passwords
- secrets

Where possible, use AWS data sources.

The EC2 instance should:

- run Linux
- receive IAM role permissions
- be able to access ECR
- be able to send logs to CloudWatch if configured
- run Docker

Create an IAM role for EC2 rather than putting access keys on the server.

==================================================
AWS RDS TERRAFORM
==================================================

Create Terraform for PostgreSQL RDS.

Requirements:

- private subnets
- security group only allowing PostgreSQL from application security group
- encryption enabled
- automated backups configured appropriately
- deletion protection disabled for the learning environment only if that makes cleanup easier
- final snapshot behavior documented

IMPORTANT:

Clearly mark the RDS settings that can generate AWS charges.

Use a small learning-friendly instance class.

Make the instance class configurable.

Do not commit a database password.

Use Terraform variables/secrets mechanisms appropriately.

==================================================
AWS S3 TERRAFORM
==================================================

Create Terraform for an S3 bucket used for profile images.

Requirements:

- block public access
- versioning where appropriate
- encryption
- lifecycle policy where appropriate

Do NOT make the bucket publicly readable.

Application should access it using IAM permissions.

==================================================
AWS IAM
==================================================

Create least-privilege IAM policies where practical.

Examples:

EC2 role:
- pull images from ECR
- write/read required CloudWatch resources
- access required S3 bucket

GitHub Actions role:
- authenticate through OIDC
- only receive permissions needed for deployment

Do not use AdministratorAccess for application runtime.

Do not embed AWS credentials in the project.

==================================================
AWS LOAD BALANCER
==================================================

Create Terraform for an Application Load Balancer.

Architecture:

Internet
   |
   v
ALB
   |
   v
Application target group
   |
   +-- EC2 instance 1
   |
   +-- EC2 instance 2

Use:

/health

as the health check endpoint.

Document:

- listener
- target group
- health checks
- security group
- DNS relationship

==================================================
AWS AUTO SCALING
==================================================

Create Terraform for an Auto Scaling Group.

Requirements:

- configurable min size
- configurable desired size
- configurable max size

Learning defaults should be conservative to control cost.

Example:

min = 1
desired = 1
max = 2

Do not automatically create large fleets.

Explain how this demonstrates horizontal scaling.

==================================================
AWS CLOUDWATCH
==================================================

Create monitoring configuration/documentation.

Include:

- EC2 metrics
- application logs
- ALB health
- RDS metrics
- alarms where practical

Create useful alarms such as:

- high CPU
- unhealthy hosts

Do not create dozens of noisy alarms.

Document what each alarm means.

==================================================
DNS
==================================================

Prepare Terraform/documentation for Route 53.

The design should be:

domain
   |
   v
Route 53
   |
   v
ALB

Do not require the user to buy a domain.

Allow the project to function through ALB DNS if no custom domain exists.

==================================================
HTTPS
==================================================

Prepare Terraform/documentation for AWS Certificate Manager.

Target:

Internet
   |
 HTTPS
   v
ALB
   |
 HTTP
   v
Application

Explain TLS termination at the ALB.

Do not require a real certificate/domain to develop locally.

==================================================
TERRAFORM
==================================================

Create Terraform modules only where they actually improve clarity.

Do not create an enormous enterprise Terraform architecture.

Suggested structure:

infra/terraform/

  env/
    dev/

  modules/
    network/
    security/
    compute/
    database/
    storage/
    load-balancer/
    iam/

Or another clean beginner-friendly structure.

Provide:

terraform init
terraform validate
terraform plan
terraform apply
terraform destroy

documentation.

IMPORTANT:

Terraform must NOT execute automatically.

I will decide when to run apply.

==================================================
KUBERNETES
==================================================

Create Kubernetes manifests for the application.

Target EKS architecture:

EKS
 |
 +-- Deployment
 |
 +-- Pods
 |
 +-- Service
 |
 +-- ConfigMap
 |
 +-- Secret reference
 |
 +-- readiness probe
 |
 +-- liveness probe

Create:

infra/kubernetes/

deployment.yaml
service.yaml
configmap.yaml
secret.example.yaml
namespace.yaml

Use environment configuration.

Do not commit real secrets.

The backend must expose:

/health
/ready

for Kubernetes probes.

==================================================
KUBERNETES DESIGN
==================================================

The Kubernetes application should pull the Docker image from ECR.

Example:

ECR
 |
 v
EKS
 |
 v
CloudBoard Deployment
 |
 +-- Pod
 |
 +-- Pod

Create reasonable resource requests/limits.

Do not assign excessive CPU/RAM.

Create a Service.

Document how this differs from the EC2 + Docker deployment.

==================================================
EKS COST WARNING
==================================================

Very important:

EKS can cost money even before considering worker nodes and other AWS resources.

Do not create an EKS cluster automatically.

Documentation must explicitly warn:

- create only when needed
- use it for learning
- test it
- destroy it afterward

==================================================
SECRETS
==================================================

No secrets may appear in:

- source code
- git history
- README
- Dockerfiles
- Terraform committed files
- GitHub workflow literals

Create examples such as:

DATABASE_URL=change-me
JWT_SECRET=change-me
AWS_REGION=change-me
S3_BUCKET=change-me

in:

.env.example

==================================================
DOCUMENTATION
==================================================

This is extremely important.

Create a large beginner-friendly:

README.md

The README should teach me the entire project.

It should contain:

1. What CloudBoard is
2. Why the project exists
3. Architecture diagram
4. Technology stack
5. Local setup
6. Docker explanation
7. Database explanation
8. AWS architecture
9. VPC explanation
10. Subnet explanation
11. Route table explanation
12. Internet Gateway explanation
13. Security Group explanation
14. EC2 explanation
15. ECR explanation
16. RDS explanation
17. S3 explanation
18. Load Balancer explanation
19. Auto Scaling explanation
20. CloudWatch explanation
21. IAM explanation
22. Route 53 explanation
23. HTTPS explanation
24. GitHub Actions explanation
25. OIDC explanation
26. Terraform explanation
27. Kubernetes explanation
28. EKS explanation
29. Cost considerations
30. Cleanup procedures
31. Troubleshooting

Write these explanations as if teaching someone who had never heard of AWS before.

==================================================
LEARNING DOCUMENTATION
==================================================

Create:

docs/

architecture.md
aws-learning-guide.md
networking.md
docker.md
ci-cd.md
terraform.md
kubernetes.md
troubleshooting.md
cost-control.md

In aws-learning-guide.md explain:

EC2
VPC
Subnet
Availability Zone
Region
Internet Gateway
Route Table
Security Group
IAM
ECR
RDS
S3
ALB
Auto Scaling
CloudWatch
Route 53
ACM

For every service explain:

WHAT IS IT?
WHY DO WE NEED IT?
WHAT PROBLEM DOES IT SOLVE?
HOW DOES IT CONNECT TO OUR PROJECT?
WHAT WOULD BREAK IF WE REMOVED IT?

==================================================
COST CONTROL
==================================================

This project is being built as a learning project with limited AWS credits.

Treat cost control as a first-class requirement.

Create:

docs/cost-control.md

Clearly classify resources as:

LOW COST / GENERALLY SAFE
POTENTIALLY COSTLY
HIGHER COST / CREATE ONLY TEMPORARILY

Pay particular attention to:

- NAT Gateway
- RDS
- EC2
- Load Balancer
- EKS
- public IPv4 addresses
- data transfer
- CloudWatch logs

Add a cleanup checklist.

==================================================
CLEANUP
==================================================

Create scripts/documentation for cleanup.

The README must explain the order in which AWS resources should be removed.

Example concept:

1. Remove Kubernetes resources
2. Delete EKS cluster
3. Remove load balancer
4. Remove Auto Scaling
5. Remove EC2
6. Remove RDS
7. Remove S3 resources
8. Remove NAT Gateway
9. Remove subnets
10. Remove route tables
11. Remove Internet Gateway
12. Remove VPC

Adjust the actual order according to Terraform dependencies.

IMPORTANT:

Never create an automatic destructive cleanup script that can run accidentally.

==================================================
GIT
==================================================

Create an appropriate .gitignore for:

Go
Node
React
Docker
Terraform
IDE files
environment files
OS files

Ignore:

.env
*.tfstate
*.tfstate.*
.terraform/

unless there is a specific reason not to.

==================================================
PROJECT QUALITY
==================================================

Use:

- clear naming
- meaningful comments
- reasonable error handling
- structured logging
- graceful shutdown for Go
- context cancellation
- database connection cleanup
- HTTP timeouts
- sensible middleware
- proper HTTP status codes

Do not:

- swallow errors
- use panic for ordinary runtime errors
- expose internal errors to users
- hardcode configuration
- use fake placeholder implementations for important functionality

==================================================
API ERROR FORMAT
==================================================

Use a consistent error response format.

For example:

{
  "error": "task not found"
}

and optionally:

{
  "error": "validation failed",
  "details": {
    "title": "title is required"
  }
}

Keep it simple.

==================================================
OBSERVABILITY
==================================================

The Go backend should log useful operational information.

Logs should include appropriate fields such as:

- timestamp
- request method
- path
- status
- duration
- request ID where useful

Do not log:

- passwords
- JWTs
- AWS credentials
- database passwords
- sensitive tokens

==================================================
HEALTH
==================================================

Implement:

GET /health

Meaning:

"The process is alive."

Implement:

GET /ready

Meaning:

"The process is ready to serve traffic and can reach required dependencies."

This distinction must be documented because it is important for:

- ALB
- Docker
- Kubernetes
- monitoring

==================================================
GRACEFUL SHUTDOWN
==================================================

The Go server should handle:

SIGINT
SIGTERM

and gracefully shut down HTTP connections and database connections.

This is important because containers and cloud orchestration systems can terminate processes.

==================================================
DEVELOPMENT EXPERIENCE
==================================================

Make local setup as easy as:

git clone ...
cd cloudboard
cp .env.example .env
docker compose up --build

Document any additional commands.

==================================================
NO UNNECESSARY CLOUD LOCK-IN
==================================================

The application itself should be reasonably portable.

The AWS-specific parts should be isolated into:

infra/
deployment/
storage implementations/
documentation/

Do not import AWS SDK everywhere in the Go application.

Use an abstraction around S3 storage.

==================================================
ARCHITECTURE DOCUMENT
==================================================

Create a Mermaid architecture diagram.

Example:

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

Create additional diagrams for:

- VPC
- CI/CD
- Docker
- Kubernetes

==================================================
IMPORTANT DEVELOPMENT PROCESS
==================================================

Do NOT attempt to generate the entire project blindly in one step.

Work systematically.

Phase 1:
Inspect workspace.

Phase 2:
Create project skeleton.

Phase 3:
Implement backend.

Phase 4:
Implement database migrations.

Phase 5:
Implement frontend.

Phase 6:
Integrate frontend and backend.

Phase 7:
Add authentication.

Phase 8:
Add S3-compatible storage abstraction and local storage.

Phase 9:
Dockerize application.

Phase 10:
Create Docker Compose.

Phase 11:
Write tests.

Phase 12:
Create GitHub Actions CI.

Phase 13:
Create ECR deployment workflow.

Phase 14:
Create Terraform.

Phase 15:
Create Kubernetes manifests.

Phase 16:
Create documentation.

After every major phase:

- run tests
- run builds
- inspect errors
- fix errors
- verify the implementation

Do not move forward while the current phase is obviously broken.

==================================================
AGENT BEHAVIOR
==================================================

You have permission to:

- create files
- edit files
- run local commands
- install project dependencies
- run tests
- run builds
- inspect generated files
- improve code
- fix issues

You do NOT have permission to:

- create AWS infrastructure
- modify AWS resources
- delete AWS resources
- use AWS credentials
- commit secrets
- spend cloud credits

unless I explicitly give a later instruction to do so.

==================================================
WHEN YOU FINISH
==================================================

At the end:

1. Verify the directory structure.
2. Run backend tests.
3. Run frontend tests.
4. Run linting.
5. Build frontend.
6. Build backend.
7. Build Docker images.
8. Verify Docker Compose configuration.
9. Run the application locally if practical.
10. Verify health endpoint.
11. Verify readiness endpoint.
12. Verify the GitHub Actions YAML syntax as much as possible.
13. Run terraform fmt.
14. Run terraform validate where possible without cloud credentials.
15. Validate Kubernetes YAML where possible.
16. Check that no secrets were accidentally created.
17. Check .gitignore.
18. Check README completeness.

Then give me a final report containing:

- what you built
- important architectural decisions
- complete directory tree
- commands to run the project
- test results
- known limitations
- AWS components prepared
- Terraform components prepared
- Kubernetes components prepared
- CI/CD components prepared
- what I should do manually next in AWS
- which AWS resources could incur significant charges

==================================================
MOST IMPORTANT GOAL
==================================================

Build a REAL, WORKING, SIMPLE application.

The application should be simple enough that I understand the business logic.

The infrastructure should be realistic enough that I can use it to learn:

Linux
Networking
AWS
Docker
ECR
EC2
VPC
Subnets
Route tables
Internet Gateway
Security Groups
IAM
RDS
S3
ALB
Auto Scaling
CloudWatch
Route 53
HTTPS
GitHub Actions
CI/CD
Terraform
Kubernetes
EKS

Do not optimize for "lots of code".

Optimize for:

UNDERSTANDABLE
WORKING
SECURE
DEPLOYABLE
OBSERVABLE
TEACHABLE

Start by inspecting the workspace and then begin Phase 1.
# CloudBoard

CloudBoard is a complete Task Management application built to demonstrate modern cloud architecture, DevOps practices, and infrastructure-as-code.

## 🚀 What is CloudBoard?
CloudBoard is a full-stack task management application with a React frontend and a Go backend. It exists primarily as a learning project for beginners to learn AWS, DevOps, Docker, Terraform, and Kubernetes.

## 🛠 Technology Stack
- **Frontend**: React, TypeScript, Vite, TailwindCSS
- **Backend**: Go (chi router), PostgreSQL, JWT Auth
- **Infrastructure**: AWS, Terraform, Docker, Kubernetes (EKS)
- **CI/CD**: GitHub Actions

## 💻 Local Setup
1. Clone the repository
2. Copy the environment file: `cp .env.example .env`
3. Spin up the entire stack using Docker Compose:
   ```bash
   docker compose up --build
   ```
4. Access the frontend at `http://localhost:5173` and the backend at `http://localhost:8080`.

## 📚 Learning Resources
Head over to the `docs/` folder to explore guides on AWS, architecture, and cost control:
- [AWS Learning Guide](docs/aws-learning-guide.md)
- [Cost Control Guide](docs/cost-control.md)
- [Architecture](docs/architecture.md)

## 🧹 AWS Cleanup
If you deploy this to AWS, remember to clean up:
1. Delete EKS cluster (if manually created).
2. `terraform destroy`
3. Delete ECR images manually.
4. Empty S3 buckets manually so Terraform can destroy them.

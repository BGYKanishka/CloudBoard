.PHONY: dev dev-db build test lint docker-build docker-up docker-down clean db-migrate

# Start local development (frontend, backend, database)
dev:
	docker compose up -d postgres
	@echo "Waiting for postgres to start..."
	sleep 3
	@echo "Starting backend and frontend..."
	# In a real shell, you'd run these in parallel or different tabs
	# For simplicity, we suggest using docker compose for the full stack:
	# docker compose up

# Run backend tests
test:
	cd backend && go test ./...

# Run backend linter
lint:
	cd backend && go vet ./...
	cd frontend && npm run lint

# Build production binaries/bundles
build:
	cd backend && go build -o bin/server ./cmd/server
	cd frontend && npm run build

# Build Docker images
docker-build:
	docker build -t cloudboard-backend -f docker/backend.Dockerfile ./backend
	docker build -t cloudboard-frontend -f docker/frontend.Dockerfile ./frontend

# Start full Docker stack
docker-up:
	docker compose up --build -d

# Stop full Docker stack
docker-down:
	docker compose down

# Clean build artifacts
clean:
	rm -rf backend/bin/
	rm -rf frontend/dist/

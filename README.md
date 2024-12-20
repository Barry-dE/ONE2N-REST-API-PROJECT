# ONE2N-REST-API-PROJECT
This repository contains a Student CRUD REST API implemented in Golang using the Gin framework. The project serves as a learning exercise to explore best practices in building RESTful APIs, containerization, CI/CD pipelines, deployment, and observability. It's main purpose is to get hands-on experience with building and managing production workloads. 

##Technology Stack
Golang, Gin, PostgreSQL, Docker, Docker-Compose, GitHub Actions, Vagrant, Nginx, Kubernetes, Helm, ArgoCD, Prometheus, Grafana, Loki, Promtail, Postman.

## Features

### Functional Requirements

- Create a new student: Add a student record to the database.
- Retrieve all students: Fetch all student records.
- Retrieve a specific student: Fetch details of a student by ID.
- Update student information: Modify details of an existing student.
- Delete a student: Remove a student record from the database.
- Healthcheck endpoint: Monitor API health.

### Non-Functional Requirements

- API versioning: Supports versioning (e.g., api/v1/<resource>).
- Meaningful logging: Emits structured logs with appropriate log levels.
- Environment variables: Configurations are externalized and injected.
- Unit testing: Includes tests for all endpoints.
- Database migrations: Automates schema creation and updates.
- Postman collection: Pre-configured requests for API testing.

## Getting Started

### Prerequisites

- Install Docker
- Install Docker Compose
- Install GNU Make
- Install Vagrant
- Install kubectl
- Install Minikube

### Local Setup

Clone the repository:

`git clone https://github.com/<your-username>/student-crud-api.git`
`cd student-crud-api`

Run database migrations:

`make db-migrate`

Build and run the API locally:

`make run`

Build Docker Image

`make docker-build`

Run Docker Container

`make docker-run`

Environment variables can be injected at runtime using a .env file or passed directly during container execution.

### One-Click Local Setup

`make compose-up`

This ReadMe is still a work in progress. 



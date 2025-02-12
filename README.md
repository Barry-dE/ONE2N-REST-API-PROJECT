# REST-API-PROJECT

This repository contains a  CRUD REST API implemented in Golang using the Gin framework. The project serves as a learning exercise to explore best practices in building RESTful APIs, containerization, CI/CD pipelines, deployment, and observability. It's main purpose is to get hands-on experience with building and managing production workloads.

## Technology Stack
Golang, Gin, PostgreSQL, Docker, GitHub Actions, Vagrant, Nginx, Kubernetes, Helm, ArgoCD, Prometheus, Grafana, Loki, Promtail, Postman.

## Features

### Functional Requirements

- Create a new student: Add a student record to the database.
- Retrieve all students: Fetch all student records.
- Retrieve a specific student: Fetch details of a student by ID.
- Update student information: Modify details of an existing student.
- Delete a student: Remove a student record from the database.
- Healthcheck endpoint: Monitor API health.

### Non-Functional Requirements

- API versioning: Supports versioning (e.g., `api/v1/<resource>`).
- Meaningful logging: Emits structured logs with appropriate log levels.
- Environment variables: Configurations are externalized and injected.
- Unit testing: Includes tests for all endpoints.
- Database migrations: Automates schema creation and updates.
- Postman collection: Pre-configured requests for API testing.

## Getting Started

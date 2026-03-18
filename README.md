RealWorld Fullstack: Multi-Tier DevOps Deployment
## Overview

This project contains the **backend (Go API)**, **frontend (Angular)** and **db (Postgres)** of the RealWorld example application.  
The objective is to **containerize and deploy the application** using Docker, and document the full workflow.

**Architecture Overview**

The system is structured as a simple multi-tier setup.

* **Frontend (Presentation Layer):**
  An Angular single-page app served through Nginx, responsible for everything the user interacts with.

* **Backend (Application Layer):**
  A REST API built with Go (using the Gin framework) that handles the core logic and communication with the database.

* **Database (Data Layer):**
  A PostgreSQL 15 instance used for storing and managing all application data.

* **Networking:**
  All services communicate over a dedicated Docker network. The backend talks to the database internally (e.g., `db:5432`), while the database itself isn’t exposed to the public, which adds an extra layer of security.

 ## Folder Structure
```text
devops-assessment/
│
├─ backend/ # Go API repo
│ ├─ Dockerfile
│ ├─ hello.go
│ ├─ go.mod
│ └─ ...
│
├─ frontend/ # Angular repo
│ ├─ Dockerfile
│ ├─ package.json
│ ├─ nginx.conf
│ └─ ...
│
├─ .github\workflows/ # Workflow for Github Action
│ ├─ ci.yml
│
└─ docker-compose.yml
```

**Steps to Run Locally**
To run the services on local machine for debugging purposes:

> Note: Running locally on Windows may encounter Go/SQLite or Node dependency issues. Docker is recommended.

1. **Backend (Go)**
```bash
cd backend
go mod download
go run ./hello.go  # may fail on Windows due to SQLite/cgo
```

2. **Frontend (Angular)**
```bash
Navigate to /frontend (cd frontend)
Install dependencies: npm install --legacy-peer-deps
Run: npm start.
```
Access via http://localhost:4200.

**Steps to Run with Docker (Recommended)**

1. **Clone the Repository:**
```bash
git clone https://github.com/MoyoBamishe/devops-assessment.git
cd devops-assessment
```

2. **Launch the Stack:**
```bash
docker-compose up -d --build
```

3. **Verify Services:**
```bash
docker compose ps
```

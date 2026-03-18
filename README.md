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
**Built Images for both Frontend and Backend:**
<img width="1920" height="1080" alt="Screenshot (21)" src="https://github.com/user-attachments/assets/c274fbf4-25f0-4729-9d59-8bc0167efcce" />


## **Deployment Implementation:**

1. **Database Deployment:**

-The database is deployed via the db service in docker-compose.yml.
-I used Docker volumes to ensure database state is preserved across container restarts.
-Credentials are handled through environment variables, keeping things a bit safer and easier to manage.


2. **CI/CD Pipeline (GitHub Actions):**

I set up a CI pipeline in .github/workflows/ci.yml to automatically validate the project on every push to the master branch.
- It Builds and starts the full stack using docker compose up --build
- Confirms that the frontend, backend, and database all come up correctly and can communicate without issues
- This helps catch integration problems early and ensures the whole system is always in a working state.


3. **Nginx Configuration:**

I added a custom Nginx config to handle Angular routing properly. Instead of returning a 404 on page refresh or direct URL access, requests are redirected to index.html, allowing the app to manage routing.



**Public URLs & Server Details:**

**Frontend:** http://moyo.nqb8.co

**API/Backend:** http://api.moyo.nqb8.co

**Deployment IP:** 144.21.53.104


**Technical Decisions:**

**Service Health Checks:** Added a health check for Postgres and configured the backend to wait until the DB is ready before connecting. This avoids startup crashes caused by timing issues.

**Multi-Stage Dockerfiles:** Used multi-stage builds for both Go and Angular to keep the final images smaller and more secure by leaving out build tools and source files.

**Environment Variables:** Centralized the DATABASE_URL to make switching between local and container environments seamless.

**Nginx Setup:** Exposed the frontend on port 80 so the app is accessible over standard HTTP.












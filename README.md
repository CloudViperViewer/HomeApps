# 💰 MyBudgetApp Backend

This is the Go-based backend server for the MyBudgetApp project. It provides an API interface for querying financial data from a MariaDB database using dynamic filtering and pagination.

---

## 🚀 Features

- RESTful API via Gin (`/api/select`)
- MariaDB integration
- Reflection-based dynamic queries
- Dockerized multi-environment setup (dev, test, prod)

---

## 🧱 Tech Stack

- Go
- Gin (Web Framework)
- MariaDB (via Docker)
- Docker & Docker Compose

---

## 🐳 Docker Requirements

Before running the backend, ensure you have the following installed:

### 📦 Installations

#### 1. Docker

Install Docker from:  
[https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)

#### 2. Docker Compose

Docker Compose is usually bundled with Docker Desktop. To verify:

```bash
docker-compose --version
```

---

## 🛠️ Running the Project

### 1. Clone the Repository

```bash
git clone https://github.com/your-org/mybudgetapp-backend.git
cd mybudgetapp-backend
```

### 2. Start Containers

```bash
docker-compose up --build
```

This will start:

- `go_api_dev`: The Go API server
- `db_dev`: The MariaDB development database

### 3. Verify

Once containers are up, the API should be available at:  
[http://localhost:8081/api/select](http://localhost:8081/api/select)

---

## 🧪 API Testing

Test the `/api/select` endpoint using Postman or curl.  
Refer to [`api_select_endpoint.md`](./api_select_endpoint.md) for full schema and examples.

---

## 📂 Project Structure

```
.
├── api.go                # Gin router setup
├── select.go             # /api/select endpoint logic
├── database.go           # DB connection handling
├── util.go               # DB query builders and helpers
├── tables/               # Structs and logic per table
├── Dockerfile
├── docker-compose.yml
├── go.mod / go.sum
└── README.md
```

---

## ⚙️ Environment Variables

Defined in `docker-compose.yml` per environment:

- `DB_USER`
- `DB_PASS`
- `DB_HOST`
- `DB_NAME`
- `DB_PORT` (optional)

---

## 📄 License

MIT (or your preferred license)

---

Happy hacking! 🚀
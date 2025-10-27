# 🚀 Go REST API Example with Gin & Gorm

A simple example project showing how to build a REST API in **Go** using **Gin**, **Gorm**, and configuration files (`config.yaml` and `.env`).

---

## 📁 Project Structure

go-restapi-example-with-gin/
├── config/ # Configuration files 
│ ├── config.yaml
│ └── config.go
├── internal/ # Core business logic
│ ├── handler/
│ ├── routes/
│ ├── database/
│ ├── utils/
│ └── model/
├── main.go # Entry point
├── go.mod
├── go.sum
├── .gitignore
└── README.md
---

## ⚙️ Example `config.yaml`

```yaml
server:
  port: 8080

database:
  host: localhost
  port: 5432
  user: postgres
  password: secret_password
  name: example_db
  sslmode: disable
```

🔐 Example .env
env

# Application

DB_PASSWORD="your_password"

🧩 Technologies Used
Go 1.25.3

Gin — HTTP web framework

Gorm — ORM for database access

PostgreSQL

YAML / env

🚀 Getting Started

# 1. Clone the repository
git clone https://github.com/Dima-Melnik/go-restapi-example-with-gin.git
cd go-restapi-example-with-gin

# 3. Install dependencies
go mod tidy

# 4. Run the application
go run main.go

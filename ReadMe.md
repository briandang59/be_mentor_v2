# Mentors Platform

Dự án xây dựng hệ thống web **Mentor/Mentee** sử dụng **Go (Gin)**, **Postgres**, **Docker**, hỗ trợ **JWT Authentication**, **WebSocket realtime** và **modular architecture** theo Clean Architecture.

---

## 🚀 Công nghệ sử dụng
- [Go](https://golang.org/) + [Gin](https://github.com/gin-gonic/gin) — Web framework
- [GORM](https://gorm.io/) — ORM cho Postgres
- [Postgres](https://www.postgresql.org/) — CSDL chính
- [Docker](https://www.docker.com/) + Docker Compose — Dev & deploy
- [JWT](https://jwt.io/) — Authentication
- WebSocket — Realtime chat, presence

---

## 📂 Cấu trúc dự án

```bash
mentors/
│── cmd/                  # Entry point cho app (main.go)
│   └── server/
│       └── main.go
│
├── config/               # Cấu hình (env, config loader, logging, constants)
│   └── config.go
│
├── internal/             # Business logic (chỉ dùng cho project này)
│   ├── app/              # Application core (use cases)
│   │   ├── user/         # Module User (domain)
│   │   │   ├── controller.go
│   │   │   ├── service.go
│   │   │   ├── repository.go
│   │   │   └── model.go
│   │   ├── guild/        # Module Guild/Server
│   │   ├── channel/      # Module Channel
│   │   ├── message/      # Module Message
│   │   ├── auth/         # Module Authentication & JWT
│   │   └── voice/        # Module Voice (WebRTC/WebSocket)
│   │
│   ├── routes/           # HTTP & WebSocket route definitions
│   │   └── routes.go
│   │
│   ├── middleware/       # Gin middleware (Auth, Logging, Rate limiting…)
│   │   └── auth_middleware.go
│   │
│   ├── ws/               # WebSocket server (realtime chat, presence)
│   ├── jobs/             # Background jobs, worker queues
│   └── utils/            # Helper functions (hash, validation, UUID…)
│
├── pkg/                  # Shared library có thể tái sử dụng
│   ├── database/         # DB connection (Postgres, Redis…)
│   ├── cache/            # Redis cache
│   ├── logger/           # Logging system
│   ├── storage/          # File storage (S3, local, GCP…)
│   └── websocket/        # Abstraction layer cho WS
│
├── migrations/           # File SQL migrate cho Postgres
├── api/                  # OpenAPI/Swagger docs
├── tests/                # Integration & unit tests
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum

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
│   │   │   ├── controller.go
│   │   │   ├── service.go
│   │   │   ├── repository.go
│   │   │   └── model.go
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
│   │   └── hub.go
│   │
│   ├── jobs/             # Background jobs, worker queues
│   │   └── email_sender.go
│   │
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
│   └── 001_init.up.sql
│
├── api/                  # OpenAPI/Swagger docs
│   └── swagger.yaml
│
├── tests/                # Integration & unit tests
│   ├── user_test.go
│   └── message_test.go
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum

# trainers-manager

REST API for AI-powered training plan generation. Accepts a training group, picks the next accent and skill in the configured cycle, calls an LLM to build a structured workout plan, and persists everything to PostgreSQL — asynchronously, with full task tracking.

## Stack

- **Go** + [Fiber v2](https://gofiber.io)
- **PostgreSQL 16** via `pgx`
- **LLM** — any OpenAI-compatible endpoint (default: Ollama `llama3.1:8b`)
- **golang-migrate** for schema migrations
- **Zerolog** for structured logging
- **Task** as the task runner

## Architecture

Clean Architecture — dependencies flow strictly inward:

```
HTTP (Fiber)
    └── controller/restapi/v1      ← parse & validate requests
            └── usecase/training   ← business logic, orchestration
                    └── repo/
                        ├── persistent/   ← PostgreSQL (pgx)
                        └── webapi/       ← LLM API client
```

All layer boundaries are interfaces (`internal/usecase/contracts.go`, `internal/repo/contracts.go`). Mocks are generated via `go.uber.org/mock`.

## Quick Start

**Prerequisites:** Go 1.24+, Docker, [Task](https://taskfile.dev), [golang-migrate](https://github.com/golang-migrate/migrate)

```bash
# 1. Start Postgres
task pg-up

# 2. Apply migrations
task mig-up

# 3. Start the server
go run ./cmd/app/ -config config/local.yaml
```

Server starts on `localhost:9045`.

For LLM generation, Ollama must be running locally:

```bash
ollama pull llama3.1:8b
ollama serve
```

## Configuration

`config/local.yaml`:

```yaml
env: local

postgre:
  host: localhost
  port: 5432
  user: postgres
  password: postgres
  name: training
  ssl: disable

server:
  host: localhost
  port: 9045
  read_timeout: 5      # seconds
  write_timeout: 5
  shutdown_timeout: 3

llm:
  base_url: "http://localhost:11434/v1"   # any OpenAI-compatible URL
  api_key: "ollama"
  model: "llama3.1:8b"

logging_handler:
  level: debug
  format: json
```

To use a different LLM provider (OpenAI, vLLM, etc.) — change `llm.base_url`, `llm.api_key`, and `llm.model`.

## API

Base path: `http://localhost:9045`

### Health

| Method | Path      | Description  |
|--------|-----------|--------------|
| GET    | `/health` | Liveness check |

### Structures — workout templates

| Method | Path                         | Description      |
|--------|------------------------------|------------------|
| POST   | `/v1/training/structure`     | Create structure |
| GET    | `/v1/training/structure`     | List all         |
| GET    | `/v1/training/structure/:id` | Get by ID        |
| PATCH  | `/v1/training/structure/:id` | Update           |
| DELETE | `/v1/training/structure/:id` | Delete           |

### Exercises

| Method | Path                        | Description     |
|--------|-----------------------------|-----------------|
| POST   | `/v1/training/exercise`     | Create exercise |
| GET    | `/v1/training/exercise`     | List all        |
| PATCH  | `/v1/training/exercise/:id` | Update          |
| DELETE | `/v1/training/exercise/:id` | Delete          |

### Training Groups

Groups define the rotation cycles used during generation (`accent_cycle`, `skill_cycle`).

| Method | Path                     | Description  |
|--------|--------------------------|--------------|
| POST   | `/v1/training/group`     | Create group |
| GET    | `/v1/training/group`     | List all     |
| PATCH  | `/v1/training/group/:id` | Update       |
| DELETE | `/v1/training/group/:id` | Delete       |

### Training Plans

| Method | Path                    | Description |
|--------|-------------------------|-------------|
| GET    | `/v1/training/plan`     | List all    |
| GET    | `/v1/training/plan/:id` | Get by ID   |
| PATCH  | `/v1/training/plan/:id` | Update      |

### Plan Generation (async)

| Method | Path                        | Description                    |
|--------|-----------------------------|--------------------------------|
| POST   | `/v1/training/generate`     | Start generation → `{task_id}` |
| GET    | `/v1/training/generate/:id` | Poll task status               |

`POST /v1/training/generate` returns HTTP 202 immediately with a `task_id`. Poll `GET /v1/training/generate/:id` until status is `DONE` or `ERROR`.

Task statuses: `CREATED → SENT → PROCESSING → DONE | ERROR`

## Generation Flow

```
POST /v1/training/generate
    │
    ├── 1. Return {task_id} immediately (HTTP 202)
    │
    └── background goroutine:
        │
        ├── 2. Fetch training group → determine next accent & skill from cycle
        ├── 3. Load structure template + exercise pool from DB
        ├── 4. Call LLM with prompt (structure + exercises as context)
        ├── 5. Validate LLM response — check exercise IDs exist in pool
        ├── 6. Write Training + TrainingPlan + exercises to DB
        └── 7. Update task status → DONE (or ERROR on failure)
```

The accent and skill are selected by cycling through `training_group.accent_cycle` and `training_group.skill_cycle` arrays, advancing on each generation for that group.

## Database Schema

```
training_structure      — workout templates (structure TEXT)
exercises               — exercise library (muscle, position, description)
training               — training sessions
training_exercises     — training ↔ exercises join table
training_group         — groups with accent_cycle[], skill_cycle[]
training_plan          — generated plans linked to training + group + structure
training_plan_history  — audit log, partitioned by quarter (JSONB snapshots)
generation_tasks       — async task tracking (status, error)
```

`training_plan_history` is a partitioned table. Partitions for the current and next quarter are created automatically on startup via `internal/app/partition.go`.

## Task Commands

```bash
task pg-up                    # Start Postgres in Docker on :5432
task pg-down                  # Stop and remove Postgres
task pg-logs                  # Tail Postgres logs

task mig-up                   # Apply all pending migrations
task mig-down                 # Roll back last migration
task mig-version              # Show current migration version
task mig-force VERSION=<n>    # Force migration version (fix dirty state)
task mig-create NAME=<name>   # Create a new migration pair

task mock-gen                 # Regenerate mocks (go.uber.org/mock)
```

## Running Tests

```bash
go test ./...
go test ./internal/usecase/training/... -run TestNextInCycle
```

## Project Structure

```
trainers-manager/
├── cmd/app/               — entry point
├── config/                — config files (YAML)
├── internal/
│   ├── app/               — wiring: migrations → DB → repos → usecases → HTTP
│   ├── config/            — config loader
│   ├── controller/
│   │   └── restapi/v1/    — Fiber handlers, request/response DTOs
│   ├── entity/            — domain structs
│   ├── repo/
│   │   ├── contracts.go   — repository interfaces
│   │   ├── persistent/    — PostgreSQL implementations + raw SQL queries
│   │   ├── webapi/        — LLM API client
│   │   └── workers/       — worker-layer types
│   └── usecase/
│       ├── contracts.go   — use case interfaces
│       └── training/      — business logic implementations
├── migrations/            — SQL migration files (golang-migrate)
└── pkg/
    ├── httpserver/        — Fiber wrapper with graceful shutdown
    ├── logger/            — Zerolog wrapper
    ├── postgres/          — pgx pool setup
    └── workers/           — generation event worker
```

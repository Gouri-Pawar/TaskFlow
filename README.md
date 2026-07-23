# TaskFlow

A full-stack task management web app with JWT authentication, priorities, due dates, and daily completion streaks — built with Go, Gin-adjacent standard `net/http`, GORM, and PostgreSQL on the backend, and a lightweight vanilla HTML/CSS/JS frontend.

> Repo: [github.com/Gouri-Pawar/TaskFlow](https://github.com/Gouri-Pawar/TaskFlow)

---

## Problem

Most personal to-do lists fail for one of two reasons:

- **They're too simple** — a flat list with no sense of priority, deadline, or ownership, so everything looks equally urgent (or equally ignorable).
- **They're too heavy** — full project-management tools (boards, sprints, tags, assignees) are overkill for someone who just wants to track their own day-to-day tasks.

There's also a motivation gap: most task apps show you *what's left to do*, but not *how consistently you've been showing up*. Nothing reinforces the habit of actually finishing things.

## Solution

TaskFlow is a focused, single-user task tracker that:

- Keeps each user's tasks private and secure behind token-based authentication.
- Lets tasks carry a **priority** (low / medium / high) and an optional **due date**, so the list can be scanned and triaged at a glance.
- Surfaces a **daily streak** (🔥) computed from completion history, turning task completion into a habit-forming feedback loop rather than a chore.
- Ships as a clean split-view login/register experience and a single-page dashboard — no framework overhead, no build step, just a Go binary and static files.

## Key Features

- 🔐 **Secure auth** — registration and login with `bcrypt`-hashed passwords and stateless **JWT** sessions (24-hour expiry).
- ✅ **Full task CRUD** — create, read, update, and delete tasks, scoped per authenticated user.
- 🚦 **Priorities** — every task is tagged low / medium / high, with an automatic fallback to `medium` for missing or invalid values.
- 📅 **Due dates** — optional due dates, with overdue tasks flagged in the UI.
- 🔥 **Completion streaks** — a "current streak" counter and a 14-day activity dot-row, computed client-side from each task's `completed_at` timestamp.
- 🗂️ **Filtering & search** — instantly filter by All / Pending / Completed, and search across task titles and descriptions.
- 🎨 **Sticky-note dashboard UI** — a responsive, collapsible-sidebar dashboard styled like a wall of sticky notes, with an add/edit modal.
- 🌐 **CORS-enabled REST API** — a clean JSON API that any frontend (or mobile client) could consume.

## Effectiveness / What It Demonstrates

This project was built as a hands-on way to practice production-style backend patterns in Go, specifically:

- Designing a **REST API from scratch** using Go's standard `net/http` (no framework), including custom middleware chaining for CORS and JWT verification.
- Implementing **stateless authentication** end-to-end: password hashing, token issuance, token verification, and per-request user-context propagation via `context.WithValue`.
- Using **GORM** with PostgreSQL for schema auto-migration, scoped queries (`WHERE user_id = ?`), and struct-tag-driven JSON serialization.
- Writing defensive handler logic — e.g. normalizing invalid `priority` values instead of rejecting the request, and only touching `CompletedAt` when the completion state actually changes (so edits to a task don't corrupt streak data).
- Building a functional frontend without any framework: state management, optimistic UI updates, and derived data (like the streak) computed entirely in the browser from raw API data.

## Tech Stack

**Backend**
- [Go](https://go.dev/) 1.26
- [`net/http`](https://pkg.go.dev/net/http) — routing and server
- [GORM](https://gorm.io/) + [`gorm.io/driver/postgres`](https://gorm.io/driver/postgres) — ORM and PostgreSQL driver
- [`golang-jwt/jwt/v5`](https://github.com/golang-jwt/jwt) — JWT issuing and verification
- [`golang.org/x/crypto/bcrypt`](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — password hashing
- [`joho/godotenv`](https://github.com/joho/godotenv) — environment variable loading

**Database**
- PostgreSQL

**Frontend**
- Vanilla HTML, CSS, and JavaScript (no build tooling, no framework)

## Project Structure

```
TaskFlow/
├── cmd/
│   └── server/
│       └── main.go          # Entry point — routes, middleware, server bootstrap
├── config/
│   └── database.go          # PostgreSQL connection + .env loading
├── handlers/
│   ├── auth_handler.go      # Register / Login
│   └── task_handler.go      # Create / Get / Update / Delete tasks
├── middleware/
│   └── jwt_middleware.go    # JWT verification, injects user_id into context
├── models/
│   ├── user.go               # User model (GORM)
│   └── task.go               # Task model (GORM)
├── frontend/
│   ├── index.html            # Landing page
│   ├── login.html            # Login page
│   ├── register.html         # Registration page
│   ├── dashboard.html        # Main app (task board, streaks, filters, modal)
│   └── css/style.css
├── go.mod / go.sum
└── .gitignore
```

## API Reference

All task routes require an `Authorization: Bearer <token>` header.

| Method | Endpoint          | Description                          | Auth required |
|--------|-------------------|---------------------------------------|:---:|
| POST   | `/register`        | Create a new user account             | ❌ |
| POST   | `/login`            | Authenticate and receive a JWT        | ❌ |
| GET    | `/get_tasks`        | Get all tasks for the logged-in user  | ✅ |
| POST   | `/tasks`            | Create a new task                     | ✅ |
| PUT    | `/tasks/update?id=` | Update an existing task               | ✅ |
| DELETE | `/tasks/delete?id=` | Delete a task                         | ✅ |

**Task object shape:**

```json
{
  "title": "string",
  "description": "string",
  "priority": "low | medium | high",
  "due_date": "RFC3339 timestamp or null",
  "completed": true,
  "completed_at": "RFC3339 timestamp or null",
  "user_id": 1
}
```

## Getting Started

### Prerequisites
- Go 1.26+
- PostgreSQL running locally (or a connection string to a remote instance)

### 1. Clone the repo
```bash
git clone https://github.com/Gouri-Pawar/TaskFlow.git
cd TaskFlow
```

### 2. Configure environment variables
Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_USER=your_pg_user
DB_PASSWORD=your_pg_password
DB_NAME=taskflow
DB_PORT=5432
JWT_SECRET=your_jwt_secret
```

### 3. Install dependencies
```bash
go mod tidy
```

### 4. Run the server
```bash
go run cmd/server/main.go
```
The API starts on `http://localhost:8080` and auto-migrates the `User` and `Task` tables on startup.

### 5. Open the frontend
Serve the `frontend/` folder with any static file server (e.g. the VS Code Live Server extension, or `python -m http.server`) and open `index.html` in your browser. The frontend calls the API at `http://localhost:8080` by default — update the `API_BASE` constant in `dashboard.html`, `login.html`, and `register.html` if you deploy the backend elsewhere.

## Roadmap / Possible Improvements

- Add automated tests (unit tests for handlers, integration tests for the API)
- Move inline dashboard/login/register JavaScript into the existing `frontend/js/` files for better separation of concerns
- Add task categories/tags and sorting options
- Add refresh tokens / logout-everywhere support
- Containerize with Docker Compose (app + Postgres) for one-command setup
- Deploy a live demo

## Author

Built by [Gouri Pawar](https://github.com/Gouri-Pawar).

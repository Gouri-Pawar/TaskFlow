# TaskFlow

A full-stack task management web app with JWT authentication, priorities, due dates, and daily completion streaks — built with Go, Gin-adjacent standard `net/http`, GORM, and PostgreSQL on the backend, and a lightweight vanilla HTML/CSS/JS frontend.

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

- **Secure auth** — registration and login with `bcrypt`-hashed passwords and stateless **JWT** sessions (24-hour expiry).
- **Full task CRUD** — create, read, update, and delete tasks, scoped per authenticated user.
- **Priorities** — every task is tagged low / medium / high, with an automatic fallback to `medium` for missing or invalid values.
- **Due dates** — optional due dates, with overdue tasks flagged in the UI.
- **Completion streaks** — a "current streak" counter and a 14-day activity dot-row, computed client-side from each task's `completed_at` timestamp.
- **Filtering & search** — instantly filter by All / Pending / Completed, and search across task titles and descriptions.
- **Sticky-note dashboard UI** — a responsive, collapsible-sidebar dashboard styled like a wall of sticky notes, with an add/edit modal.
- **CORS-enabled REST API** — a clean JSON API that any frontend (or mobile client) could consume.
  

## Screenshots

### Landing Page
<img width="1920" height="1014" alt="Screenshot 2026-07-24 141344" src="https://github.com/user-attachments/assets/43e98690-b315-400d-9e30-4e20d97742d0" />

### Login Page
<img width="1920" height="1011" alt="Screenshot 2026-07-24 141437" src="https://github.com/user-attachments/assets/dc0d0336-5bc5-4d1f-b812-48be340abaa2" />

### Dashboard
<img width="1920" height="1017" alt="Screenshot (58)" src="https://github.com/user-attachments/assets/b23fc199-48f8-4332-b0a8-2ed9b5d3b8d1" />

### Task Management
<img width="1920" height="1015" alt="Screenshot (57)" src="https://github.com/user-attachments/assets/39e2b385-5d7e-4f28-8af1-8875c192d36d" />


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
- Vanilla HTML
- CSS
- JavaScript 

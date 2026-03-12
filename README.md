# Cigarland: Go, Databases, and Good Friends

Welcome to **Cigarland**-a project built by two friends learning **Go**, diving into **databases**, and logging our journey through code, cigars, and craftsmanship.

This isn't just about syntax or schema-it's about sharpening our tools, improving our systems, and enjoying a damn fine cigar while we do it.

---

## Tech Stack

| Layer         | Tooling                         |
|---------------|---------------------------------|
| Language      | Go 1.25                         |
| Web Framework | Gin                             |
| ORM           | GORM                            |
| Database      | PostgreSQL 17                   |
| Frontend      | HTML / CSS / Vanilla JS         |
| Infra         | Docker Compose, NGINX           |
| Auth          | Google OAuth2 + session cookies |

---

## Features

- [x] Google OAuth2 login
- [x] Per-user permission system (add, delete, admin)
- [x] Add cigars with brand, wrapper, profile, strength, flavor notes
- [x] Per-user ratings and reviews (Kyle + John)
- [x] Filtered cigar queries via dynamic WHERE API
- [x] Delete cigars
- [x] Admin panel to manage user permissions
- [x] Containerized with Docker Compose + NGINX reverse proxy
- [ ] Search, sort, and filter UI
- [ ] Save shared tasting notes with `@jhohnharmon`
- [ ] Humidor inventory tracking
- [ ] Future: track humidor inventory, humidity/temp over time

---

## Running the App

### With Docker Compose

Create a `.env` file in the project root:

```env
GOOGLE_CLIENT_ID=...
GOOGLE_CLIENT_SECRET=...
GOOGLE_REDIRECT_URL=http://localhost/auth/google/callback
```

Then:

```bash
docker compose up --build
```

The app will be available at `http://localhost`.

### Locally

```bash
cd go_app

# First-time setup: create tables and seed admin user
go run . --create-db

# Start the server
go run . --start-server --port 8080
```

Set DB connection via env vars (`DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`) or pass `--db-dsn` directly.

---

## Auth & Permissions

Login is handled via Google OAuth2. On first login, every user is registered with no permissions. The admin (`kyle@15kmr.com`) is seeded with full access on `--create-db`.

| Permission   | Controls                        |
|--------------|---------------------------------|
| `can_add`    | Create new cigars               |
| `can_delete` | Delete cigars                   |
| `can_admin`  | View and edit user permissions  |

---

## Cigar Review Sample

```json
{
  "origin": "Cuba",
  "brand": "Arturo Fuente",
  "name": "Hemingway Short Story",
  "wrapper": "Cameroon",
  "profile": "Medium",
  "tasty_tip": false,
  "pressed": false,
  "binder": "string",
  "spicy": 7,
  "kyle_rating": 4,
  "john_rating": 6,
  "length": 60,
  "ring": 50,
  "review": "Smooth draw, nutty on the front, spice at the end. Perfect burn. Smoked with JHarmon while talking Go interfaces.",
  "john_review": "",
  "kyle_review": "",
  "authentic_human_review": ""
}
```

---

## Learning Topics

This repo explores real-world fundamentals while staying fun:

- Go modules, structs, interfaces
- HTTP servers, routers, middleware
- SQL design, querying, schema evolution
- RESTful API patterns
- OAuth2 and session-based auth
- Testing and validation
- Logging and error handling
- Docker Compose and containerization
- Hosting, reverse proxying (NGINX), and deployment basics
- Too much html and javascript :(

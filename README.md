# 🥃 Cigarland: Go, Databases, and Good Friends

Welcome to **Cigarland**—a project built by two friends learning **Go**, diving into **databases**, and logging our journey through code, cigars, and craftsmanship.

This isn’t just about syntax or schema—it’s about sharpening our tools, improving our systems, and enjoying a damn fine cigar while we do it.

---

## 🚀 Project Goals

- **Learn Go (Golang)**: idiomatic practices, concurrency, API building
- **Master Databases**: SQL/NoSQL, ORM vs raw queries, performance patterns
- **Build a Full App**: local + web-based cigar review tracker
- **Experiment & Deploy**: play with containers, systemd, CI/CD
- **Have Fun**: learn with intent, but never take it too seriously

---

## 🧱 Tech Stack

| Layer        | Tooling                  |
|--------------|--------------------------|
| Language     | Go (Golang)              |
| Web Framework| `net/http`, `Gin`, `Fiber` (TBD) |
| Database     | SQLite (starter), PostgreSQL (target maybe?????????) |
| Frontend     | Minimal HTML/CSS |
| Infra        | systemd, docker (future), NGINX |
| Dev Workflow | Git, Makefile |

---

## 🔥 Feature Concepts

- [x] Add cigars with brand, wrapper, country, strength, flavor notes
- [] Leave 1–10 star reviews and tasting notes
- [ ] Search, sort, and filter cigars by flavor profile or score
- [ ] Save shared tasting notes with `@jhohnharmon`
- [ ] Deployable ANYWHERE
- [ ] Future: track humidor inventory, humidity/temp over time

---

## 🧠 Learning Topics

This repo explores real-world fundamentals while staying fun:

- Go modules, structs, interfaces
- HTTP servers, routers, middleware
- SQL design, querying, schema evolution
- RESTful API patterns
- Testing and validation
- Logging and error handling
- systemd services and environment configuration
- Hosting, reverse proxying (NGINX), and deployment basics

---

## 🍂 Cigar Review Sample

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

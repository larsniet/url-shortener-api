# URL Shortener

## 🏗️ Architecture Overview

| Component               | Choice                              | Why?                                                                  |
| ----------------------- | ----------------------------------- | --------------------------------------------------------------------- |
| **Folder Structure**    | `cmd/`, `internal/`, `pkg/`         | Follows Go standards (`cmd/` for entry, `internal/` for private code) |
| **Domain-Based Layout** | `internal/url/`, `internal/health/` | Groups logic by feature, not by technical layer                       |
| **Layered Design**      | Handler → Service → Repository      | Separates HTTP, business logic, and data access cleanly               |

---

## 💾 Database Handling

| Component              | Choice                     | Why?                                                     |
| ---------------------- | -------------------------- | -------------------------------------------------------- |
| **Database Driver**    | `lib/pq` or `pgx`          | Standard, reliable PostgreSQL drivers                    |
| **Raw SQL**            | No ORM, use `database/sql` | Simple, idiomatic, testable, avoids abstraction overhead |
| **Repository Pattern** | Interface for data access  | Enables testability (mocking), decouples service & DB    |
| **Slug Generation**    | In service layer           | Business logic stays out of DB layer                     |
| **Migrations**         | Atlas (with `atlas.hcl`)   | Versioned schema, repeatable deployments                 |

---

## 🛠️ Key Technologies

| Tool/Library        | Purpose                   | Why Chosen?                             |
| ------------------- | ------------------------- | --------------------------------------- |
| **Chi Router**      | HTTP routing              | Lightweight, idiomatic Go router        |
| **Swagger/OpenAPI** | API documentation         | Clear, interactive API docs             |
| **Docker**          | Containerization          | Consistent local and prod environments  |
| **GitHub Actions**  | CI/CD pipeline            | Automated builds, tests, multi-platform |
| **Air**             | Hot reload in development | Speeds up development                   |

---

## ⚙️ Design Patterns

| Pattern                  | Purpose                           | Why?                                    |
| ------------------------ | --------------------------------- | --------------------------------------- |
| **Dependency Injection** | Inject repo into service          | Decoupling, testability                 |
| **Repository Interface** | Abstract data storage             | Swappable DB, easier unit testing       |
| **Environment Config**   | `.env` + 12-Factor app principles | Flexible, environment-specific settings |
| **Multi-stage Docker**   | Separate build/runtime containers | Smaller, more secure production images  |

---

## 🧪 Testing Strategy

| Layer      | Test Type                   | Why?                               |
| ---------- | --------------------------- | ---------------------------------- |
| Repository | Mock DB or integration      | Validate SQL and data handling     |
| Service    | Pure logic unit tests       | Ensures business rules are correct |
| Handler    | HTTP request/response tests | Ensures correct API behavior       |

---

## 🌐 API Design

| Principle              | Why?                                         |
| ---------------------- | -------------------------------------------- |
| **RESTful Routes**     | Standard, predictable API usage              |
| **OpenAPI Spec**       | Auto-generated docs, easy for clients to use |
| **Health Check Route** | Enables observability, uptime monitoring     |

---

## 🚢 Deployment Strategy

| Choice                    | Why?                                             |
| ------------------------- | ------------------------------------------------ |
| **Docker Compose**        | Dev environment management                       |
| **GHCR + GitHub Actions** | Automated builds/releases, secure image registry |
| **Environment Variables** | Clean config, no hardcoded secrets               |

---

## 📊 Summary

This architecture balances **simplicity**, **testability**, and **scalability** by:

- Following **Go idioms** and **standard folder structures**
- Keeping logic **modular and decoupled**
- Using **raw SQL** with PostgreSQL for **maximum control and performance**
- Leveraging **containers** and **CI/CD** for seamless deployment

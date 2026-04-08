# Go Microservice API

A Go REST API demonstrating modular service boundaries (users, products, orders, and related commerce features) with JWT auth and MongoDB.

---

## Jira — Description

Build a microservice-oriented API in Go that separates concerns into clear domains (identity, catalog, commerce, and supporting services). The implementation uses a single deployable binary with route groups and packages that mirror how independent microservices would be split (users, products, orders, cart, notifications, etc.), Gorilla Mux for routing, MongoDB for persistence, and JWT for protected routes. The system includes operational endpoints (health, version, metrics), middleware (logging, recovery, CORS, rate limiting, request IDs), and environment-based configuration suitable for staging and production.

---

## Jira — Acceptance Criteria

- [ ] Implement RESTful API endpoints using Gorilla Mux (including path parameters and protected subroutes).
- [ ] Provide user management: registration, login, JWT issuance, and CRUD-style user APIs (protected).
- [ ] Provide product management: list, create, get, update, delete products (protected).
- [ ] Provide order processing: list/create orders, get order, update order status, cancel order for the authenticated user (protected).
- [ ] Add JWT authentication middleware and secure `/api/*` routes that require a valid `Authorization: Bearer` token.
- [ ] Integrate MongoDB for data storage with configurable URI and database name.
- [ ] Implement health check and operational endpoints: `GET /api/health`, `GET /api/version`, `GET /api/metrics`.
- [ ] Add structured logging, panic recovery, CORS, optional rate limiting, and `X-Request-ID` propagation.
- [ ] Document the API surface (this README), environment variables, and primary routes for consumers.

---

## Technical Stack

| Layer | Technology |
|--------|------------|
| Language | Go 1.21+ |
| HTTP router | Gorilla Mux |
| Database | MongoDB (official driver) |
| Authentication | JWT (golang-jwt/jwt/v5) |
| Password hashing | golang.org/x/crypto/bcrypt |
| Configuration | Environment variables + `godotenv` (optional `.env`) |

---

## Architecture Notes

- **Modular packages**: `internal/handlers`, `internal/models`, `internal/middleware`, `internal/auth`, `internal/config`, `internal/services`, `internal/repository`, etc.
- **Auth context**: JWT claims are attached to the request context using typed keys (`user_id`, `email`, `role`).
- **Public vs protected**: Registration, login, health, version, metrics, and coupon validation are public; most business routes live under `/api` with JWT + rate limiting.

---

## Environment Variables

| Variable | Purpose | Default |
|----------|---------|---------|
| `PORT` | HTTP listen port | `8080` |
| `MONGO_URI` | MongoDB connection string | `mongodb://localhost:27017` |
| `MONGO_DB` | Database name | `microservice_db` |
| `JWT_SECRET` | HMAC secret for signing tokens | `your-secret-key` (change in production) |
| `CORS_ORIGIN` | `Access-Control-Allow-Origin` | `*` |
| `APP_ENV` | Environment label | `development` |

---

## Installation

```bash
go mod download
go run main.go
```

---

## API Overview

### Public

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/health` | Liveness / health |
| GET | `/api/version` | Build version metadata |
| GET | `/api/metrics` | Basic request/error counters |
| POST | `/api/auth/register` | Register |
| POST | `/api/auth/login` | Login (returns JWT) |
| GET | `/api/coupons/validate?code=` | Validate coupon code |

### Protected (`Authorization: Bearer <token>`)

| Area | Examples |
|------|----------|
| Users | `GET/PUT/DELETE /api/users`, `GET /api/users/{id}` |
| Products | `GET/POST /api/products`, `GET/PUT/DELETE /api/products/{id}` |
| Reviews | `GET/POST /api/products/{product_id}/reviews` |
| Categories | `GET/POST /api/categories`, `GET/DELETE /api/categories/{id}` |
| Cart | `GET /api/cart`, `PUT /api/cart/lines` |
| Addresses | `GET/POST /api/addresses`, `DELETE /api/addresses/{id}` |
| Coupons | `POST /api/coupons`, `DELETE /api/coupons/{id}` |
| Notifications | `GET/POST /api/notifications`, `PATCH /api/notifications/{id}/read` |
| Wishlist | `GET/POST /api/wishlist`, `DELETE /api/wishlist/{id}` |
| Orders | `GET/POST /api/orders`, `GET /api/orders/{id}`, `PATCH /api/orders/{id}/status`, `POST /api/orders/{id}/cancel` |

---

## Error Handling & Logging

- Handlers return appropriate HTTP status codes; middleware recovers from panics and returns `500` without crashing the process.
- Request logging records method, path, and duration; metrics track aggregate request and error counts.

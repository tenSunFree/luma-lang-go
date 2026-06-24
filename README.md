# luma-lang-go

![CI](https://github.com/tenSunFree/rest-boilerplate-refined-go/actions/workflows/ci.yml/badge.svg)

---

## Introduction

If you're interested in Go (Gin / PostgreSQL / sqlx / Redis / JWT / Clean Architecture / Docker), feel free to take a look.

go-rest-boilerplate  
https://github.com/snykk/go-rest-boilerplate

This project is for learning and technical practice.

It can also be paired with my Kotlin Multiplatform app project to demonstrate a full-stack mobile application architecture.

---

## Related App Client

This backend can be used together with my Kotlin Multiplatform app project:

- [luma-lang-kmp](https://github.com/tenSunFree/luma-lang-kmp)

The app project provides a cross-platform client foundation built with Kotlin Multiplatform, Compose Multiplatform, Clean Architecture, MVI ViewModel, Navigation3, and Koin.

It can serve as the Android / iOS client-side foundation for connecting to RESTful APIs, handling authentication flows, and demonstrating modern shared mobile architecture.

---

## Preview

<p align="left">
  <img src="https://i.postimg.cc/597mrH0y/2026-06-23-071232.png" width="500"/>
  <img src="https://i.postimg.cc/pVGBS9Tx/2026-06-23-071318.png" width="500"/>
  <img src="https://i.postimg.cc/HWZtNJxk/2026-06-23-071331.png" width="500"/>
  <img src="https://i.postimg.cc/Kcq5WKzm/2026-06-23-071343.png" width="500"/>
</p> 

---

## Features

### Authentication & Account Security
- JWT-based login with access/refresh token support and refresh token rotation
- Account registration with email-based OTP verification before activation
- Forgot and reset password flow using time-limited reset tokens delivered via async email dispatch
- Change password endpoint with old-password verification
- Audit logging for auth-sensitive actions

### Live Course Sessions
- Real-time live class support powered by Agora RTC, with per-user UID allocation and RTC token generation on join
- Live join response includes teacher camera/screen stream UIDs and WebSocket chat configuration
- Live course listing filterable by status (`scheduled`, `live`, `ended`, `cancelled`)
- Per-user reminder state reflected in the live course list response

### Content Delivery
- Language lesson/content listing, full-text search, and detail retrieval
- Two-layer caching with in-process Ristretto cache backed by Redis to reduce redundant Postgres queries
- Pagination support across content and user listing endpoints
- Database seeding and JSON-based content import/export tooling for populating lesson data

### Observability & Reliability
- Prometheus metrics for cache behavior, mailer outcomes, and Postgres connection pool usage
- Distributed tracing across HTTP and usecase layers
- Structured logging with Zap, request-scoped fields, request IDs, and access logs
- Per-IP rate limiting, request body size limits, and security headers via middleware

### API & Engineering Practices
- RESTful API built with Gin and documented via generated Swagger/OpenAPI specs
- Clean Architecture: handlers → usecases → repositories, with interfaces decoupling business logic from Postgres/sqlx
- Versioned SQL migrations with a dedicated CLI (`cmd/migration`)
- Unit tests, end-to-end tests, and Testcontainers-based integration tests for Postgres and Redis flows
- CI pipeline and local check scripts covering linting, tests, integration tests, Swagger drift checks, and binary builds

---

## Tech Stack

---

## Environment

---

## Credits

This project is created for independent learning and demonstration purposes.
Special thanks to the original author for their open-source contribution.

---

## Notes

Image resources are for learning and purposes only. Please do not use them for commercial purposes.

If there is any infringement, please contact me for removal. Thank you.

---

## License

This repository is intended for learning and demonstration.

If you plan to open-source it, please choose a license and confirm third-party asset usage rights.

---

## Project Structure

```
```

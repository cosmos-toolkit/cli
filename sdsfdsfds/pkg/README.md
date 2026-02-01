# Reusable packages (pkg)

Packages in `pkg/` extend the application without breaking Clean Architecture’s dependency rule: **dependencies point inward**. `cmd` and adapters (e.g. `internal/repository/postgres`) may import `pkg/`; `entity` and `usecase` must not depend on `pkg/`.

## Typical usage

- **pkg/env** – environment variables / .env (used in `cmd` or config).
- **pkg/database** – SQL client; `usecase.Repository` implementation in `internal/repository/postgres` that uses `pkg/database`.
- **pkg/cache** – Redis or similar (injected into use case via interface if needed).
- **pkg/auth** – JWT, API Key (middleware in HTTP delivery).
- **pkg/logger** – structured logger (injected in handlers or use case via interface).

## Example: Postgres repository as a plugin

1. Create `pkg/database/postgres` with the client.
2. Create `internal/repository/postgres/repository.go` implementing `usecase.Repository` and using `pkg/database/postgres`.
3. In `cmd/api/main.go`, instantiate `repository/postgres.New(cfg)` and pass it to `usecase.New(repo)`.

This way the use case layer still depends only on the `usecase.Repository` interface; the concrete implementation lives in the adapters.

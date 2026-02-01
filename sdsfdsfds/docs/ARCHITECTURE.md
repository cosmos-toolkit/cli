# Clean Architecture (Uncle Bob)

This template organizes code according to **Clean Architecture**: concentric layers with the **dependency rule** — code may only depend on inner layers.

## Layers (innermost to outermost)

```
                    +------------------+
                    | Frameworks       |  HTTP, DB driver, CLI (cmd)
                    | & Drivers        |
                    +--------+---------+
                             |
                    +--------v---------+
                    | Interface       |  Controllers (HTTP), Gateways (Repository impl)
                    | Adapters        |
                    +--------+---------+
                             |
                    +--------v---------+
                    | Use Cases       |  Application rules + Repository interface
                    +--------+---------+
                             |
                    +--------v---------+
                    | Entities        |  Enterprise business rules
                    +------------------+
```

### Entities (`internal/entity`)

- Domain entities and errors.
- Enterprise-level business rules; **no dependencies** on frameworks, DB, or UI.
- Files: `entity.go`, `errors.go`.

### Use Cases (`internal/usecase`)

- Application rules (transaction orchestration).
- Depends only on **entity** and the **interface** `Repository` (defined here).
- Knows nothing about HTTP, SQL, or external libraries.
- Files: `usecase.go`, `repository.go` (interface).

### Interface Adapters

- **Controllers**: `internal/delivery/http` — convert HTTP into use case calls and format the response.
- **Gateways**: `internal/repository/memory` (or `postgres`) — implement `usecase.Repository` and access persistence.

### Frameworks & Drivers

- **cmd/api/main.go** — instantiates repository and use case, sets up the HTTP server.
- HTTP, DB, etc. libraries are used only in **cmd**, **delivery**, and **repository**.

## Dependency rule

- **entity** does not import anything from `internal` or `pkg`.
- **usecase** imports only `entity` and defines the `Repository` interface; it does not import `repository/memory` or `delivery`.
- **delivery** and **repository** import **usecase** (and **entity** only if they need the types).
- **cmd** imports all adapters and performs the wiring.

## Naming (Go)

- **Directories**: single name, lowercase (`entity`, `usecase`, `delivery`, `repository`).
- **Files**: `snake_case.go` for multiple words (`user_repository.go`).
- **Interfaces**: defined on the **consuming** side (e.g. `Repository` in `usecase`).

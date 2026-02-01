# API Clean Architecture (Cosmos Template)

Go API template using **Clean Architecture** (Uncle Bob), following Go community conventions and the dependency rule (dependencies point inward).

## Project structure

```
api-clean-arch/
├── cmd/
│   └── api/                 # Entrypoint (Frameworks & Drivers)
│       └── main.go
├── internal/
│   ├── entity/              # Entities – enterprise business rules
│   │   ├── entity.go
│   │   └── errors.go
│   ├── usecase/             # Use Cases – application rules + Repository interface
│   │   ├── repository.go    # Persistence gateway interface
│   │   └── usecase.go
│   ├── repository/          # Gateway implementations (Interface Adapters)
│   │   └── memory/
│   └── delivery/            # Controllers (Interface Adapters)
│       └── http/
│           ├── handler/
│           └── router/
├── pkg/                     # Reusable packages (env, database, logger, etc.)
├── configs/
├── build/
├── docs/
│   └── ARCHITECTURE.md
├── go.mod
├── Makefile
└── README.md
```

## Conventions (Go)

- **Directories**: lowercase, single word when possible (`entity`, `usecase`, `delivery`, `repository`).
- **Files**: `snake_case.go` for compound names (e.g. `user_repository.go`).
- **Interfaces**: defined in the package that **consumes** them (e.g. `Repository` in `usecase`).
- **Dependency rule**: entity ← usecase ← (repository, delivery); `cmd` does the wiring.

## How to use

1. Copy the template and set the `module` in `go.mod` (e.g. `github.com/your-org/your-api`).
2. Replace `internal/entity` with your entities and errors.
3. Extend `internal/usecase` with your use cases and gateway interfaces.
4. Implement repositories in `internal/repository` (e.g. `postgres`) or use packages in `pkg/`.
5. Run:
   ```bash
   make run
   # or: go run ./cmd/api
   ```

## Endpoints

- `GET /health`
- `GET /api/v1/entities/{id}`
- `POST /api/v1/entities` (body: `{"id": "123"}`)

## Build and tests

```bash
make build   # bin/api
make run     # start the API
make test    # run tests
make lint    # golangci-lint (requires installation)
```

## Packages (pkg)

Packages in `pkg/` are used by `cmd` and by adapters (repository, delivery); they are not imported by `entity` or `usecase`. See `pkg/README.md` for examples (env, database, logger).

## License

As per the Cosmos Toolkit project.

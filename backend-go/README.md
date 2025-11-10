# FMR Tracking System - FIPS-Compliant Go Backend

This is a rewritten backend for the FMR Tracking System using native Go 1.24+ with FIPS 140-3 compliance.

## Key Features

- **FIPS 140-3 Compliant**: Uses Go 1.24's built-in FIPS mode (`GODEBUG=fips140=on`)
- **PostgreSQL**: FIPS-compatible database with pgx/v5 driver
- **REST API**: Clean, maintainable API matching the original Encore endpoints
- **No Framework Lock-in**: Uses standard Go with Chi router

## Architecture

```
backend-go/
├── cmd/server/          # Main application entry point
├── internal/
│   ├── handlers/        # HTTP request handlers
│   ├── repository/      # Database access layer
│   ├── models/          # Data models and types
│   └── middleware/      # HTTP middleware (CORS, logging)
├── migrations/          # SQL database migrations
├── Dockerfile           # FIPS-enabled Docker build
├── go.mod
└── README.md
```

## Prerequisites

- Go 1.24 or later (for FIPS 140-3 support)
- PostgreSQL 13+
- Docker (optional, for containerized deployment)

## Environment Variables

```bash
DATABASE_URL=postgres://user:password@localhost:5432/fmr_db?sslmode=prefer
PORT=4000
GODEBUG=fips140=on  # Enables FIPS mode
```

## Running Locally

### 1. Set up PostgreSQL

```bash
# Using Docker
docker run --name fmr-postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=fmr_db \
  -p 5432:5432 \
  -d postgres:15-alpine

# Run migrations
psql postgres://postgres:postgres@localhost:5432/fmr_db < migrations/001_create_tables.up.sql
# ... (run all migrations in order)
```

### 2. Run the server

```bash
# Install dependencies
go mod download

# Set environment variables
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/fmr_db?sslmode=prefer"
export GODEBUG=fips140=on

# Run the server
go run cmd/server/main.go
```

### 3. Verify FIPS Mode

Check the server logs on startup for FIPS initialization messages:
```
FIPS Mode: Check runtime logs for crypto/internal/fips140 initialization
```

## Building with Docker

```bash
# Build the image
docker build -t fmr-backend-fips .

# Run the container
docker run -p 4000:4000 \
  -e DATABASE_URL="postgres://postgres:postgres@host.docker.internal:5432/fmr_db" \
  fmr-backend-fips
```

## API Endpoints

All endpoints match the original Encore API structure for frontend compatibility:

### FMR Requests
- `POST /fmr` - Create new FMR
- `GET /fmr` - List FMRs (with filters)
- `GET /fmr/{id}` - Get single FMR
- `PUT /fmr/{id}` - Update FMR

### Contacts
- `POST /contacts` - Create contact
- `GET /contacts` - List contacts
- `PUT /contacts/{id}` - Update contact
- `DELETE /contacts/{id}` - Delete contact

### Notes
- `POST /notes` - Create note
- `GET /notes/fmr/{fmrId}` - List notes for FMR

### Presets
- `POST /presets` - Create filter preset
- `GET /presets` - List filter presets
- `DELETE /presets/{id}` - Delete filter preset
- `POST /column-presets` - Create column preset
- `GET /column-presets` - List column presets
- `DELETE /column-presets/{id}` - Delete column preset

### Reports
- `GET /reports/export` - Export FMR data to CSV

## FIPS Compliance Details

### How FIPS Mode Works in Go 1.24+

Go 1.24 introduced native FIPS 140-3 compliance through the `crypto/internal/fips140` module. When `GODEBUG=fips140=on` is set:

1. **Cryptographic Operations**: All crypto operations use FIPS-validated implementations
2. **TLS Connections**: Database and HTTP connections use FIPS-approved cipher suites
3. **Random Number Generation**: Uses FIPS-approved DRBG (Deterministic Random Bit Generator)
4. **Hashing**: Only FIPS-approved hash functions (SHA-2, SHA-3) are available

### Verification

To verify FIPS mode is active:

1. Check server startup logs for FIPS initialization
2. Verify `GODEBUG=fips140=on` environment variable
3. Test TLS connections use only FIPS-approved cipher suites

## Differences from Encore Backend

| Aspect | Encore Backend | New Go Backend |
|--------|----------------|----------------|
| Framework | Encore.dev (transpiled to Go 1.21) | Native Go 1.24+ |
| FIPS Support | ❌ No (Go 1.21) | ✅ Yes (Go 1.24+) |
| Router | Encore framework | Chi router |
| Database Driver | Encore SQLDatabase | pgx/v5 |
| Build Process | `encore build` | Standard `go build` |
| Deployment | Encore cloud or Go binary | Docker, Go binary, any host |

## Migration from Encore

The API endpoints remain unchanged, so the frontend requires no modifications. Simply:

1. Stop the Encore backend
2. Start this Go backend on the same port (4000)
3. Frontend will continue to work without changes

## Development

### Adding a New Endpoint

1. Define models in `internal/models/`
2. Add repository methods in `internal/repository/`
3. Create handler in `internal/handlers/`
4. Register route in `cmd/server/main.go`

### Database Migrations

Place new migrations in `migrations/` directory following the naming convention:
```
NNN_description.up.sql
```

Apply manually or use a migration tool like golang-migrate.

## Troubleshooting

### Database Connection Issues

Ensure PostgreSQL is running and accessible:
```bash
psql $DATABASE_URL -c "SELECT 1"
```

### FIPS Mode Not Active

1. Verify Go version is 1.24+: `go version`
2. Check `GODEBUG` environment variable: `echo $GODEBUG`
3. Review server startup logs for FIPS messages

### Port Already in Use

Change the PORT environment variable:
```bash
export PORT=8080
go run cmd/server/main.go
```

## License

Same as the main FMR Tracking System project.

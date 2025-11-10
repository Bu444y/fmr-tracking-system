# FMR Tracking System

A comprehensive Field Modification Request (FMR) tracking system with **FIPS 140-3 compliant backend**.

## 🚀 Quick Start (Automatic Setup)

**Just run this one command and everything will be configured automatically:**

```bash
bun run setup
```

or

```bash
./setup.sh
```

That's it! The script will:
- ✅ Build all containers (frontend, backend, database)
- ✅ Start all services with correct configuration
- ✅ Wait for everything to be ready
- ✅ Open the app in your browser automatically

**Access the application at: http://localhost:5173**

## 📦 What's Included

- **Frontend**: React 19 + Vite + Tailwind CSS (port 5173)
- **Backend**: Go 1.24+ with FIPS 140-3 compliance (port 4000)
- **Database**: PostgreSQL 15 (port 5432)
- **Automatic configuration** - no manual setup needed!

## 🎯 Common Commands

```bash
# Start everything (automatic setup)
bun run start

# Stop all services
bun run stop

# View logs
bun run logs              # All services
bun run logs:backend      # Backend only
bun run logs:frontend     # Frontend only
bun run logs:db           # Database only

# Restart services
bun run restart

# Check service status
bun run status

# Clean everything (removes volumes)
bun run clean
```

## 🔧 Manual Development (Optional)

If you prefer to run services individually:

```bash
# Terminal 1: Start database
docker-compose up postgres

# Terminal 2: Start backend
bun run dev:backend

# Terminal 3: Start frontend
bun run dev:frontend
```

## 🏗️ Architecture

### Frontend (`/frontend`)
- **Framework**: React 19 with TypeScript
- **Build Tool**: Vite
- **UI Library**: Radix UI + Tailwind CSS
- **State Management**: TanStack Query
- **Port**: 5173

### Backend (`/backend-go`)
- **Language**: Go 1.24+
- **Router**: Chi
- **Database Driver**: pgx/v5 (FIPS-compatible)
- **FIPS Mode**: Enabled via `GODEBUG=fips140=on`
- **Port**: 4000

### Database
- **System**: PostgreSQL 15
- **Port**: 5432
- **Credentials**: postgres/postgres (change in production)

## 🔒 FIPS 140-3 Compliance

The backend uses Go 1.24's native FIPS 140-3 compliance:

- ✅ All cryptographic operations use FIPS-validated implementations
- ✅ TLS connections use FIPS-approved cipher suites
- ✅ Database connections secured with FIPS-compliant crypto
- ✅ Runtime flag: `GODEBUG=fips140=on`

### Verifying FIPS Mode

Check the backend logs after startup:
```bash
bun run logs:backend
```

Look for FIPS initialization messages confirming FIPS mode is active.

## 📂 Project Structure

```
fmr-tracking-system/
├── frontend/              # React frontend application
│   ├── src/
│   ├── Dockerfile
│   └── package.json
├── backend-go/           # FIPS-compliant Go backend
│   ├── cmd/server/       # Main application
│   ├── internal/         # Business logic
│   ├── migrations/       # Database migrations
│   ├── Dockerfile
│   └── README.md
├── docker-compose.yml    # Full-stack orchestration
├── setup.sh              # Automated setup script
└── package.json          # Root package with scripts
```

## 🔄 Migration from Encore Backend

The old Encore.dev backend has been replaced with a native Go implementation for FIPS compliance. The API endpoints remain 100% compatible, so:

- ✅ No frontend changes required
- ✅ Same API structure
- ✅ Same database schema
- ✅ Seamless migration

## 🐳 Docker Configuration

The `docker-compose.yml` automatically configures:

1. **PostgreSQL** with migrations applied on startup
2. **Go Backend** with FIPS mode enabled
3. **React Frontend** connected to backend
4. **Network** for inter-service communication
5. **Volume** for persistent database storage

## 🔍 Troubleshooting

### Services won't start
```bash
# Check status
bun run status

# View logs
bun run logs

# Clean and restart
bun run clean
bun run setup
```

### Port already in use
Stop other services using ports 4000, 5173, or 5432:
```bash
# Check what's using the port
lsof -i :4000
lsof -i :5173
lsof -i :5432

# Stop FMR services
bun run stop
```

### Database connection issues
```bash
# Check if PostgreSQL is running
docker-compose ps postgres

# Restart database
docker-compose restart postgres

# View database logs
bun run logs:db
```

### FIPS mode not working
```bash
# Check backend logs for FIPS initialization
bun run logs:backend | grep -i fips

# Verify GODEBUG is set
docker-compose exec backend env | grep GODEBUG
```

## 📊 Features

### FMR Management
- Create, view, update, and archive FMRs
- Track through lifecycle: Draft → Submitted → In-Progress → Resolved
- 60+ fields covering all aspects of failure tracking

### Contact Management
- Store contact information with team assignments
- Organization tracking
- Quick lookup and filtering

### Notes System
- Add notes to FMRs
- Multiple note types: general, technician, customer
- Chronological tracking

### Reporting & Export
- CSV export functionality
- Date range filtering
- Weekly grouping
- Export by unit

### Presets
- Filter presets for quick access
- Column presets for customized views
- Save and share configurations

## 🚢 Production Deployment

For production:

1. Update database credentials in `docker-compose.yml`
2. Set appropriate environment variables
3. Use production build targets
4. Enable HTTPS/TLS
5. Configure backups

See `backend-go/README.md` for detailed deployment instructions.

## 📝 License

[Your License Here]

## 🤝 Contributing

[Your Contributing Guidelines Here]

## 📞 Support

[Your Support Information Here]

---

**Built with FIPS 140-3 compliance for secure government and enterprise deployments** 🔒

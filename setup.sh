#!/bin/bash

# FMR Tracking System - Automated Setup Script
# This script automatically sets up and runs the entire system

set -e

echo "🚀 FMR Tracking System - Automated Setup"
echo "=========================================="
echo ""

# Check for required tools
echo "📋 Checking prerequisites..."

if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker first."
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

echo "✅ Prerequisites met"
echo ""

# Stop any existing containers
echo "🛑 Stopping existing containers..."
docker-compose down 2>/dev/null || true
echo ""

# Build containers
echo "🏗️  Building containers (this may take a few minutes)..."
docker-compose build
echo ""

# Start services
echo "🚀 Starting services..."
docker-compose up -d
echo ""

# Wait for services to be healthy
echo "⏳ Waiting for services to be ready..."
echo "   - Waiting for PostgreSQL..."
until docker-compose exec -T postgres pg_isready -U postgres &>/dev/null; do
    sleep 1
done
echo "   ✅ PostgreSQL is ready"

echo "   - Waiting for backend..."
until curl -f http://localhost:4000/health &>/dev/null; do
    sleep 1
done
echo "   ✅ Backend is ready"

echo "   - Waiting for frontend..."
until curl -f http://localhost:5173 &>/dev/null; do
    sleep 1
done
echo "   ✅ Frontend is ready"
echo ""

# Display status
echo "✅ FMR Tracking System is running!"
echo ""
echo "📍 Access Points:"
echo "   Frontend:  http://localhost:5173"
echo "   Backend:   http://localhost:4000"
echo "   Database:  localhost:5432"
echo ""
echo "🔧 Useful commands:"
echo "   bun run logs         - View all logs"
echo "   bun run logs:backend - View backend logs"
echo "   bun run logs:frontend- View frontend logs"
echo "   bun run stop         - Stop all services"
echo "   bun run restart      - Restart services"
echo "   bun run status       - Check service status"
echo ""
echo "🔒 FIPS Mode: Enabled (GODEBUG=fips140=on)"
echo ""
echo "Opening frontend in browser..."
sleep 2

# Try to open browser (works on most systems)
if command -v xdg-open &> /dev/null; then
    xdg-open http://localhost:5173
elif command -v open &> /dev/null; then
    open http://localhost:5173
else
    echo "Please open http://localhost:5173 in your browser"
fi

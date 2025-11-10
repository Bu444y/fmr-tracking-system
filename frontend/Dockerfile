# Frontend Dockerfile for FMR Tracking System
FROM oven/bun:1-alpine AS builder

WORKDIR /app

# Copy package files
COPY package.json bun.lockb* ./

# Install dependencies
RUN bun install --frozen-lockfile

# Copy source code
COPY . .

# Build argument for API endpoint
ARG VITE_CLIENT_TARGET=http://localhost:4000
ENV VITE_CLIENT_TARGET=$VITE_CLIENT_TARGET

# Development stage (default)
FROM builder AS development
EXPOSE 5173
CMD ["bun", "run", "vite", "--host", "0.0.0.0"]

# Production build stage
FROM builder AS build
RUN bun run vite build

# Production runtime stage
FROM nginx:alpine AS production
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

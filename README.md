# URL Shortener (Go + Fiber)

Minimal URL shortener with two routes: one to register new short URLs and one to resolve them by short code. Uses PostgreSQL, caching, and rate limiting.

## Tech
- Go + Fiber
- PostgreSQL (docker-compose)
- Fiber middleware: cache (30s), limiter

## Project layout
- cmd/main.go
- internal/
    - db
    - env
    - handlers
    - models
    - router
    - scripts

## Environment
Create a .env file (or export these vars):
```env
DB_NAME="url_shortener"
DB_USER="postgres"
DB_PASSWORD="postgres"
DB_HOST="localhost"
DB_PORT=5432
TIMEZONE="UTC"
PORT=8080
```

## Database (Docker)
docker-compose.yml:
```yaml
services:
    postgres:
        image: postgres:15-alpine
        ports:
            - "5432:5432"
        environment:
            - POSTGRES_USER=postgres
            - POSTGRES_PASSWORD=postgres
            - POSTGRES_DB=url_shortener
        volumes:
            - postgres_data:/var/lib/postgresql/data
        restart: unless-stopped

volumes:
    postgres_data:
```

Start DB:
```sh
docker compose up -d
```

## Run
```sh
go run ./cmd/main.go
# or
go build -o url-shortener ./cmd && ./url-shortener
```

## Routes

- GET /:shortCode
    - Resolves a short code and redirects to the original URL.
    - Example:
        ```sh
        curl -I http://localhost:8080/abc123
        ```
        Expected: 301/302 redirect if found, 404 if not.

- POST /api/v1/url/register
    - Registers a new short URL.
    - Cached for 30s and protected by a rate limiter.
    - Example request payload:
        ```sh
        curl -X POST http://localhost:8080/api/v1/url/register \
            -H "Content-Type: application/json" \
            -d '{"target_path":"https://example.com"}'
        ```
    - Example response:
        ```json
        {
            "short_path": "http://localhost:8080/7288",
            "target_path": "https://google.com"
        }
        ```

Notes:
- The API group uses a 30s cache and a rate limiter; bursty requests may receive 429.
- Ensure the DB is up and env vars are set before running.

## Development
- Router entry: internal/router/router.go
- Handlers: internal/handlers
- Models and DB access: internal/models, internal/db
- Env loading: internal/env
- Entry point: cmd/main.go
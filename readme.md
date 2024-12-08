# Pixel Board Game

A real-time pixel painting game where users can collaboratively paint a board, with updates reflected instantly for all participants.

## Getting Started

### Prerequisites

- Docker and Docker Compose installed
- Go installed
- Node.js and npm installed

### Running Locally

1. **Start Traefik**

   Traefik acts as a reverse proxy, allowing multiple services (backend, frontend, Centrifugo) to be exposed through a single entry point. This is necessary because Ngrok's free plan only allows one port. Using Traefik, you can configure exposure of all necessary services during development.

   ```bash
   make docker
   ```

2. **Configure Environment Variables**

   Create a `.env` file in the project's root directory based on the example below. Ensure to keep the `SECRET_KEY` and `CENTRIFUGO_SECRET_KEY` unchanged as they are aligned with the `centrifugo.json` configuration.

   ```env
   MYSQL_DSN="username:password@tcp(127.0.0.1:3306)/pixel?charset=utf8mb4&parseTime=True"
   TELEGRAM_TOKEN="your_telegram_token"
   WEBAPP_URL="https://your-ngrok-url.ngrok-free.app"
   SECRET_KEY="e7c48b48-d80f-4249-bf89-e90ab9641cd0"
   CENTRIFUGO_ADDR_API="http://localhost/ws/api"
   CENTRIFUGO_SECRET_KEY="83892ae7-2bb8-47fb-b93f-52c23e20f8af"
   TEST_TOKEN_REPLACE="your_test_token"
   NGROK_URL=your-ngrok-url.ngrok-free.app
   ```

3. **Run the Application**

   ```bash
   make serve
   ```

4. **Run the Client**

   Navigate to the `ui` directory and start the client:

   ```bash
   cd ui
   npm run dev
   ```

### Generating TypeScript Types

Serializers bridge the client and backend by defining data structures. To generate TypeScript types from the serializers:

1. Add your serializer to the `cmd/ts.go` file.
2. Run the following command:

   ```bash
   make ts
   ```

This will generate the TypeScript definitions in `./ui/src/types/serializer.ts`.

## Available Make Commands

- `make ts` - Generate TypeScript types from serializers
- `make docker` - Start development Docker environment
- `make serve` - Run the application
- `make stop-docker` - Stop Docker services
- `make ngrok` - Start Ngrok with the specified URL
- `make test` - Run tests

## Traefik Configuration

The project uses Traefik as a reverse proxy. Ensure the following configuration files are present:

- **traefik.yml**

  ```yaml
  entryPoints:
    web:
      address: ":80"

  providers:
    file:
      filename: "/dynamic.yml"
      watch: true

  api:
    dashboard: true
  ```

- **dynamic.yml**

  ```yaml
  http:
    routers:
      pixel-frontend:
        rule: "PathPrefix(`/pixel`)"
        entryPoints:
          - web
        service: pixel-frontend-service
      pixel-backend:
        rule: "PathPrefix(`/pixel/api`)"
        entryPoints:
          - web
        middlewares:
          - strip-pixel-api-prefix
        service: pixel-backend-service
      pixel-events:
        rule: "PathPrefix(`/pixel/events`)"
        entryPoints:
          - web
        middlewares:
          - strip-pixel-events-prefix
        service: pixel-events-service

    services:
      pixel-frontend-service:
        loadBalancer:
          servers:
            - url: "http://localhost:3000"
      pixel-backend-service:
        loadBalancer:
          servers:
            - url: "http://localhost:8001"
      pixel-events-service:
        loadBalancer:
          servers:
            - url: "http://localhost:8000"

    middlewares:
      strip-pixel-api-prefix:
        stripPrefix:
          prefixes:
            - "/pixel/api"
      strip-pixel-events-prefix:
        stripPrefix:
          prefixes:
            - "/pixel/events"
  ```

- **docker-compose.yml**

  ```yaml
  services:
    traefik:
      image: traefik:v3.2
      container_name: local-traefik
      restart: always
      volumes:
        - "./traefik.yml:/traefik.yml:ro"
        - "./dynamic.yml:/dynamic.yml:ro"
      network_mode: host
  ```

Ensure Traefik is running correctly to manage the routing of frontend, backend, and event services.

## License

This project is licensed under the MIT License.
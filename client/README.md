# Client

Static web UI for the Lamrim HTTP server.

## Run

```bash
go run ./cmd/lamrim -serve
```

Open http://localhost:8080

## API used

- `GET /api/workflows`
- `POST /api/workflows/{name}/run`

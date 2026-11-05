# API contract

This is the agreement between the frontend and backend. If both sides follow it,
they work together. **Update this file first** whenever you add or change an endpoint.

Base URL (development): `http://localhost:8080`
All bodies are JSON. Errors always look like: `{ "error": "message" }`

## Todo object

```json
{ "id": 1, "title": "Buy milk", "done": false }
```

## Endpoints

| Method | Path              | Body                                 | Success              | Errors              | Status        |
|--------|-------------------|--------------------------------------|----------------------|---------------------|---------------|
| GET    | `/api/health`     | –                                    | `200 {"status":"ok"}`| –                   | ✅ done        |
| GET    | `/api/todos`      | –                                    | `200 [Todo, ...]`    | –                   | ✅ done        |
| GET    | `/api/todos/{id}` | –                                    | `200 Todo`           | `400` bad id, `404` | ⬜ Step 3      |
| POST   | `/api/todos`      | `{"title":"Buy milk"}`               | `201 Todo`           | `400` empty title   | ⬜ Step 3      |
| PATCH  | `/api/todos/{id}` | `{"title"?: "...", "done"?: true}`   | `200 Todo`           | `400`, `404`        | ⬜ Step 3      |
| DELETE | `/api/todos/{id}` | –                                    | `204` (no body)      | `400`, `404`        | ⬜ Step 3      |

## Testing with curl

```bash
curl -i http://localhost:8080/api/todos
curl -i -X POST http://localhost:8080/api/todos -H "Content-Type: application/json" -d '{"title":"Buy milk"}'
curl -i -X PATCH http://localhost:8080/api/todos/1 -H "Content-Type: application/json" -d '{"done":true}'
curl -i -X DELETE http://localhost:8080/api/todos/1
```

`-i` prints the status code and headers, so you can check you return the right code.

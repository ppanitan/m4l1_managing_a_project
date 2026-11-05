# Step-by-step guide

Each step has a **Goal**, the **Concepts** it teaches, **Tasks**, **Hints** (hidden, so try
first), a **Checkpoint** that tells you when you're done, and **Study** links.

Work in order. Commit after every step: `git add . && git commit -m "Step 3: create todo"`.

---

## Step 0: Run it

**Goal:** both servers running and talking to each other.

**Tasks**
1. Install [Go](https://go.dev/dl/) and [Node.js LTS](https://nodejs.org/).
2. Follow *Quick start* in the README.
3. Stop the Go server (Ctrl+C) and reload the browser. What happens? Start it again.

**Checkpoint:** `http://localhost:3000` lists 2 todos, and `http://localhost:8080/api/todos` shows raw JSON.

---

## Step 1: Trace one request

**Goal:** understand how the existing code fits together before changing anything.

**Concepts:** routing, handlers, JSON encoding, Server Components.

**Tasks.** Read these files in this order and follow the arrows:

```
frontend/app/page.tsx        HomePage() calls getTodos()
        │
frontend/lib/api.ts          getTodos() → fetch("http://localhost:8080/api/todos")
        │  HTTP GET
backend/middleware.go        withCORS adds headers, passes the request on
        │
backend/main.go              mux matches "GET /api/todos" → store.listTodos
        │
backend/handlers.go          listTodos → writeJSON(200, s.todos)
        │
backend/todo.go              the Todo struct → JSON via `json:"..."` tags
```

Then answer for yourself:
- Why does `page.tsx` have no `"use client"`, while `AddTodoForm.tsx` does?
- What would break if you renamed `json:"title"` to `json:"name"` in Go?
- Why does `Store` have a `sync.Mutex`?

<details><summary>Answers</summary>

- `page.tsx` only fetches and displays, so it can run on the server. The form needs `useState` and `onSubmit`, which only exist in the browser.
- The frontend expects `title` (see the `Todo` type in `lib/api.ts`). The list would show empty items. This is why the API contract matters.
- Go handles each request in its own goroutine, so two requests could modify the slice at the same time. The mutex makes them take turns.
</details>

**Study:** [Go by Example](https://gobyexample.com/) (structs, methods, slices) · [Next.js: Server and Client Components](https://nextjs.org/docs/app/getting-started/server-and-client-components)

---

## Step 2: Explore the API

**Goal:** get comfortable seeing raw HTTP.

**Tasks**
1. Run the curl commands in `docs/API.md`. Note the status codes (`-i`).
2. In the browser, open DevTools (F12) → **Network** tab and reload `localhost:3000`.
   Notice that you **don't** see a request to `:8080`. Why?
3. Optional: install [Bruno](https://www.usebruno.com/) or Postman to send requests with a UI.

<details><summary>Why no :8080 request in DevTools?</summary>

`page.tsx` is a Server Component, so the Next.js **server** called Go, not your browser.
The browser received finished HTML. In Step 4 your form calls Go **from the browser**,
and then you'll see the request (including an `OPTIONS` preflight, which is CORS at work).
</details>

**Study:** [MDN: HTTP overview](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview) · [HTTP status codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

---

## Step 3: Backend CRUD (Go)

**Goal:** implement the 4 stub handlers in `backend/handlers.go`.

**Concepts:** path parameters, decoding JSON, validation, status codes, error handling.

**Tasks** (one at a time, testing each with curl before moving on):
1. `getTodo`: uncomment its route in `main.go`, implement, and test.
2. `createTodo`: decode body, validate, assign `s.nextID`, append, return `201`.
3. `deleteTodo`: find by id, remove from slice, return `204`.
4. `updateTodo`: partial update. Only change the fields that were sent.
5. Refactor: you'll write "find todo by id" several times. Extract a helper method.
6. Refactor: write a `writeError(w, status, msg)` helper so all errors look the same.

<details><summary>Hint: reading the id</summary>

```go
id, err := strconv.Atoi(r.PathValue("id"))
if err != nil {
    writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
    return
}
```
</details>

<details><summary>Hint: decoding the body</summary>

```go
var input struct {
    Title string `json:"title"`
}
if err := json.NewDecoder(r.Body).Decode(&input); err != nil { /* 400 */ }
input.Title = strings.TrimSpace(input.Title)
```
</details>

<details><summary>Hint: removing from a slice</summary>

`s.todos = append(s.todos[:i], s.todos[i+1:]...)`, or `slices.Delete(s.todos, i, i+1)` (Go 1.21+).
</details>

<details><summary>Hint: partial update with pointers</summary>

```go
var input struct {
    Title *string `json:"title"`
    Done  *bool   `json:"done"`
}
// after decoding:
if input.Done != nil { todo.Done = *input.Done }
```
Careful: if you loop with `for _, t := range s.todos`, `t` is a **copy**.
Use `for i := range s.todos { s.todos[i].Done = ... }` instead.
</details>

**Checkpoint:** every curl command in `docs/API.md` returns the documented status. Also check that
`POST` with `{"title":""}` returns `400`, and `DELETE /api/todos/999` returns `404`.
Update the *Status* column in `API.md`.

**Study:** [Go by Example: JSON](https://gobyexample.com/json), [HTTP servers](https://gobyexample.com/http-servers) · [net/http routing patterns (Go 1.22)](https://go.dev/blog/routing-enhancements)

---

## Step 4: Frontend: add a todo

**Goal:** typing in the form creates a todo via the Go API and the list updates.

**Concepts:** Client Components, `useState`, `fetch` with POST, `router.refresh()`, CORS.

**Tasks**
1. In `lib/api.ts`, write `createTodo(title: string): Promise<Todo>`.
2. In `components/AddTodoForm.tsx`, replace the `alert` with a real call (follow the TODO).
3. Don't allow submitting an empty title (disable the button).
4. Show an error message if the request fails. Test it by stopping the Go server.
5. Open DevTools → Network and find your `POST` and the `OPTIONS` preflight before it.

<details><summary>Hint: createTodo</summary>

```ts
export async function createTodo(title: string): Promise<Todo> {
  const res = await fetch(`${API_URL}/api/todos`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ title }),
  });
  if (!res.ok) throw new Error(`create failed: ${res.status}`);
  return res.json();
}
```
</details>

**Experiment:** in `backend/middleware.go`, temporarily change the allowed origin to
`http://localhost:9999` and try adding a todo. Read the red CORS error in the browser
console. Seeing that error once makes it much easier to debug later. Then change it back.

**Checkpoint:** a new todo appears without a manual page reload, and it's still there after a reload.
(It disappears when Go restarts. Step 6 fixes that.)

**Study:** [React: State](https://react.dev/learn/state-a-components-memory) · [React: Responding to events](https://react.dev/learn/responding-to-events) · [MDN: CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)

---

## Step 5: Frontend: toggle and delete

**Goal:** each todo has a checkbox (done or not) and a delete button.

**Concepts:** component design, props, splitting Server and Client Components.

**Tasks**
1. Add `toggleTodo(id, done)` and `deleteTodo(id)` to `lib/api.ts`.
2. Create `components/TodoItem.tsx` (a Client Component) that receives a `todo` prop and renders
   the checkbox, title and delete button.
3. In `page.tsx`, render `<TodoItem todo={t} />` for each item. Keep `page.tsx` a Server Component.
4. Ask before deleting with `confirm("Delete?")`.

<details><summary>Hint: why a separate component?</summary>

You want `page.tsx` to stay on the server (it fetches data). Only the small interactive part
needs to be a Client Component. This "push `use client` down the tree" pattern keeps the amount
of JavaScript sent to the browser small.
</details>

**Checkpoint:** you can add, complete and delete todos in the browser, and the UI matches `curl /api/todos`.

**Study:** [React: Passing props](https://react.dev/learn/passing-props-to-a-component) · [React: Rendering lists](https://react.dev/learn/rendering-lists)

🎉 **You now have a full working CRUD app.** Steps 6 to 9 make it "real".

---

## Step 6: A real database (PostgreSQL)

**Goal:** data survives a server restart.

**Concepts:** SQL, connection pools, interfaces, migrations, SQL injection, environment config.

**Tasks**
1. Run Postgres with Docker:
   `docker run --name todo-db -e POSTGRES_PASSWORD=dev -p 5432:5432 -d postgres:17`
2. Create the table (save this as `backend/migrations/001_create_todos.sql`):
   ```sql
   CREATE TABLE todos (
     id    SERIAL PRIMARY KEY,
     title TEXT    NOT NULL,
     done  BOOLEAN NOT NULL DEFAULT false
   );
   ```
3. `go get github.com/jackc/pgx/v5`, then connect using an env var
   `DATABASE_URL=postgres://postgres:dev@localhost:5432/postgres`.
4. **Design step:** define an interface so handlers don't care where the data lives:
   ```go
   type TodoRepo interface {
       List(ctx context.Context) ([]Todo, error)
       Get(ctx context.Context, id int) (Todo, error)
       Create(ctx context.Context, title string) (Todo, error)
       Update(ctx context.Context, id int, title *string, done *bool) (Todo, error)
       Delete(ctx context.Context, id int) error
   }
   ```
   Make your in-memory store implement it, then write a `PostgresRepo` that also implements it.
   Switch between them in `main.go`.
5. Always use placeholders: `WHERE id = $1`. **Never** build SQL with `fmt.Sprintf`.

**Checkpoint:** add todos, restart Go, and they're still there. The frontend needed **zero** changes.
(That's the value of an API contract.)

**Study:** [SQLBolt](https://sqlbolt.com/) (interactive SQL) · [pgx docs](https://github.com/jackc/pgx) · [Go by Example: Interfaces](https://gobyexample.com/interfaces) · [OWASP: SQL Injection](https://owasp.org/www-community/attacks/SQL_Injection) · later: [golang-migrate](https://github.com/golang-migrate/migrate) or [sqlc](https://sqlc.dev/)

---

## Step 7: Backend polish

**Tasks**
1. `withLogging` middleware (see TODO in `middleware.go`) prints method, path and duration.
2. Move the allowed CORS origin into an env var (`FRONTEND_ORIGIN`).
3. Write tests with `net/http/httptest` for at least create and delete (`handlers_test.go`, run `go test ./...`).
4. Optional: reorganize into folders (`internal/handlers`, `internal/store`) as the project grows.

**Study:** [Go by Example: Testing](https://gobyexample.com/testing-and-benchmarking) · [httptest](https://pkg.go.dev/net/http/httptest) · [log/slog](https://go.dev/blog/slog)

---

## Step 8: Frontend polish

**Tasks**
1. Add `app/loading.tsx` (shown while the page fetches) and `app/error.tsx` (shown on failure).
2. Add a second page `app/todos/[id]/page.tsx` that shows one todo (uses your `GET /api/todos/{id}`), and link to it.
3. Style it with [Tailwind CSS](https://tailwindcss.com/docs/installation/framework-guides/nextjs).
4. Optional: try **Server Actions** as an alternative to calling the API from the browser, and compare the two approaches.

**Study:** [Next.js: Loading UI](https://nextjs.org/docs/app/api-reference/file-conventions/loading) · [Error handling](https://nextjs.org/docs/app/getting-started/error-handling) · [Dynamic routes](https://nextjs.org/docs/app/api-reference/file-conventions/dynamic-routes) · [Server Actions](https://nextjs.org/docs/app/getting-started/updating-data)

---

## Step 9 (stretch): Users, Docker, deploy

Pick any:
- **Auth:** `users` table, `POST /api/register` and `/api/login`, bcrypt passwords, a JWT in an
  **HttpOnly cookie**, a middleware that rejects requests without a valid token, and a `user_id` column on todos so each user sees only their own.
  (CORS then needs `Access-Control-Allow-Credentials: true`, and fetch needs `credentials: "include"`.)
- **Docker Compose:** one `docker compose up` that starts db + backend + frontend.
- **Deploy:** frontend on Vercel, backend and database on Fly.io, Railway or Render.

**Study:** [OWASP Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html) · [jwt.io intro](https://jwt.io/introduction) · [Docker Get Started](https://docs.docker.com/get-started/)

---

## When you get stuck

1. **Read the error message fully.** Go errors and the browser console usually say exactly what's wrong.
2. **Find which side is broken.** Does `curl` work? Then the backend is fine and the bug is in the frontend.
3. **DevTools → Network:** check the request URL, method, body, status and response.
4. Add `log.Printf(...)` in Go and `console.log(...)` in React to see values.
5. Commit working code often, so you can always go back (`git diff`, `git stash`).

# Todo Starter: Next.js + Go

A deliberately **unfinished** full-stack project for learning.
The pieces are wired together end to end, and **you** write the rest.

```
 Browser ──► Next.js (frontend, :3000) ──HTTP/JSON──► Go API (backend, :8080) ──► (later) PostgreSQL
```

## What already works

- Go server with `GET /api/health` and `GET /api/todos` (in-memory data)
- CORS middleware so the browser may call the API
- Next.js page that fetches the todos from Go and lists them
- A form component that does nothing yet

## Quick start

You need **Go 1.22+** and **Node.js 20+**.

```bash
# Terminal 1: backend
cd starter-app/backend
go run .                      # → http://localhost:8080/api/todos

# Terminal 2: frontend
cd starter-app/frontend
cp .env.example .env.local
npm install
npm run dev                   # → http://localhost:3000
```

You should see two todos in the browser. That data came from Go.

## Folder map

```
starter-app/
├── docs/
│   ├── GUIDE.md          ← START HERE: step-by-step tasks
│   └── API.md            ← the contract between front and back
├── backend/              (Go)
│   ├── main.go           routes + server start
│   ├── handlers.go       request handlers (4 stubs for you to finish)
│   ├── todo.go           Todo type + in-memory Store
│   └── middleware.go     CORS (+ your logging middleware later)
└── frontend/             (Next.js, App Router)
    ├── app/layout.tsx    shared wrapper for all pages
    ├── app/page.tsx      home page (Server Component)
    ├── components/
    │   └── AddTodoForm.tsx   form (Client Component, unfinished)
    └── lib/api.ts        all fetch calls to the Go API
```

## Progress checklist

Tick these off as you go (edit this file and commit, which is good Git practice too).

- [ ] Step 0: both servers run, todos show in the browser
- [ ] Step 1: traced one request through the code
- [ ] Step 2: explored the API with curl and DevTools
- [ ] Step 3: backend CRUD: get / create / update / delete
- [ ] Step 4: frontend: add a todo
- [ ] Step 5: frontend: toggle and delete
- [ ] Step 6: PostgreSQL instead of memory
- [ ] Step 7: backend polish: logging, tests
- [ ] Step 8: frontend polish: loading/error states, styling
- [ ] Step 9 (stretch): users and login, Docker, deploy

Search the code for `TODO (Step N)` to find where each step starts.

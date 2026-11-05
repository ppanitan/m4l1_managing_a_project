// All calls to the Go backend live here, so pages and components never
// hard-code URLs. This file is the frontend side of the "API contract"
// described in docs/API.md.

export const API_URL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

// Must match the Go struct `Todo` in backend/todo.go
export type Todo = {
  id: number;
  title: string;
  done: boolean;
};

export async function getTodos(): Promise<Todo[]> {
  const res = await fetch(`${API_URL}/api/todos`, { cache: "no-store" });
  if (!res.ok) throw new Error(`GET /api/todos failed: ${res.status}`);
  return res.json();
}

// TODO (Step 4): createTodo(title)  → POST   /api/todos
// TODO (Step 5): toggleTodo(id, done) → PATCH  /api/todos/{id}
// TODO (Step 5): deleteTodo(id)     → DELETE /api/todos/{id}
// Hint: fetch(url, { method: "POST", headers: { "Content-Type": "application/json" },
//                    body: JSON.stringify({ title }) })

// This is a SERVER Component (no "use client" at the top).
// It runs on the Next.js server, calls the Go API, and sends finished HTML.
import { getTodos } from "@/lib/api";
import AddTodoForm from "@/components/AddTodoForm";

export default async function HomePage() {
  let todos;
  try {
    todos = await getTodos();
  } catch {
    return <p className="error">Cannot reach the Go backend. Is it running on :8080?</p>;
  }

  return (
    <>
      <h1>My Todos</h1>

      <AddTodoForm />

      <ul>
        {todos.map((t) => (
          <li key={t.id} className={t.done ? "done" : ""}>
            {t.title}
            {/* TODO (Step 5): add a checkbox to toggle and a delete button.
                They need onClick → that means a Client Component. */}
          </li>
        ))}
      </ul>
    </>
  );
}

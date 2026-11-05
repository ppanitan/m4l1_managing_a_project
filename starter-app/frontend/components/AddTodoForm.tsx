// A CLIENT Component: it runs in the browser because it needs state and events.
"use client";

import { useState } from "react";

export default function AddTodoForm() {
  const [title, setTitle] = useState("");

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    // TODO (Step 4):
    //   1. call createTodo(title) from lib/api.ts
    //   2. clear the input
    //   3. refresh the list: const router = useRouter() (from "next/navigation"),
    //      then router.refresh() re-runs the Server Component in page.tsx
    //   4. bonus: show an error message if the request fails
    alert(`Not implemented yet: would create "${title}"`);
  }

  return (
    <form onSubmit={handleSubmit}>
      <input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="What needs doing?"
      />
      <button type="submit">Add</button>
    </form>
  );
}

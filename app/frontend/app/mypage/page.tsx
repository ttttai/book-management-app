"use client";

import { useState } from "react";
import Header from "../components/header";

export default function MyPage() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <Header title="マイページ" />
      <div className="h-screen flex items-center justify-center bg-blue-100">
        <h1 className="text-2xl font-bold">マイページ</h1>

        <button className="bg-blue-500" onClick={() => setCount(count + 1)}>
          Click!
        </button>

        <p>{count}</p>
      </div>
    </div>
  );
}

"use client";

import { useRouter } from "next/navigation";
import Header from "../components/header";

export default function BookshelfPage() {
  const router = useRouter();

  return (
    <div>
      <Header title="マイ本棚" />
      <div className="h-screen flex items-center justify-center bg-purple-100">
        <h1 className="flex text-2xl font-bold items-center justify-center">
          マイ本棚
        </h1>
        <button
          className="w-60 px-4 py-2 bg-blue-500 text-white text-lg font-semibold rounded-lg shadow-md hover:bg-blue-600 transition mx-4"
          onClick={() => router.push("/bookshelf/book/1")}
        >
          IDページへ
        </button>
      </div>
    </div>
  );
}

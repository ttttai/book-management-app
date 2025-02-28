"use client";

import { useRouter } from "next/navigation";

export default function BookshelfPage() {
  const router = useRouter();

  return (
    <div className="h-screen flex items-center justify-center">
      <h1 className="flex text-2xl font-bold items-center justify-center">
        マイ本棚
      </h1>
      <button
        className="w-60 px-4 py-2 bg-blue-500 text-white text-lg font-semibold rounded-lg shadow-md hover:bg-blue-600 transition mx-4"
        onClick={() => router.push("/bookshelf/book/2")}
      >
        IDページへ
      </button>
    </div>
  );
}

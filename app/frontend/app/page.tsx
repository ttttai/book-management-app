"use client";

import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  return (
    <div className="flex flex-col items-center justify-center h-screen bg-gray-100">
      <h1 className="text-3xl font-bold mb-8">ホーム画面</h1>
      <div className="space-y-4">
        <button
          className="w-60 px-4 py-2 bg-blue-500 text-white text-lg font-semibold rounded-lg shadow-md hover:bg-blue-600 transition mx-4"
          onClick={() => router.push("/mypage")}
        >
          マイページ
        </button>
        <button
          className="w-60 px-4 py-2 bg-green-500 text-white text-lg font-semibold rounded-lg shadow-md hover:bg-green-600 transition mx-4"
          onClick={() => router.push("/search")}
        >
          本を検索
        </button>
        <button
          className="w-60 px-4 py-2 bg-purple-500 text-white text-lg font-semibold rounded-lg shadow-md hover:bg-purple-600 transition mx-4"
          onClick={() => router.push("/bookshelf")}
        >
          マイ本棚
        </button>
      </div>
    </div>
  );
}

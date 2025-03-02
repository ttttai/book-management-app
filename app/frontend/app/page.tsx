"use client";

import { useRouter } from "next/navigation";
import { Search, Library, BookOpen, Star, Bot } from "lucide-react";

export default function Home() {
  const router = useRouter();

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-50 to-blue-100">
      {/* ヒーローセクション */}
      <div className="flex flex-col items-center justify-center text-center py-20">
        <h1 className="text-5xl font-bold text-gray-800">
          MyLibraryへようこそ！
        </h1>
        <p className="text-lg text-gray-600 mt-4">
          あなたの読書体験をもっと楽しく、もっと便利に。
        </p>
      </div>

      {/* 機能紹介セクション */}
      <div className="max-w-5xl mx-auto grid grid-cols-1 md:grid-cols-3 gap-6 mt-10">
        {/* カード：本を検索 */}
        <div
          className="p-6 bg-white shadow-md rounded-lg flex flex-col items-center text-center shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95"
          onClick={() => router.push("/search")}
        >
          <Search className="w-12 h-12 text-blue-500" />
          <h2 className="text-xl font-semibold mt-3">本を検索</h2>
          <p className="text-gray-600 mt-2">タイトルや著者名で検索できます。</p>
        </div>

        {/* カード：本棚管理 */}
        <div
          className="p-6 bg-white shadow-md rounded-lg flex flex-col items-center text-center shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95"
          onClick={() => router.push("/bookshelf")}
        >
          <Library className="w-12 h-12 text-green-500" />
          <h2 className="text-xl font-semibold mt-3">本棚管理</h2>
          <p className="text-gray-600 mt-2">
            購入済や読書中の本を整理できます。
          </p>
        </div>

        {/* カード：読書進捗 */}
        <div
          className="p-6 bg-white shadow-md rounded-lg flex flex-col items-center text-center shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95"
          onClick={() => router.push("/recommendation")}
        >
          <Bot className="w-12 h-12 text-blue-600" />
          <h2 className="text-xl font-semibold mt-3">AIのおすすめ</h2>
          <p className="text-gray-600 mt-2">読書履歴からAIが本を推薦します。</p>
        </div>
      </div>
    </div>
  );
}

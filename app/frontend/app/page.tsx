import { Search, Library, Bot } from "lucide-react";
import Link from "next/link";
import { getServerSession } from "next-auth";
import { authOptions } from "@/app/lib/authOptions";

export default async function Home() {
  const session = await getServerSession(authOptions);
  return (
    <div>
      {session ? (
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
            <Link
              href={"/search"}
              className="p-6 bg-white shadow-md rounded-lg flex flex-col items-center text-center shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95"
            >
              <Search className="w-12 h-12 text-blue-500" />
              <h2 className="text-xl font-semibold mt-3">本を検索</h2>
              <p className="text-gray-600 mt-2">
                タイトルや著者名で検索できます。
              </p>
            </Link>

            {/* カード：本棚管理 */}
            <Link
              href={"/bookshelf"}
              className="p-6 bg-white shadow-md rounded-lg flex flex-col items-center text-center shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95"
            >
              <Library className="w-12 h-12 text-green-500" />
              <h2 className="text-xl font-semibold mt-3">本棚管理</h2>
              <p className="text-gray-600 mt-2">
                購入済や読書中の本を整理できます。
              </p>
            </Link>

            {/* カード：読書進捗 */}
            <Link
              href={"/recommendation"}
              className="p-6 bg-white shadow-md rounded-lg flex flex-col items-center text-center shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95"
            >
              <Bot className="w-12 h-12 text-blue-600" />
              <h2 className="text-xl font-semibold mt-3">AIのおすすめ</h2>
              <p className="text-gray-600 mt-2">
                読書履歴からAIが本を推薦します。
              </p>
            </Link>
          </div>
        </div>
      ) : (
        <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-100 to-gray-300">
          <div className="bg-white p-10 rounded-2xl shadow-xl text-center max-w-md space-y-6">
            <h1 className="text-3xl font-bold text-gray-800">ようこそ 👋</h1>
            <p className="text-gray-600 text-sm">
              この機能を利用するにはログインが必要です。
            </p>
            <p className="text-gray-500 text-sm">
              アカウントをお持ちでない方は、新規登録を行ってください。
            </p>
            <Link
              href={"/login"}
              className="inline-block bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition"
            >
              ログインはこちらから
            </Link>
          </div>
        </div>
      )}
    </div>
  );
}

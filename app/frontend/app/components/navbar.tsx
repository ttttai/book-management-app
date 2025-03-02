"use client";

import { useRouter } from "next/navigation";
import { usePathname } from "next/navigation";
import { Bot, CircleUserRound, Library, Search } from "lucide-react";

export default function Navbar() {
  const router = useRouter();
  const pathname = usePathname();

  // 現在のページのボタンに適用するクラス
  const activeClass = "text-blue-600 font-bold border-b-2 border-blue-600";
  const baseClass = "text-gray-700 hover:text-blue-500 transition";

  return (
    <nav className="fixed top-0 w-full bg-white shadow-md z-50">
      <div className="max-w-6xl mx-auto flex items-center justify-between px-6 py-4">
        {/* 左側：ロゴ */}
        <div
          className="text-xl font-bold cursor-pointer text-gray-700 hover:text-gray-900 hover:bg-gray-100 p-2 rounded-md transition scale-100 active:scale-95"
          onClick={() => router.push("/")}
        >
          📚 MyLibrary
        </div>

        {/* 右側：メニュー */}
        <div className="flex space-x-6">
          <button
            className={`${pathname === "/search" ? activeClass : baseClass}`}
            onClick={() => router.push("/search")}
          >
            <div className="flex gap-1">
              <Search />
              本を検索
            </div>
          </button>
          <button
            className={`${pathname === "/bookshelf" ? activeClass : baseClass}`}
            onClick={() => router.push("/bookshelf")}
          >
            <div className="flex gap-1">
              <Library />
              マイ本棚
            </div>
          </button>
          <button
            className={`${
              pathname === "/recommendation" ? activeClass : baseClass
            }`}
            onClick={() => router.push("/recommendation")}
          >
            <div className="flex gap-1">
              <Bot />
              AIのおすすめ
            </div>
          </button>
          <button
            className={`${pathname === "/mypage" ? activeClass : baseClass}`}
            onClick={() => router.push("/mypage")}
          >
            <div className="flex gap-1">
              <CircleUserRound />
              マイページ
            </div>
          </button>
        </div>
      </div>
    </nav>
  );
}

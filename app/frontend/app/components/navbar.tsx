"use client";

import { useRouter } from "next/navigation";
import { usePathname } from "next/navigation";
import { Bot, CircleUserRound, Library, Search } from "lucide-react";
import Link from "next/link";

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
        <Link
          href={"/"}
          className="text-xl font-bold cursor-pointer text-gray-700 hover:text-gray-900 hover:bg-gray-100 p-2 rounded-md transition scale-100 active:scale-95"
        >
          📚 MyLibrary
        </Link>

        {/* 右側：メニュー */}
        <div className="flex space-x-6">
          <Link
            href={"/search"}
            className={`${pathname === "/search" ? activeClass : baseClass}`}
          >
            <div className="flex gap-1">
              <Search />
              本を検索
            </div>
          </Link>
          <Link
            href={"/bookshelf"}
            className={`${pathname === "/bookshelf" ? activeClass : baseClass}`}
          >
            <div className="flex gap-1">
              <Library />
              マイ本棚
            </div>
          </Link>
          <Link
            href={"/recommendation"}
            className={`${
              pathname === "/recommendation" ? activeClass : baseClass
            }`}
          >
            <div className="flex gap-1">
              <Bot />
              AIのおすすめ
            </div>
          </Link>
          <Link
            href={"/mypage"}
            className={`${pathname === "/mypage" ? activeClass : baseClass}`}
          >
            <div className="flex gap-1">
              <CircleUserRound />
              マイページ
            </div>
          </Link>
        </div>
      </div>
    </nav>
  );
}

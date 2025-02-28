"use client";

import { useRouter } from "next/navigation";
import Header from "../components/header";

// type Book = {
//   id: number;
//   isbn: number;
//   title_name: string;
//   title_name_kana: string;
//   publisher_name: string;
//   publisher_name_kana: string;
//   publish_date: string;
//   price: number;
//   status: number;
// };

// type Author = {
//   id: number;
//   name: string;
//   name_kana: string;
// };

// type Subject = {
//   id: number;
//   subject_name: string;
//   subject_kana: string;
// };

// type BookInfo = {
//   book: Book;
//   authors: Author[];
//   subjects: Subject[];
// };

export default function BookshelfPage() {
  const router = useRouter();
  //   let data = await fetch("http://localhost:8080/book/1");
  //   let bookInfo: BookInfo = data.json();

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

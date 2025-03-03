// import { useEffect, useState } from "react";
import { Bot } from "lucide-react";
// import Loading from "./loading";
import BookInfoDisplay from "../components/bookInfoDisplay";
import camelcaseKeys from "camelcase-keys";

export default async function BookRecommendations() {
  let bookInfo: BookInfo[] = [];

  try {
    const apiUrl = `${process.env.API_URL}/book/gemini`;
    const res = await fetch(apiUrl, { cache: "force-cache" });
    const data = await res.json();
    const dataCamel = camelcaseKeys(data, { deep: true });

    if (!res.ok) {
      console.error("Failed to fetch book");
    } else {
      bookInfo = dataCamel;
    }
  } catch (err: any) {
    console.error("Error fetching books:", err.message);
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-green-50 to-blue-100 flex flex-col items-center justify-center px-6 py-12">
      <div className="flex items-center space-x-3 mb-6 mt-20">
        <Bot className="w-8 h-8 text-blue-600" />
        <h1 className="text-3xl font-bold text-gray-800">AIのおすすめ</h1>
      </div>
      <p className="text-gray-600 mb-6 text-center">
        AIがあなたの読書履歴からおすすめの本を選びました
      </p>

      <div className="grid grid-cols-4 gap-4">
        {bookInfo.map((bookInfoItem) => (
          <BookInfoDisplay
            key={bookInfoItem.book.id}
            bookInfoItem={bookInfoItem}
          />
        ))}
      </div>
    </div>
  );
}

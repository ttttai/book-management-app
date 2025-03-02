"use client";

import { useEffect, useState } from "react";
import { Bot } from "lucide-react";
import Loading from "./loading";
import BookInfoDisplay from "../components/bookInfoDisplay";

export default function BookRecommendations() {
  const [bookInfo, setBookInfo] = useState<BookInfo[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const fetchBook = async () => {
    try {
      setIsLoading(true);
      const res = await fetch(`/api/book/gemini`);
      if (!res.ok) {
        console.error("Failed to fetch book");
        return;
      }
      const data = await res.json();
      setBookInfo(data);
      setIsLoading(false);
    } catch (err: any) {
      console.log(err.message);
    }
  };

  useEffect(() => {
    fetchBook();
  }, []);

  return (
    <div className="min-h-screen bg-gradient-to-br from-green-50 to-blue-100 flex flex-col items-center justify-center px-6 py-12">
      <div className="flex items-center space-x-3 mb-6 mt-20">
        <Bot className="w-8 h-8 text-blue-600" />
        <h1 className="text-3xl font-bold text-gray-800">AIのおすすめ</h1>
      </div>
      <p className="text-gray-600 mb-6 text-center">
        AIがあなたの読書履歴からおすすめの本を選びました
      </p>

      {/* 本のリスト */}
      <div className="flex items-center justify-center">
        <div className="flex items-center justify-center px-30 pt-5">
          <div>
            {isLoading ? (
              <Loading />
            ) : (
              <div>
                <div className="grid grid-cols-4 gap-4">
                  {bookInfo.map((bookInfoItem) => (
                    <BookInfoDisplay
                      key={bookInfoItem.book.id}
                      bookInfoItem={bookInfoItem}
                    />
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

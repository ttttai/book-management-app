"use client";

import { useState, Suspense, useRef } from "react";
import Header from "../components/header";
import Loading from "./loading";

export default function SearchPage() {
  const [bookInfo, setBookInfo] = useState<BookInfo[]>([]);
  const [query, setQuery] = useState("");
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const isComposing = useRef(false);

  const searchBooks = async () => {
    try {
      setIsLoading(true);
      const res = await fetch(`/api/book/search?title=${query}`);
      if (!res.ok) {
        throw new Error("Failed to fetch books");
      }
      const bookInfo = await res.json();
      if (bookInfo == null) {
        setBookInfo([]);
      } else {
        setBookInfo(bookInfo);
      }
      setIsLoading(false);
    } catch (err: any) {
      console.log(err.message);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" && !isComposing.current) {
      searchBooks();
    }
  };

  const handleCompositionStart = () => {
    isComposing.current = true;
  };

  const handleCompositionEnd = () => {
    isComposing.current = false;
  };

  return (
    <div>
      <Header title="検索" />
      <div>
        <div className="flex items-center justify-center">
          <input
            type="text"
            placeholder="検索..."
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onKeyDown={handleKeyDown}
            onCompositionStart={handleCompositionStart}
            onCompositionEnd={handleCompositionEnd}
            className="border p-2 my-8 rounded-md w-200 h-15 bg-white"
          />
          <button
            className="w-24 h-12 mx-4 bg-blue-500 text-white text-lg font-semibold rounded-lg shadow-md 
            hover:bg-blue-600 active:bg-blue-700 
            transform active:scale-90 transition-all duration-150"
            onClick={searchBooks}
          >
            検索
          </button>
        </div>
        <div className="h-screen flex items-center justify-center">
          <div>
            {isLoading ? (
              <Loading />
            ) : (
              <div>
                {bookInfo.length == 0 ? (
                  <div>見つかりませんでした</div>
                ) : (
                  <ul>
                    {bookInfo.map((bookInfoItem) => (
                      <li key={bookInfoItem.book.id}>
                        <div>ID: {bookInfoItem.book.id}</div>
                        <div>ISBN: {bookInfoItem.book.isbn}</div>
                        <div>タイトル: {bookInfoItem.book.titleName}</div>
                      </li>
                    ))}
                  </ul>
                )}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

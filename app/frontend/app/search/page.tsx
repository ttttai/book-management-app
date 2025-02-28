"use client";

import { useState, useRef } from "react";
import Loading from "./loading";
import { useRouter } from "next/navigation";
import { BOOK_STATUS_LABELS, BOOK_STATUS_COLORS } from "../constants";

export default function SearchPage() {
  const [bookInfo, setBookInfo] = useState<BookInfo[]>([]);
  const [query, setQuery] = useState("");
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isFirstVisit, setIsFirstVisit] = useState<boolean>(true);
  const isComposing = useRef(false);
  const router = useRouter();

  const searchBooks = async () => {
    try {
      setIsLoading(true);
      setIsFirstVisit(false);
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
      setIsLoading(false);
      alert(err.message);
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
      <div className="flex items-center justify-center">
        <input
          type="text"
          placeholder="検索..."
          maxLength={40}
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
      <div className="flex items-center justify-center px-30 pt-10">
        <div>
          {isLoading ? (
            <Loading />
          ) : (
            <div>
              {bookInfo.length == 0 ? (
                <div>{!isFirstVisit && "見つかりませんでした"}</div>
              ) : (
                <div className="grid grid-cols-4 gap-4">
                  {bookInfo.map((bookInfoItem) => (
                    <div
                      key={bookInfoItem.book.id}
                      className="flex flex-col items-center border border-gray-300 p-4 rounded-lg shadow-md bg-white cursor-pointer 
                        hover:shadow-xl hover:scale-105 hover:border-blue-500 hover:bg-gray-100 
                        transition-all duration-200"
                      onClick={() =>
                        router.push(`/bookshelf/book/${bookInfoItem.book.id}`)
                      }
                    >
                      <div className="aspect-[2/3] w-full max-w-xs overflow-hidden rounded-md">
                        <img
                          src={`${process.env.NEXT_PUBLIC_THUMBNAIL_URL}/${bookInfoItem.book.isbn}.jpg`}
                          alt={bookInfoItem.book.titleName}
                          className="w-full h-full object-contain"
                        />
                      </div>
                      <div className="mt-3 text-center flex flex-col flex-grow">
                        <div className="text-md font-semibold">
                          {bookInfoItem.book.titleName}
                        </div>
                        {bookInfoItem.authors ? (
                          <div>
                            {bookInfoItem.authors
                              .map((author) => author.name)
                              .join(", ")}
                          </div>
                        ) : (
                          <div></div>
                        )}
                        <div className="text-sm text-gray-600">
                          {bookInfoItem.book.publisherName}
                        </div>
                        <div className="text-sm text-gray-600">
                          ISBN: {bookInfoItem.book.isbn}
                        </div>
                      </div>
                      <div className="mt-auto">
                        <span
                          className={`inline-block px-3 py-1 rounded-md text-white ${
                            BOOK_STATUS_COLORS[bookInfoItem.book.status]
                          }`}
                        >
                          {BOOK_STATUS_LABELS[bookInfoItem.book.status]}
                        </span>
                      </div>
                    </div>
                  ))}
                </div>
              )}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

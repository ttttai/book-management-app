"use client";

import { useState } from "react";
import Loading from "./loading";
import BookInfoDisplay from "../components/bookInfoDisplay";
import SearchBox from "../components/searchBox";

export default function SearchPage() {
  const [bookInfo, setBookInfo] = useState<BookInfo[]>([]);
  const [query, setQuery] = useState("");
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isFirstVisit, setIsFirstVisit] = useState<boolean>(true);

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
    }
  };

  return (
    <div>
      <div className="flex items-center justify-center">
        <SearchBox setQuery={setQuery} searchBooks={searchBooks} />
        <button
          className="w-24 h-12 mx-4 bg-blue-500 text-white text-lg font-semibold rounded-lg shadow-md 
            hover:bg-blue-600 active:bg-blue-700 
            transform active:scale-90 transition-all duration-150"
          onClick={searchBooks}
        >
          検索
        </button>
      </div>
      <div className="flex items-center justify-center px-30 pt-5">
        <div>
          {isLoading ? (
            <Loading />
          ) : (
            <div>
              {bookInfo.length == 0 ? (
                <div>
                  {!isFirstVisit && (
                    <div className="text-lg">見つかりませんでした</div>
                  )}
                </div>
              ) : (
                <div className="grid grid-cols-4 gap-4">
                  {bookInfo.map((bookInfoItem) => (
                    <BookInfoDisplay
                      key={bookInfoItem.book.id}
                      bookInfoItem={bookInfoItem}
                    />
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

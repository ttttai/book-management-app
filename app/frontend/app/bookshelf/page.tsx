"use client";

import { useEffect, useState } from "react";
import BookInfoDisplay from "../components/bookInfoDisplay";
import Loading from "./loading";
import SearchBox from "../components/searchBox";
import {
  BOOK_STATUS_PURCHASED,
  BOOK_STATUS_READING,
  BOOK_STATUS_READ_COMPLETED,
  BOOK_STATUS_COLORS,
  BOOK_STATUS_LABELS,
} from "../constants";
import { HelpCircle } from "lucide-react";
import { useSearchParams } from "next/navigation";

export default function BookshelfPage() {
  const searchParams = useSearchParams();
  const status = searchParams.get("status");
  const [bookInfo, setBookInfo] = useState<BookInfo[]>([]);
  const [query, setQuery] = useState("");
  const SELECTABLE_BOOK_STATUS = [
    BOOK_STATUS_PURCHASED,
    BOOK_STATUS_READING,
    BOOK_STATUS_READ_COMPLETED,
  ];
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isHovered, setIsHovered] = useState(false);

  let firstStatus: number[];
  if (status == null) {
    firstStatus = SELECTABLE_BOOK_STATUS;
  } else {
    firstStatus = [Number(status)];
  }
  const [selectedStatuses, setSelectedStatuses] =
    useState<number[]>(firstStatus);

  const handleStatusChange = (status: number) => {
    setSelectedStatuses((prev) =>
      prev.includes(status)
        ? prev.filter((s) => s !== status)
        : [...prev, status]
    );
  };

  useEffect(() => {
    fetchBooks();
  }, [selectedStatuses]);

  const fetchBooks = async () => {
    try {
      setIsLoading(true);
      if (selectedStatuses.length == 0) {
        setBookInfo([]);
      } else {
        const res = await fetch(
          `/api/book?title=${query}&status=${selectedStatuses}`
        );
        if (!res.ok) {
          throw new Error("Failed to fetch books");
        }

        const bookInfo = await res.json();
        if (bookInfo == null) {
          setBookInfo([]);
        } else {
          setBookInfo(bookInfo);
        }
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
        <SearchBox setQuery={setQuery} searchBooks={fetchBooks} />
        <button
          className="w-24 h-12 mx-4 bg-blue-500 text-white text-lg font-semibold rounded-lg shadow-md 
                    hover:bg-blue-600 active:bg-blue-700 
                    transform active:scale-90 transition-all duration-150"
          onClick={fetchBooks}
        >
          検索
        </button>
      </div>
      <div className="flex flex-col items-center">
        <div className="flex flex-wrap mb-4 justify-center">
          <span className="font-semibold text-gray-700 py-1">ステータス</span>
          <HelpCircle
            className="w-5 h-5 text-gray-500 hover:text-gray-700 cursor-pointer"
            onMouseEnter={() => setIsHovered(true)}
            onMouseLeave={() => setIsHovered(false)}
          />
          {isHovered && (
            <div className="absolute top-30 w-56">
              <div className="bg-gray-900 text-white text-sm p-2 rounded-lg shadow-md text-center">
                ⓘ ステータスをクリックして本を絞り込めます
              </div>
            </div>
          )}
          <span className="font-semibold text-gray-700 py-1">：</span>
          {SELECTABLE_BOOK_STATUS.map((value) => (
            <button
              key={value}
              className={`ml-3 px-3 py-1 rounded-md text-white font-semibold transition-all duration-150 
                ${
                  selectedStatuses.includes(Number(value))
                    ? `${BOOK_STATUS_COLORS[value]} border-gray-800 shadow-lg scale-105 cursor-pointer`
                    : "bg-gray-300 hover:bg-gray-400 border-gray-500 cursor-pointer"
                } 
                active:scale-95`}
              onClick={() => handleStatusChange(Number(value))}
            >
              {BOOK_STATUS_LABELS[value]}
            </button>
          ))}
        </div>
      </div>
      <div className="flex items-center justify-center">
        検索結果：{bookInfo.length}件
      </div>
      <div className="flex items-center justify-center">
        <div className="flex items-center justify-center px-30 pt-5">
          <div>
            {isLoading ? (
              <Loading />
            ) : (
              <div>
                {bookInfo.length == 0 ? (
                  <div className="text-lg">見つかりませんでした</div>
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
    </div>
  );
}

"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import {
  BOOK_STATUS_PURCHASED,
  BOOK_STATUS_READING,
  BOOK_STATUS_READ_COMPLETED,
  BOOK_STATUS_LABELS,
} from "../constants";

export default function MyPage() {
  const router = useRouter();
  const [bookInfo, setBookInfo] = useState<BookInfo[]>([]);
  const [statusCounts, setStatusCounts] = useState({
    [BOOK_STATUS_PURCHASED]: 0,
    [BOOK_STATUS_READING]: 0,
    [BOOK_STATUS_READ_COMPLETED]: 0,
  });

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const status = [
          BOOK_STATUS_PURCHASED,
          BOOK_STATUS_READING,
          BOOK_STATUS_READ_COMPLETED,
        ];
        const res = await fetch(`/api/book?status=${status}`);
        if (!res.ok) {
          throw new Error("Failed to fetch books");
        }
        const data: BookInfo[] = await res.json();
        setBookInfo(data);

        const counts = {
          [BOOK_STATUS_PURCHASED]: data.filter(
            (book) => book.book.status === BOOK_STATUS_PURCHASED
          ).length,
          [BOOK_STATUS_READING]: data.filter(
            (book) => book.book.status === BOOK_STATUS_READING
          ).length,
          [BOOK_STATUS_READ_COMPLETED]: data.filter(
            (book) => book.book.status === BOOK_STATUS_READ_COMPLETED
          ).length,
        };
        setStatusCounts(counts);
      } catch (error) {
        console.error(error);
      }
    };

    fetchBooks();
  }, []);

  return (
    <div className="p-16">
      <div className="max-w-4xl mx-auto p-6 bg-white shadow-md rounded-lg">
        <h1 className="text-2xl font-bold mb-4">マイページ</h1>

        <div className="grid grid-cols-3 gap-4 text-center">
          <div
            className="p-4 bg-yellow-100 rounded-lg shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95"
            onClick={() =>
              router.push(`/bookshelf?status=${BOOK_STATUS_PURCHASED}`)
            }
          >
            <p className="text-lg font-semibold">
              {BOOK_STATUS_LABELS[BOOK_STATUS_PURCHASED]}
            </p>
            <p className="text-xl font-bold">
              {statusCounts[BOOK_STATUS_PURCHASED]}冊
            </p>
          </div>

          <div
            className="p-4 bg-blue-100 rounded-lg shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95"
            onClick={() =>
              router.push(`/bookshelf?status=${BOOK_STATUS_READING}`)
            }
          >
            <p className="text-lg font-semibold">
              {BOOK_STATUS_LABELS[BOOK_STATUS_READING]}
            </p>
            <p className="text-xl font-bold">
              {statusCounts[BOOK_STATUS_READING]}冊
            </p>
          </div>

          <div
            className="p-4 bg-green-100 rounded-lg shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95"
            onClick={() =>
              router.push(`/bookshelf?status=${BOOK_STATUS_READ_COMPLETED}`)
            }
          >
            <p className="text-lg font-semibold">
              {BOOK_STATUS_LABELS[BOOK_STATUS_READ_COMPLETED]}
            </p>
            <p className="text-xl font-bold">
              {statusCounts[BOOK_STATUS_READ_COMPLETED]}冊
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

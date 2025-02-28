"use client";

import { useState, useEffect } from "react";
import { useParams } from "next/navigation";

export default function BookDetail({ params }: { params: { id: string } }) {
  const BOOK_STATUS_NOT_PURCHASED = 0;
  const BOOK_STATUS_PURCHASED = 1;
  const BOOK_STATUS_READING = 2;
  const BOOK_STATUS_READ_COMPLETED = 3;

  const { id } = useParams();
  const [bookInfo, setBook] = useState<BookInfo | null>(null);

  useEffect(() => {
    const fetchBook = async () => {
      const res = await fetch(`/api/book/${id}`);
      if (!res.ok) {
        console.error("Failed to fetch book");
        return;
      }
      const data = await res.json();
      setBook(data);
    };

    fetchBook();
  }, []);

  return (
    <div className="min-h-screen pt-16">
      <div className="max-w-4xl mx-auto p-6 bg-white shadow-md rounded-lg">
        {/* 本の情報 */}
        <div className="flex flex-col md:flex-row">
          {/* 左側: 本の画像 */}
          <div className="w-full h-full md:w-1/3 flex justify-center">
            <img
              src={`${process.env.NEXT_PUBLIC_THUMBNAIL_URL}/${bookInfo?.book.isbn}.jpg`}
              alt={bookInfo?.book.titleName}
              className="rounded-md shadow-md"
            />
          </div>

          {/* 右側: 詳細情報 */}
          <div className="md:w-2/3 mt-4 md:mt-0 md:ml-6">
            <h1 className="text-2xl font-bold">{bookInfo?.book.titleName}</h1>
            <p className="text-gray-600">{bookInfo?.book.titleNameKana}</p>

            {/* 著者 */}
            <p className="mt-2 text-gray-800">
              <span className="font-semibold">著者:</span>{" "}
              {bookInfo?.authors.map((author) => author.name).join(", ")}
            </p>

            {/* 出版社 */}
            <p className="mt-1 text-gray-800">
              <span className="font-semibold">出版社:</span>{" "}
              {bookInfo?.book.publisherName}
            </p>

            {/* 価格 */}
            {bookInfo?.book.price != 0 && (
              <p className="mt-1 text-gray-800">
                <span className="font-semibold">価格:</span> ¥
                {bookInfo?.book.price.toLocaleString()}
              </p>
            )}

            {/* ジャンル */}
            {bookInfo?.subjects.length != 0 && (
              <p className="mt-1 text-gray-800">
                <span className="font-semibold">ジャンル:</span>{" "}
                {bookInfo?.subjects
                  .map((subject) => subject.subjectName)
                  .join(", ")}
              </p>
            )}

            {/* ISBN */}
            <p className="mt-1 text-gray-800">
              <span className="font-semibold">ISBN:</span> {bookInfo?.book.isbn}
            </p>

            {/* ステータス */}
            <div className="mt-3">
              <span className="font-semibold">ステータス:</span>{" "}
              <span
                className={`inline-block px-3 py-1 rounded-md text-white ${
                  bookInfo?.book.status === BOOK_STATUS_NOT_PURCHASED
                    ? "bg-gray-500"
                    : bookInfo?.book.status === BOOK_STATUS_PURCHASED
                    ? "bg-yellow-500"
                    : bookInfo?.book.status === BOOK_STATUS_READING
                    ? "bg-blue-500"
                    : bookInfo?.book.status === BOOK_STATUS_READ_COMPLETED
                    ? "bg-green-500"
                    : ""
                }`}
              >
                {bookInfo?.book.status === BOOK_STATUS_NOT_PURCHASED
                  ? "未購入"
                  : bookInfo?.book.status === BOOK_STATUS_PURCHASED
                  ? "積読中"
                  : bookInfo?.book.status === BOOK_STATUS_READING
                  ? "読んでいる"
                  : bookInfo?.book.status === BOOK_STATUS_READ_COMPLETED
                  ? "読んだ"
                  : ""}
              </span>
            </div>

            {/* ボタン */}
            <div className="mt-5 flex gap-3">
              <button className="px-4 py-2 bg-blue-600 text-white rounded-md shadow-md hover:bg-blue-700">
                本棚に追加
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

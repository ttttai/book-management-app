"use client";

import { useState, useEffect } from "react";
import { useParams } from "next/navigation";
import { toast } from "react-toastify";
import {
  BOOK_STATUS_NOT_PURCHASED,
  BOOK_STATUS_PURCHASED,
  BOOK_STATUS_READING,
  BOOK_STATUS_READ_COMPLETED,
  BOOK_STATUS_LABELS,
} from "../../../constants";

export default async function BookDetail({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const BOOK_STATUSES = [
    {
      value: BOOK_STATUS_NOT_PURCHASED,
      label: BOOK_STATUS_LABELS[BOOK_STATUS_NOT_PURCHASED],
    },
    {
      value: BOOK_STATUS_PURCHASED,
      label: BOOK_STATUS_LABELS[BOOK_STATUS_PURCHASED],
    },
    {
      value: BOOK_STATUS_READING,
      label: BOOK_STATUS_LABELS[BOOK_STATUS_READING],
    },
    {
      value: BOOK_STATUS_READ_COMPLETED,
      label: BOOK_STATUS_LABELS[BOOK_STATUS_READ_COMPLETED],
    },
  ];

  const { id } = await params;
  const [bookInfo, setBookInfo] = useState<BookInfo | null>(null);
  const [selectedStatus, setSelectedStatus] = useState<number>(0);
  const [readingStartDate, setReadingStartDate] = useState<string>("");
  const [readingEndDate, setReadingEndDate] = useState<string>("");

  useEffect(() => {
    const fetchBook = async () => {
      const res = await fetch(`/api/book/${id}`);
      if (!res.ok) {
        console.error("Failed to fetch book");
        return;
      }
      const data = await res.json();
      setBookInfo(data);
      setSelectedStatus(data.book.status);
      if (data.book.readingStartDate != null) {
        setReadingStartDate(data.book.readingStartDate.split("T")[0]);
      }
      if (data.book.readingEndDate != null) {
        setReadingEndDate(data.book.readingEndDate.split("T")[0]);
      }
    };

    fetchBook();
  }, []);

  const updateStatus = async () => {
    try {
      const res = await fetch(`/api/book/${id}/status`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          status: selectedStatus,
        }),
      });
      if (!res.ok) {
        console.error("Failed to update status");
        return;
      }
      const data = await res.json();
      const newBookInfo: BookInfo = {
        book: data,
        authors: bookInfo?.authors ?? [],
        subjects: bookInfo?.subjects ?? [],
      };
      setBookInfo(newBookInfo);
      toast.success("ステータスを更新しました！", {
        position: "top-right",
        autoClose: 2000,
      });
    } catch (error: any) {
      console.log(error.message);
    }
  };

  const updateReadingStartDate = async () => {
    try {
      if (bookInfo == null) {
        return;
      }
      let requestBody = bookInfo.book;
      requestBody.readingStartDate = readingStartDate;

      const res = await fetch(`/api/book/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(requestBody),
      });
      if (!res.ok) {
        console.error("Failed to update status");
        return;
      }
      const data = await res.json();
      const newBookInfo: BookInfo = {
        book: data,
        authors: bookInfo?.authors ?? [],
        subjects: bookInfo?.subjects ?? [],
      };
      setBookInfo(newBookInfo);
      toast.success("読書開始日を更新しました！", {
        position: "top-right",
        autoClose: 2000,
      });
    } catch (error: any) {
      console.log(error.message);
    }
  };

  const updateReadingEndDate = async () => {
    try {
      if (bookInfo == null) {
        return;
      }
      let requestBody = bookInfo.book;
      requestBody.readingEndDate = readingEndDate;

      const res = await fetch(`/api/book/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(requestBody),
      });
      if (!res.ok) {
        console.error("Failed to update status");
        return;
      }
      const data = await res.json();
      const newBookInfo: BookInfo = {
        book: data,
        authors: bookInfo?.authors ?? [],
        subjects: bookInfo?.subjects ?? [],
      };
      setBookInfo(newBookInfo);
      toast.success("読書終了日を更新しました！", {
        position: "top-right",
        autoClose: 2000,
      });
    } catch (error: any) {
      console.log(error.message);
    }
  };

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

            {/* 読書開始日 */}
            <div className="mt-3 flex items-center">
              <span className="font-semibold">読書開始日:</span>
              <input
                type="date"
                value={readingStartDate}
                onChange={(e) => setReadingStartDate(e.target.value)}
                className="ml-2 p-2 border rounded-md bg-white"
              />
              <button
                className="ml-2 px-4 py-2 bg-blue-600 text-white rounded-md shadow-md hover:bg-blue-700"
                onClick={updateReadingStartDate}
              >
                更新
              </button>
            </div>

            {/* 読書終了日 */}
            <div className="mt-3 flex items-center">
              <span className="font-semibold">読書終了日:</span>
              <input
                type="date"
                value={readingEndDate}
                onChange={(e) => setReadingEndDate(e.target.value)}
                className="ml-2 p-2 border rounded-md bg-white"
              />
              <button
                className="ml-2 px-4 py-2 bg-blue-600 text-white rounded-md shadow-md hover:bg-blue-700"
                onClick={updateReadingEndDate}
              >
                更新
              </button>
            </div>

            <div className="mt-3">
              <span className="font-semibold">ステータス:</span>{" "}
              <select
                value={selectedStatus}
                onChange={(e) => setSelectedStatus(Number(e.target.value))}
                className="ml-2 p-2 border rounded-md bg-white"
              >
                {BOOK_STATUSES.map((status) => (
                  <option key={status.value} value={status.value}>
                    {status.label}
                  </option>
                ))}
              </select>
              <button
                className="mx-4 px-4 py-2 bg-blue-600 text-white rounded-md shadow-md hover:bg-blue-700"
                onClick={updateStatus}
              >
                更新
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

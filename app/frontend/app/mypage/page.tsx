import Link from "next/link";
import {
  BOOK_STATUS_PURCHASED,
  BOOK_STATUS_READING,
  BOOK_STATUS_READ_COMPLETED,
  BOOK_STATUS_LABELS,
} from "../constants";

export default async function MyPage() {
  const bookStatuses = [
    BOOK_STATUS_PURCHASED,
    BOOK_STATUS_READING,
    BOOK_STATUS_READ_COMPLETED,
  ];
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/book?status=${bookStatuses}`
  );
  const booksData: BookInfo[] = await res.json();

  const purchasedBooks = booksData.filter(
    (book) => book.book.status === BOOK_STATUS_PURCHASED
  );
  const readingBooks = booksData.filter(
    (book) => book.book.status === BOOK_STATUS_READING
  );
  const readCompletedBooks = booksData.filter(
    (book) => book.book.status === BOOK_STATUS_READ_COMPLETED
  );

  return (
    <div className="p-16">
      <div className="max-w-4xl mx-auto p-6 bg-white shadow-md rounded-lg">
        <h1 className="text-2xl font-bold mb-4">マイページ</h1>

        <div className="grid grid-cols-3 gap-4 text-center">
          <Link href={`/bookshelf?status=${BOOK_STATUS_PURCHASED}`}>
            <div className="p-4 bg-yellow-100 rounded-lg shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95">
              <p className="text-lg font-semibold">
                {BOOK_STATUS_LABELS[BOOK_STATUS_PURCHASED]}
              </p>
              <p className="text-xl font-bold">{purchasedBooks.length}冊</p>
            </div>
          </Link>

          <Link href={`/bookshelf?status=${BOOK_STATUS_READING}`}>
            <div className="p-4 bg-blue-100 rounded-lg shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95">
              <p className="text-lg font-semibold">
                {BOOK_STATUS_LABELS[BOOK_STATUS_READING]}
              </p>
              <p className="text-xl font-bold">{readingBooks.length}冊</p>
            </div>
          </Link>

          <Link href={`/bookshelf?status=${BOOK_STATUS_READ_COMPLETED}`}>
            <div className="p-4 bg-green-100 rounded-lg shadow-md cursor-pointer transition transform hover:scale-105 hover:shadow-lg active:scale-95">
              <p className="text-lg font-semibold">
                {BOOK_STATUS_LABELS[BOOK_STATUS_READ_COMPLETED]}
              </p>
              <p className="text-xl font-bold">{readCompletedBooks.length}冊</p>
            </div>
          </Link>
        </div>
      </div>
    </div>
  );
}

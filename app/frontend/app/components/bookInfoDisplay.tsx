import { useRouter } from "next/navigation";
import { BOOK_STATUS_LABELS, BOOK_STATUS_COLORS } from "../constants";

type BookInfoDisplayProps = {
  bookInfoItem: BookInfo;
};

export default function BookInfoDisplay({
  bookInfoItem,
}: BookInfoDisplayProps) {
  const router = useRouter();
  return (
    <div
      key={bookInfoItem.book.id}
      className="flex flex-col items-center border border-gray-300 p-4 rounded-lg shadow-md bg-white cursor-pointer
                          hover:shadow-xl hover:scale-105 hover:border-blue-500 hover:bg-gray-100 
                          transition-all duration-200"
      onClick={() => router.push(`/bookshelf/book/${bookInfoItem.book.id}`)}
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
            {bookInfoItem.authors.map((author) => author.name).join(", ")}
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
  );
}

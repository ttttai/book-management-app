import { Bot } from "lucide-react";
import BookInfoDisplay from "../components/bookInfoDisplay";

export default async function BookRecommendations() {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/book/gemini`,
    { cache: "no-store" }
  );
  const bookInfo: BookInfo[] = await res.json();

  return (
    <div className="min-h-screen bg-gradient-to-br from-green-50 to-blue-100 flex flex-col items-center justify-center px-6 py-12">
      <div className="flex items-center space-x-3 mb-6 mt-20">
        <Bot className="w-8 h-8 text-blue-600" />
        <h1 className="text-3xl font-bold text-gray-800">AIのおすすめ</h1>
      </div>
      <p className="text-gray-600 mb-6 text-center">
        AIがあなたの読書履歴からおすすめの本を選びました
      </p>

      {bookInfo.length === 0 ? (
        <p className="text-lg text-gray-600">
          おすすめの本が見つかりませんでした。
        </p>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          {bookInfo.map((bookInfoItem) => (
            <BookInfoDisplay
              key={bookInfoItem.book.id}
              bookInfoItem={bookInfoItem}
            />
          ))}
        </div>
      )}
    </div>
  );
}

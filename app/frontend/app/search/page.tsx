"use client";

import { useState, useEffect, useCallback, useRef } from "react";
import Loading from "./loading";
import BookInfoDisplay from "../components/bookInfoDisplay";
import SearchBox from "../components/searchBox";

export default function SearchPage() {
  const [bookInfo, setBookInfo] = useState<BookInfo[]>([]);
  const [query, setQuery] = useState("");
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isFirstVisit, setIsFirstVisit] = useState<boolean>(true);
  const [offset, setOffset] = useState(1);
  const loader = useRef<HTMLDivElement | null>(null);
  const [prevQuery, setPrevQuery] = useState("");

  const searchBooks = async () => {
    try {
      setIsLoading(true);
      setIsFirstVisit(false);
      setPrevQuery(query);

      if (prevQuery != query) {
        setBookInfo([]);
      }

      const res = await fetch(
        `/api/book/search?title=${query}&offset=${offset}`
      );
      if (!res.ok) {
        throw new Error("Failed to fetch books");
      }

      const data = await res.json();
      if (data != null) {
        if (prevQuery == query) {
          setBookInfo((prev) => [...prev, ...data]);
        } else {
          setBookInfo(data);
        }
      }

      setIsLoading(false);
    } catch (err: any) {
      console.log(err.message);
      setIsLoading(false);
    }
  };

  // targetと交差すると発火し、offsetを変更する
  const handleObserver = useCallback(
    (entities: IntersectionObserverEntry[]) => {
      const target = entities[0];
      if (target.isIntersecting) {
        setOffset((prev) => prev + 30);
      }
    },
    []
  );

  useEffect(() => {
    const observer = new IntersectionObserver(handleObserver, {
      root: null,
      rootMargin: "20px",
      threshold: 0.9,
    });
    if (loader.current) observer.observe(loader.current);

    return () => {
      observer.disconnect();
    };
  }, [handleObserver]);

  useEffect(() => {
    if (bookInfo.length != 0) {
      searchBooks();
    }
  }, [offset]);

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
          <div>
            {bookInfo.length == 0 ? (
              <div>
                {!isFirstVisit && !isLoading && (
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
          {isLoading && <Loading />}
        </div>
      </div>
      <div ref={loader}></div>
    </div>
  );
}

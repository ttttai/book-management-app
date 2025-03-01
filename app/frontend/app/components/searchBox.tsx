"use client";

import { useRef } from "react";
import { Dispatch, SetStateAction } from "react";

export default function SearchBox(props: {
  setQuery: Dispatch<SetStateAction<string>>;
  searchBooks: () => Promise<void>;
}) {
  const { setQuery, searchBooks } = props;
  const isComposing = useRef(false);

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
    <input
      type="text"
      placeholder="検索..."
      maxLength={40}
      onChange={(e) => setQuery(e.target.value)}
      onKeyDown={handleKeyDown}
      onCompositionStart={handleCompositionStart}
      onCompositionEnd={handleCompositionEnd}
      className="border p-2 my-12 rounded-md w-200 h-15 bg-white"
    />
  );
}

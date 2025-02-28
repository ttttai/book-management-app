"use client";

import { useParams } from "next/navigation";
import Header from "../../../components/header";

export default function BookPage() {
  const { id } = useParams();

  return (
    <div>
      <Header title="本" />
      <div className="h-screen flex items-center justify-center bg-purple-100">
        <p>Hello Book!</p>
        <h2>ID : {id}</h2>
      </div>
    </div>
  );
}

import Header from "../components/header";

export default function SearchPage() {
  return (
    <div>
      <Header title="検索" />
      <div className="h-screen flex items-center justify-center bg-green-100">
        <h1 className="text-2xl font-bold">本を検索</h1>
      </div>
    </div>
  );
}

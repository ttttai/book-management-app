import camelcaseKeys from "camelcase-keys";
import { NextRequest, NextResponse } from "next/server";

export async function GET(req: NextRequest) {
  try {
    const { searchParams } = new URL(req.url);
    const title = searchParams.get("title") ?? "";
    const status = searchParams.getAll("status");
    if (status.length == 0) {
      throw new Error("Invalid Parameter Error");
    }
    const statusSplitted = status[0].split(",");
    const statusQuery = statusSplitted.map((s) => `status=${s}`).join("&");

    const apiUrl = `${process.env.API_URL}/book?title=${title}&${statusQuery}`;
    const res = await fetch(apiUrl);
    if (!res.ok) {
      throw new Error(`API Error: ${res.statusText}`);
    }

    const data = await res.json();
    const dataCamel = camelcaseKeys(data, { deep: true });
    return NextResponse.json(dataCamel, { status: 200 });
  } catch (error) {
    return NextResponse.json(
      { error: "Failed to fetch books" },
      { status: 500 }
    );
  }
}

import { NextRequest, NextResponse } from "next/server";
import camelcaseKeys from "camelcase-keys";
import snakecaseKeys from "snakecase-keys";

export async function GET(req: NextRequest) {
  try {
    const apiUrl = `${process.env.API_URL}/book/gemini`;
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

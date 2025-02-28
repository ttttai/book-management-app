import { NextRequest, NextResponse } from "next/server";
import camelcaseKeys from "camelcase-keys";

export async function GET(
  req: NextRequest,
  { params }: { params: Promise<{ id: string }> }
) {
  try {
    const id = (await params).id;
    if (!id) {
      return NextResponse.json(
        { error: "Missing required parameter: id" },
        { status: 400 }
      );
    }

    const apiUrl = `${process.env.API_URL}/book/${id}`;
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

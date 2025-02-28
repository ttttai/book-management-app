import camelcaseKeys from "camelcase-keys";
import { NextRequest, NextResponse } from "next/server";

export async function PUT(
  req: NextRequest,
  { params }: { params: Promise<{ id: string }> }
) {
  try {
    const id = (await params).id;
    const { status } = await req.json();

    if (typeof status !== "number" || isNaN(status)) {
      return NextResponse.json(
        { error: "Invalid status value" },
        { status: 400 }
      );
    }

    const apiUrl = `${process.env.API_URL}/book/${id}/status`;
    const res = await fetch(apiUrl, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        status: status,
      }),
    });
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

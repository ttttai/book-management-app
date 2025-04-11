import { NextResponse, NextRequest } from "next/server";
import { hash } from "bcryptjs";

export async function POST(req: NextRequest) {
  const { email, password } = await req.json();
  console.log("email", email);
  console.log("password", password);

  if (!email || !password) {
    return NextResponse.json(
      { message: "必須項目が不足しています" },
      { status: 400 }
    );
  }

  const hashedPassword = await hash(password, 10);
  const requestBody = {
    email: email,
    password: hashedPassword,
  };

  const apiUrl = `${process.env.API_URL}/authentication/register`;
  const res = await fetch(apiUrl, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(requestBody),
  });
  if (!res.ok) {
    return NextResponse.json(
      { message: "登録に失敗しました" },
      { status: 500 }
    );
  }

  return NextResponse.json({ message: "登録成功" }, { status: 200 });
}

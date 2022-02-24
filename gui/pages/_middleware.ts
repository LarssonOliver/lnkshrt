import { NextRequest, NextResponse } from "next/server";

export default function middleware(req: NextRequest) {
  const url = req.nextUrl.clone();

  if (url.pathname === "/path")
    return NextResponse.json({ path: process.env.API_URL });

  if (!url.pathname.startsWith("/api")) return NextResponse.next();

  // Rewrite all /api urls.
  const path = url.pathname.replace("/api", "");
  const rewrite = `${process.env.API_URL}${path}`;

  return NextResponse.rewrite(rewrite);
}

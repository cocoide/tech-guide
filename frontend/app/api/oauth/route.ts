import { ResponseCookie } from 'next/dist/compiled/@edge-runtime/cookies';
import { NextRequest, NextResponse } from 'next/server';

const corsHeaders = {
    'Access-Control-Allow-Origin': process.env.NEXT_PUBLIC_API_BASE_URL,
    'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
    'Access-Control-Allow-Headers': 'Content-Type',
}

const cookieOptions: Partial<ResponseCookie> = {
    httpOnly: true,
    domain: ".tech-guide.jp",
    maxAge: 60 * 60,// 1時間,
    secure: true,
    path: "/",
    sameSite: 'lax',
}

export async function GET(request: NextRequest) {
    const token = request.nextUrl.searchParams.get("token")
    if (token === null || token === "") {
        console.log("Failed to get token")
        return NextResponse.redirect("www.tech-guide.jp")
    }
    const response = NextResponse.redirect("www.tech-guide.jp")
    return response.cookies.set("accessToken", token, cookieOptions)
}
import { NextRequest, NextResponse } from 'next/server';

const corsHeaders = {
    'Access-Control-Allow-Origin': process.env.NEXT_PUBLIC_API_BASE_URL,
    'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
    'Access-Control-Allow-Headers': 'Content-Type',
}

export async function GET(request: NextRequest) {
    const token = request.nextUrl.searchParams.get("token")
    const response = NextResponse.redirect("https://www.tech-guide.jp")
    response.headers.set('Access-Control-Allow-Origin', process.env.NEXT_PUBLIC_API_BASE_URL)
    response.headers.set('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS')
    if (token === null || token === "") {
        console.log("Failed to get token")
        return response
    }
    if (token) {
        response.cookies.set({
            name: 'accessToken',
            domain: '.tech-guide.jp',
            value: token,
            httpOnly: true,
            sameSite: 'lax',
            maxAge: 60 * 60,// 1時間,
            path: '/',
            secure: true,
        })
    }
    return response
}
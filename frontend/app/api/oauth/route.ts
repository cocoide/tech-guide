import { NextRequest, NextResponse } from 'next/server';

const corsHeaders = {
    'Access-Control-Allow-Origin': process.env.NEXT_PUBLIC_API_BASE_URL,
    'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
    'Access-Control-Allow-Headers': 'Content-Type',
}

export async function GET(request: NextRequest) {
    const accessToken = request.nextUrl.searchParams.get("access")
    const refreshToken = request.nextUrl.searchParams.get("refresh")
    const response = NextResponse.redirect("https://www.tech-guide.jp")
    response.headers.set('Access-Control-Allow-Origin', process.env.NEXT_PUBLIC_API_BASE_URL)
    response.headers.set('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS')
    if (!accessToken||!refreshToken) {
        console.log("Failed to get tokens")
        return response
    }
        response.cookies.set({
            name: 'accessToken',
            domain: '.tech-guide.jp',
            value: accessToken,
            httpOnly: true,
            sameSite: 'lax',
            maxAge: 60 * 60 * 24 * 7, //7日
            path: '/',
            secure: true,
        })

    response.cookies.set({
        name: 'refreshToken',
        domain: '.tech-guide.jp',
        value: refreshToken,
        httpOnly: true,
        sameSite: 'lax',
        maxAge: 60 * 60 * 24 * 30,// 30日
        path: '/',
        secure: true,
    })

    return response
}
import { cookies } from 'next/headers';
import { NextRequest, NextResponse } from 'next/server';

export async function POST(request: NextRequest) {
    const token = request.nextUrl.searchParams.get("token")

    const response = NextResponse.redirect("https://www.tech-guide.jp")
    if (token) {
        response.cookies.set({
            name: 'accessToken',
            domain: '.tech-guide.jp',
            value: token,
            httpOnly: true,
            sameSite: 'lax',
            path: '/',
            secure: true,
        })
    }
    return response
}

export async function GET() {
    const cookieStore = cookies()
    const token = cookieStore.get("accessToken")
    return NextResponse.json(token, { status: 200 })
}
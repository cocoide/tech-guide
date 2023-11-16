import { cookies } from 'next/headers';
import { NextRequest, NextResponse } from 'next/server';


export async function GET(request: NextRequest) {
    const sessionId = request.nextUrl.searchParams.get("sessionId")
    if (!sessionId) {
        return NextResponse.json({ status: 403 })
    }
    cookies().set({
        name: 'techguide.sessionId',
        domain: '.tech-guide.jp',
        value: sessionId,
        httpOnly: true,
        sameSite: 'lax',
        maxAge: 60 * 60 * 24,// 1日
        path: '/',
        secure: true,
    })
    return NextResponse.redirect("https://www.tech-guide.jp/onboarding")
}
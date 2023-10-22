import { NextResponse } from 'next/server';
import {NextRequest} from "next/server";

export async function GET(request: NextRequest) {
    const token=request.nextUrl.searchParams.get("token")

    const response = NextResponse.redirect("https://www.tech-guide.jp/")
    if (token){
   response.cookies.set({
        name: 'accessToken',
        domain: '.tech-guide.jp',
        value: token,
        httpOnly: true,
        sameSite: 'none',
        path: '/',
    })
    }
    return response
}
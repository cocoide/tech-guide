import { cookies } from "next/headers";
import { NextResponse } from 'next/server';
import {NextRequest} from "next/server.js";
export async function GET(request: NextRequest) {
    const token=request.nextUrl.searchParams.get("token")
    if (token){
    cookies().set({
        name: 'accessToken',
        domain: '.tech-guide.jp',
        value: token,
        httpOnly: true,
        sameSite: 'none',
        path: '/',
    })
    }
    return NextResponse.redirect("www.tech-guide.jp/")
}
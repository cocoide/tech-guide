import { ResponseCookie } from 'next/dist/compiled/@edge-runtime/cookies';
import { cookies } from 'next/headers';
import { NextRequest, NextResponse } from 'next/server';
import { z } from 'zod';
export const corsHeaders = {
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

export async function POST(request: NextRequest) {
    const requestSchema = z.object({
        access_token: z.string(),
        // refresh_token: z.string().nullable(),
        display_name: z.string().nullable(),
        avatar_url: z.string().nullable(),
    });
    const body = await request.json();
    const validatedData = requestSchema.parse(body);
    const cookieStore = cookies()
    cookieStore.set("accessToken", validatedData.access_token, cookieOptions)
    return NextResponse.json("success", { status: 200, headers: corsHeaders })
}

export async function GET() {
    const cookieStore = cookies()
    const token = cookieStore.get("accessToken")
    return NextResponse.json(token, { status: 200 })
}
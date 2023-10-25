import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';
import { serverAuthFunc } from '../../../_server_actions/auth';

export async function GET() {
    const token = await serverAuthFunc.GetAccessToken()
    if (!token) {
        return NextResponse.json({ status: 403 })
    }
    return NextResponse.json(token, { status: 200 })
}

export async function DELETE(){
    const cookieStore = cookies()
    cookieStore.delete("accessToken").delete("refreshToken")
    return NextResponse.json("Cookies deleted", { status: 200 })
}


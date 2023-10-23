import { VerifyJwt } from '@/utils/jwt';
import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';

export async function GET() {
    const cookieStore = cookies()
    const tokenCookie = cookieStore.get("accessToken")
    if (!tokenCookie) {
        return NextResponse.json({ status: 403 })
    }
    var accessToken: string= tokenCookie.value
    const response = await VerifyJwt(accessToken)
    if (response!==null){
        if (response.accountID){

        }
        if (response.updatedToken){
            accessToken=response.updatedToken
        }
    }
    return NextResponse.json(accessToken, { status: 200 })
}

export async function DELETE(){
    const cookieStore = cookies()
    cookieStore.delete("accessToken").delete("refreshToken")
    return NextResponse.json("Cookies deleted", { status: 200 })
}


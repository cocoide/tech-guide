import jwt, {JwtPayload, VerifyOptions} from 'jsonwebtoken';
import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';
import { api } from "@/app/_functions/API";

export async function GET() {
    const cookieStore = cookies()
    const tokenCookie = cookieStore.get("accessToken")
    if (!tokenCookie) {
        return NextResponse.json({ status: 403 })
    }
    var accessToken: string= tokenCookie.value
    const response =await verifyToken(accessToken)
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

type VerifyResponse ={
    updatedToken?: string
    accountID?: number
}

async function verifyToken(token: string): Promise<VerifyResponse | null> {
    var response: VerifyResponse={}
    const option: VerifyOptions = {
        algorithms: ['HS256'],
    }
    const decoded = jwt.decode(token, option)
    if (decoded === null || typeof decoded === 'string' || !decoded.exp) {
        return null
    }
    if (Date.now() < decoded.exp * 1000) {
        const {data: accessToken}=await api.put<string>("/oauth/refresh",undefined)
        const cookieStore=cookies()
        if(!accessToken){
            return null
        }
        decoded["account_id"]=response.accountID
        cookieStore.set("accessToken",accessToken)
        response.updatedToken=accessToken
    }
    return response
}
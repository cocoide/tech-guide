import jwt, { VerifyOptions } from 'jsonwebtoken';
import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';

export async function GET() {
    const cookieStore = cookies()
    const token = cookieStore.get("accessToken")
    if (!token) {
        return NextResponse.json({ status: 403 })
    }
    // const claims = verifyToken(token.value)
    // if (claims === null) {
    //     return NextResponse.json({ status: 403 })
    // }
    return NextResponse.json(token?.value, { status: 200 })
}

type CustomClaims = {
    account_id: number
}

function verifyToken(token: string): CustomClaims | null {
    const option: VerifyOptions = {
        algorithms: ['HS256'],
    }
    const decoded = jwt.decode(token, option)
    if (decoded === null || typeof decoded === 'string' || !decoded.exp) {
        return null
    }
    if (Date.now() < decoded.exp * 1000) {
        // Add refresh token later
        return null
    }
    const account_id = decoded["account_id"] as number
    if (!account_id) {
        return null
    }
    const claims: CustomClaims = { account_id: account_id }
    return claims
}
import { api } from '@/app/_functions/API';
import { AccountSession } from '@/types/model';
import { decodeJwt } from '@/utils/jwt';
import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';

export async function GET() {
    try {
        var accessToken = cookies().get("accessToken")?.value
        if (!accessToken) {
            throw new Error(`Error getting token`)
        }
        const claims = await decodeJwt(accessToken)
        if (Date.now() < claims.exp * 1000) {
            console.log(claims.exp)
            console.log(Date.now())
            const refreshToken = cookies().get("refreshToken")?.value
            if (!refreshToken) {
                throw new Error(`Error getting refreshToken in cookie`)
            }
            const params = { "token": refreshToken }
            const { data: accessToken, error } = await api.pos<string>("/oauth/refresh", undefined, undefined, params)
            if (error || !accessToken) {
                throw new Error(`Failed to refresh token: ${error}`)
            }
            // CookieのAccessTokenも更新
            cookies().set({
                name: 'accessToken',
                domain: '.tech-guide.jp',
                value: accessToken,
                httpOnly: true,
                sameSite: 'lax',
                maxAge: 60 * 60 * 24 * 30,// 7日
                path: '/',
                secure: true,
            })
        }
        const { data: session, error } = await api.get<AccountSession>("/account/session", "no-store", accessToken)
        if (!session || error) {
            throw new Error(`Error getting session: ${error}`)
        }
        return NextResponse.json(session, { status: 200 })
    } catch (error) {
        console.error(error)
        return NextResponse.json(error, { status: 403 })
    }
}
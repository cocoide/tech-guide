import { api } from '@/app/_functions/API'
import { AccountSession } from '@/types/model'
import { VerifyJwt } from '@/utils/jwt'
import { cookies } from 'next/headers'
import { NextResponse } from 'next/server'

export async function GET() {
    const cookieStore = cookies()
    var accessToken = cookieStore.get("accessToken")?.value
    if (!accessToken) {
        return NextResponse.json({ status: 403 })
    }
    const response = await VerifyJwt(accessToken)
    if (response !== null) {
        if (response.accountID) {

        }
        if (response.updatedToken) {
            accessToken = response.updatedToken
        }
    }
    const { data: session, ok } = await api.get<AccountSession>("/account/session", "no-store", accessToken)
    if (!ok) {
        return NextResponse.json({ status: 403 })
    }
    return NextResponse.json(session, { status: 200 })
}
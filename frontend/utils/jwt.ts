import { api } from '@/app/_functions/API'
import jwt, { VerifyOptions } from 'jsonwebtoken'
import { cookies } from 'next/headers'

type VerifyResponse = {
    updatedToken?: string
    accountID?: number
}

export async function VerifyJwt(token: string): Promise<VerifyResponse | null> {
    "use server"

    var response: VerifyResponse = {}
    const option: VerifyOptions = {
        algorithms: ['HS256'],
    }
    const decoded = jwt.decode(token, option)
    if (decoded === null || typeof decoded === 'string' || !decoded.exp) {
        return null
    }
    if (Date.now() < decoded.exp * 1000) {
        const refresh_token = cookies().get("refresh_token")?.value
        if (!refresh_token || refresh_token.length==0){
            return null
        }
        const params ={"token":refresh_token }
        const { data: accessToken } = await api.put<string>("/oauth/refresh", undefined, undefined, params)
        const cookieStore = cookies()
        if (!accessToken) {
            return null
        }
        decoded["account_id"] = response.accountID
        cookieStore.set("accessToken", accessToken)
        response.updatedToken = accessToken
    }
    return response
}
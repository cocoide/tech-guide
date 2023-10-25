import jwt, { VerifyOptions } from 'jsonwebtoken'
import { cookies } from 'next/headers'
import { api } from '../_functions/API'

export const serverAuthFunc = {
    async GetAccessToken() {
        "use server"

        const cookieStore = cookies()
        var response = ""
        var accessToken = cookieStore.get("accessToken")?.value
        if (!accessToken) {
            return
        }
        const option: VerifyOptions = {
            algorithms: ['HS256'],
        }
        const decoded = jwt.decode(accessToken, option)
        if (decoded === null || typeof decoded === 'string' || !decoded.exp) {
            return
        }
        // const account_id = decoded["account_id"]
        // if (Date.now() < decoded.exp * 1000) {
        //     const updateAccessToken = await refreshToken()
        //     if (!updateAccessToken) {
        //         return
        //     }
        //     response = updateAccessToken
        //     cookieStore.set("accessToken", updateAccessToken)
        // }
        return response
    },
}

async function refreshToken(): Promise<string | undefined> {
    "use server"

    const cookieStore = cookies()
    const refresh_token = cookieStore.get("refresh_token")?.value
    if (!refresh_token) {
        return
    }
    const params = { "token": refresh_token }
    const { data: accessToken } = await api.pos<string>("/oauth/refresh", undefined, undefined, params)
    if (!accessToken) {
        return
    }
    return accessToken
}
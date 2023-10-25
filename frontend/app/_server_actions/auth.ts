import jwt, { VerifyOptions } from 'jsonwebtoken'
import { cookies } from 'next/headers'
import { api } from '../_functions/API'

export const serverAuthFunc = {
    async GetAccessToken() {
        "use server"

        var response = ""
        var accessToken = cookies().get("accessToken")?.value
        if (!accessToken) {
            return
        }
        response = accessToken
        const option: VerifyOptions = {
            algorithms: ['HS256'],
        }
        const decoded = jwt.decode(accessToken, option)
        if (decoded === null || typeof decoded === 'string' || !decoded.exp) {
            return
        }
        if (Date.now() < decoded["exp"] * 1000) {
            const refresh_token = cookies().get("refresh_token")?.value
            if (!refresh_token) {
                return
            }
            const updateAccessToken = await refreshToken(refresh_token)
            if (!updateAccessToken) {
                return
            }
            cookies().set("accessToken", updateAccessToken)
            response = updateAccessToken
        }
        return response
    },
}

async function refreshToken(refreshToken: string): Promise<string | undefined> {
    const params = { "token": refreshToken }
    const { data: accessToken } = await api.pos<string>("/oauth/refresh", undefined, undefined, params)
    if (!accessToken) {
        return
    }
    return accessToken
}
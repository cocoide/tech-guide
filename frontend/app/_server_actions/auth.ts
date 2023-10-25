import jwt, { VerifyOptions } from 'jsonwebtoken'
import { cookies } from 'next/headers'
import { api } from '../_functions/API'

export const serverAuthFunc = {
    async GetAccessToken() {
        "use server"

        try {
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
            if (decoded === null || typeof decoded === 'string' || !decoded["exp"]) {
                throw new Error("Failed to decode accessToken")
        }
        if (Date.now() < decoded["exp"] * 1000) {
            const updateAccessToken = await refreshToken()
            if (!updateAccessToken) {
                throw new Error("Failed to refresh token")
            }
            cookies().set("accessToken", updateAccessToken)
            response = updateAccessToken
        }
        } catch (error) {
            console.error("Error in GetAccessToken:", error);
            return
        }
        return response
    },
}

async function refreshToken(): Promise<string | undefined> {
    "use server"

    const refreshToken = cookies().get("refreshToken")?.value
    if (!refreshToken) {
        throw new Error("Error getting refresToken in cookies")
    }
    const params = { "token": refreshToken }
    const { data: accessToken, error } = await api.pos<string>("/oauth/refresh", undefined, undefined, params)
    if (error || !accessToken) {
        throw new Error(`Failed to refresh token: ${error}`)
    }
    return accessToken
}
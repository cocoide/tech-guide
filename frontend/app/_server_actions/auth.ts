import { AccountSession } from '@/types/model'
import { cookies } from 'next/headers'
import { api } from '../_functions/API'

export const serverAuthFunc = {
    async GetAccountSession() {
        "use server"

        let response: AccountSession = { account_id: 0, avatar_url: "", display_name: "" }
        try {
            const token = await this.GetAccessToken()
            const { data: session, error } = await api.get<AccountSession>("/account/session", "no-store", token)
            if (!session || error) {
                throw new Error(`Failed to get session for: ${error}`)
            }
            response = session
        } catch (error) {
            console.error(error)
            return
        }
        return response
    },
    async GetAccessToken() {
        "use server"

        try {
            const cookieStore = cookies()
            var accessToken = cookieStore.get("accessToken")?.value
        if (!accessToken) {
            throw new Error("Error getting accessToken in cookies")
        }
            // const option: VerifyOptions = {
            //     algorithms: ['HS256'],
            // }
            // const decoded = jwt.decode(accessToken, option)
            //     if (decoded === null || typeof decoded === 'string' || !decoded["exp"]) {
            //         throw new Error("Failed to decode accessToken")
            // }
            // if (Date.now() < decoded["exp"] * 1000) {
            //     const newAccessToken = await this.refreshToken()
            //     if (!newAccessToken) {
            //         throw new Error(`Failed to refresh accessToken`)
            //     }
            //     cookies().set({
            //         name: 'accessToken',
            //         domain: '.tech-guide.jp',
            //         value: newAccessToken,
            //         httpOnly: true,
            //         sameSite: 'lax',
            //         maxAge: 60 * 60 * 24 * 7,// 7日
            //         path: '/',
            //         secure: true,
            //     })
            //     return newAccessToken
            // }
            return accessToken
        } catch (error) {
            console.log(`Error in GetAccessToken: ${error}`);
            throw error;
        }
    },
    async refreshToken() {
        "use server"

    const refreshToken = cookies().get("refreshToken")?.value
    try {
    if (!refreshToken) {
        throw new Error("Error getting refreshToken in cookies")
    }
    const params = { "token": refreshToken }
    const { data: accessToken, error } = await api.pos<string>("/oauth/refresh", undefined, undefined, params)
    if (error || !accessToken) {
        throw new Error(`Failed to refresh token: ${error}`)
    }
        return accessToken
    } catch (error) {
        console.log(`Failed to refresh token: ${error}`)
        throw error
    }
    },
}


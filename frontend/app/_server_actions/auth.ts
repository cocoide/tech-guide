import { AccountSession } from '@/types/model'
import { decodeJwt, refreshAccessToken } from '@/utils/jwt'
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
            var accessToken = cookies().get("accessToken")?.value
            if (!accessToken) {
                throw new Error(`Error getting token`)
            }
            const claims = await decodeJwt(accessToken)
            if (Date.now() < claims.exp * 1000) {
                const refreshToken = cookies().get("refreshToken")?.value
                if (!refreshToken) {
                    throw new Error(`Error getting refreshToken in cookie`)
                }
                accessToken = await refreshAccessToken(refreshToken)
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
            return accessToken
        } catch (error) {
            console.error(`Failed to get accessToken: ${error}`)
            return
        }
    },
}


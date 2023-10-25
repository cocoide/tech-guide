import { api, apiRoute } from '@/app/_functions/API'
import { AccountSession } from '@/types/model'

export const authAPI ={
    async GetAccessToken() {
        const { data: token, error } = await apiRoute.get<string>("/oauth/token")
        if (error) {
            console.log(error)
        }
        return token
    },
    async GetAccountSession(token: string) {
        const { data: session, error } = await api.get<AccountSession>("/session", "no-store", token)
        if (error) {
            console.log(error)
        }
        return session
    }
}
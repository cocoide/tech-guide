import { apiRoute } from '@/app/_functions/API'
import { AccountSession } from '@/types/model'

export const authAPI ={
    async GetAccessToken() {
        const { data: token, error } = await apiRoute.get<string>("/oauth/token")
        if (error) {
            console.log(error)
        }
        return token
    },
    async GetAccountSession() {
        const { data: session, error } = await apiRoute.get<AccountSession>("/session", "no-store")
        if (error) {
            console.log(error)
        }
        return session
    }
}
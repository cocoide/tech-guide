import { api, apiRoute } from '@/app/_functions/API'

export const authAPI ={
    async GetAccessToken() {
        const { data: token, error, status } = await apiRoute.get<string>("/oauth/token")
        console.log(token)
        console.log(error)
        console.log(status)
        return token
    },
    async RefreshToken(refresh_token: string){
        return await api.pos<string>("/refresh",refresh_token)
    },
}
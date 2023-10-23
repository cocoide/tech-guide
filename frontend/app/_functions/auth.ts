import { api, apiRoute } from '@/app/_functions/API'

export type SignupRequest={
    email:string,
    name: string,
    image: string,
}
export type LoginRequest={
    email:string,
}
type LoginResponse={
    token:string,
    token_expires: number
    uid: number,
    image: string,
    name: string,
}

type RefreshTokenResponse={
    token:string,
    token_expires: number
}
export const authAPI ={
    async GetAccessToken() {
        const { data: token, error, status } = await apiRoute.get<string>("/oauth/token")
        console.log(token)
        console.log(error)
        console.log(status)
        return token
    },
    async SignUp(req: SignupRequest){
        return await api.pos<LoginResponse>("/signup",req)
    },
    async IsEmailUsed(email: string): Promise<boolean|undefined>{
        const { data: isRegisterd, ok } = await api.get<boolean>(`/email?email=${email}`,"no-store")
        if (!ok) {
          return undefined
        }
        return isRegisterd
    },
    async Login(req: LoginRequest){
      return await api.pos<LoginResponse>("/login",req)
    },
    async RefreshToken(refresh_token: string){
        return await api.pos<RefreshTokenResponse>("/refresh",refresh_token)
    },
}
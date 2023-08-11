import { api } from '@/app/_functions/API'
import { authOptions } from '@/libs/next-auth'
import { getServerSession } from 'next-auth'

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
    async GetAuthSession(){
        const session = await getServerSession(authOptions)
        return {
            token: session?.token,
            token_expires: session?.token_expires,
            user: {
                uid: session?.user?.uid,
                name: session?.user?.name,
                image: session?.user?.image,
            },
        }
    }
}
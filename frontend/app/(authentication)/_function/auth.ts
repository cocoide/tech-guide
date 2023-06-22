import { api } from '@/app/_functions/API'
import { User } from 'next-auth'

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
    uid: number,
    image: string,
    name: string,
}
export const authAPI ={
    async SignUp(req: SignupRequest){
        return await api.pos<User>("/signup",req)
    },
    async IsEmailUsed(email: string): Promise<boolean|undefined>{
        const { data: isRegisterd, ok } = await api.get<boolean>(`/email?email=${email}`,"no-store")
        if (!ok) {
          return undefined
        }
        return isRegisterd
    },
    async Login(req: LoginRequest): Promise<LoginResponse|null>{
        const { data } =await api.pos<LoginResponse>("/login",req)
        if(!data){
            return null
        }
        return data
    }
}
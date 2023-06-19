import { api } from '@/app/_functions/API'
import { User } from 'next-auth'

export type LoginRequest={
    email:string,
    password:string
}
export const authAPI ={
    async Login(body: LoginRequest){
        return await api.pos<User>("/login",body)
    },
}
import { useSession } from 'next-auth/react'

export const useAuth=()=>{
    const {data: session, status} = useSession()
    return {
        token : session?.token,
        expires: session?.token_expires,
        user:{
            uid: session?.user.uid,
            name: session?.user.name,
            image: session?.user.image,
        },
        status: status
    }
}
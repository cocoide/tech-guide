import { authAPI } from '@/app/_functions/auth'
import { useQuery } from '@tanstack/react-query'

export const useAuth=()=>{
    const { data: token, isLoading } = useQuery({
        queryFn: async () => await authAPI.GetAccessToken(),
        queryKey: ["access_token"],
    })
    var status: string
    if (token !== "" || token !== undefined) {
        status = "authenticated"
    } else if (isLoading) {
        status = "loading"
    } else {
        status = "unauthenticated"
    }
    return {
        token: token,
        status: status,
        user:{
            uid: 6,
            name: "test",
            image: "",
        },
    }
}
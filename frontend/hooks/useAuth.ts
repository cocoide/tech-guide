import { authAPI } from '@/app/_functions/auth';
import { useQuery } from '@tanstack/react-query';

type AuthStatus = 'authenticated' | 'unauthenticated' | 'loading'
type AuthData = {
    token?: string;
    status: AuthStatus;
    user: UserData;
}
type UserData = {
    uid: number;
    image?: string;
    name: string;
}

export const useAuth = (): AuthData => {
    const { data: token, isLoading } = useQuery({
        queryFn: async () => await authAPI.GetAccessToken(),
        queryKey: ["access_token"],
    })
    var status: AuthStatus
    if (token !== "" || token !== undefined) {
        console.log(token)
        status = "authenticated"
    } else if (isLoading) {
        status = "loading"
    } else {
        status = "unauthenticated"
    }
    return {
        token: token,
        status: status,
        user: {
            uid: 6,
            name: "test",
            image: "",
        },
    }
}
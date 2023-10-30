import { authAPI } from '@/app/_functions/auth';
import { useQuery } from '@tanstack/react-query';

type AuthStatus = 'authenticated' | 'unauthenticated' | 'loading'
type AuthData = {
    token?: string;
    status: AuthStatus;
}

export const useAuth = (): AuthData => {
    const { data: token, isLoading } = useQuery({
        queryFn: async () => await authAPI.GetAccessToken(),
        queryKey: ["access_token"],
    })
    var status: AuthStatus
    if ( token != undefined && token.length>0 ) {
        status = "authenticated"
    } else if (isLoading) {
        status = "loading"
    } else {
        status = "unauthenticated"
    }
    return {
        token: token,
        status: status,
    }
}

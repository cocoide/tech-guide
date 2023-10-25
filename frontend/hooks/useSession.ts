import { authAPI } from '@/app/_functions/auth'
import { AccountSession } from '@/types/model'
import { useQuery } from '@tanstack/react-query'
import { useAuth } from './useAuth'

export const useSession = (): AccountSession => {
	var response: AccountSession = {
		account_id: 0,
		display_name: '',
		avatar_url: '',
		features: []
	}
	const { token } = useAuth()
	if (token) {
		// eslint-disable-next-line react-hooks/rules-of-hooks
		const { data: session } = useQuery({
			queryFn: async () => await authAPI.GetAccountSession(token),
			queryKey: ["account_session"],
		})
		if (session) {
			response = session
		}
	}
	return response
}

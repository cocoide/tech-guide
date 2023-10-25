import { authAPI } from '@/app/_functions/auth'
import { AccountSession } from '@/types/model'
import { useQuery } from '@tanstack/react-query'

export const useSession = (): AccountSession => {
	var response: AccountSession = {
		account_id: 0,
		display_name: '',
		avatar_url: '',
		features: []
	}
		const { data: session } = useQuery({
			queryFn: async () => await authAPI.GetAccountSession(),
			queryKey: ["account_session"],
		})
		if (session) {
			response = session
		}
	return response
}

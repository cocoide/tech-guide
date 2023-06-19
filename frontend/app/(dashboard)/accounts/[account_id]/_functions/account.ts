import { api } from '@/app/_functions/API';
import { Account } from '@/app/_models';

export type GetPublicProfile =Pick<Account,'avatar_url'|'display_name'>

export const accountAPI={
    async getProfile(accountId: number){
        const res= await api.get<GetPublicProfile>(`/account/profile/${accountId}`,'no-store')
        return res
    }
}
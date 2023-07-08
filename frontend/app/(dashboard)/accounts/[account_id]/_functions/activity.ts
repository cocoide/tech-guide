import { api } from '@/app/_functions/API'
import { Contribution } from '@/types/model'

export const activityAPI={
    async GetContributions(account_id: string){
        return await api.get<Contribution[]>(`/contribution/${account_id}`,"no-store")
    }
}
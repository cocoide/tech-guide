import { api } from '@/app/_functions/API'
import { Collection } from '@/app/_models'

export const collectionAPI={
    async getCollections(accountId: number){
        const res= await api.get<Collection[]>(`/account/collection/${accountId}`,"no-store")
        return res
    },
    async getCollectionForBookmark(token?: string) {
        return  await api.get<Collection[]>(`/account/collection`, 'no-store', token)
    }
}
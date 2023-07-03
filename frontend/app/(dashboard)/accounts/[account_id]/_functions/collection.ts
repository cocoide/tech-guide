import { NewCollectionRequest } from '@/app/_components/layouts/CollectionDialog'
import { api } from '@/app/_functions/API'
import { Collection } from '@/types/model'

export const collectionAPI={
    async getCollections(accountId: number){
        const res= await api.get<Collection[]>(`/account/collection/${accountId}`,"no-store")
        return res
    },
    async getCollectionForBookmark(token?: string) {
        return  await api.get<Collection[]>(`/account/collection`, 'no-store', token)
    },
    async makeCollectionWithBookmark(collection: NewCollectionRequest, articleId: number, token?: string) {
        return  await api.pos(`/account/collection`,collection, token, {"articleId": articleId} )
    }
}
import { api } from '@/app/_functions/API'
import { Collection } from '@/types/model'

export const collectionAPI={
    async GetCollectionData(collectionId: string){
        return await api.get<Collection>(`/collection/${collectionId}`,"no-store")
    }
}
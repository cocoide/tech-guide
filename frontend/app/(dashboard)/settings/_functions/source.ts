import { api } from '@/app/_functions/API'
import { Source } from '@/types/model'

export const sourceAPI ={
    async GetFollowingSources(token?: string){
        return await api.get<Source[]>("/account/source/follow","no-store",token)
    },
    async GetAllSources(){
        return await api.get<Source[]>('/source')
    },
    async DoFollowSource(source_id: number, token?: string){
        return await api.put(`/account/source/follow/${source_id}`,undefined,token)
    },
    async UnFollowSource(source_id: number, token?: string){
        return await api.del(`/account/source/follow/${source_id}`,token)
    },
}
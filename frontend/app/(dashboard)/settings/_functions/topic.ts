import { api } from '@/app/_functions/API'
import { Category, Topic } from '@/types/model'

export const topicAPI ={
    async GetFollowingTopics(token?: string){
        return await api.get<Topic[]>("/account/topic/follow","no-store",token)
    },
    async GetAllTopics(){
        return await api.get<Topic[]>('/topic')
    },
    async GetAllCategories() {
        return await api.get<Category[]>('/category')
    },
    async DoFollowTopic(topic_id: number, token?: string){
        return await api.put(`/account/topic/follow/${topic_id}`,undefined,token)
    },
    async UnFollowTopic(topic_id: number, token?: string){
        return await api.del(`/account/topic/follow/${topic_id}`,token)
    },
}
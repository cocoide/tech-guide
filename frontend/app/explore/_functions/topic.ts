import { api } from '@/app/_functions/API'

type Topic={
    id: number
    name: string
}

export const topicAPI ={
    async GetPopularTopics() {
        return await api.get<Topic[]>('/topic/popular')
    }
}
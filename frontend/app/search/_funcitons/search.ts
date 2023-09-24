import { api } from '@/app/_functions/API'
import { Article } from '@/types/model'

export const searchAPI = {
    async findArticlesByTitle(title: string) {
        return api.get<Article[]>(`/search/title?title=${title}`)
    }
}
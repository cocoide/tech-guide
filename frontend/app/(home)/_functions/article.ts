import { api } from '@/app/_functions/API';
import { Article } from '@/app/_models';
import { MakeArticle } from '../type';

export const articleAPI={
    async getLatestArticles() {
        const res =  await api.get<Article[]>("/article",60)
        return res.data
    },
    async createArticle(article:MakeArticle){
       return await api.pos("/article",article)
    }
}
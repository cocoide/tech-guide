import { api } from '@/app/_functions/API';
import { Article } from '@/types/model';
import { MakeArticle } from '../type';

export const articleAPI={
    async getLatestArticles() {
        const res =  await api.get<Article[]>("/article",60)
        return res.data
    },
    async createArticle(article:MakeArticle){
       return await api.pos("/article",article)
    },
    async GetRelatedArticles(article_id: string){
        return await api.get<Article[]>(`/article/related/${article_id}`,60*60)
    },
    async GetArticleDetail(article_id: string){
        return await api.get<Article>(`/article/${article_id}`)
    },
    async GetRecommendArticles(token: string){
        return await api.get<Article[]>('/account/article/recommend', 'no-store', token)
    }
}
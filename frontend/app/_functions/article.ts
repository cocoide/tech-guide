import { api } from '@/app/_functions/API';
import { Article, Comment } from '@/types/model';
export type Header = {
    content: string
    subHeader?: Header
}
type MakeArticle=Pick<Article,"original_url">

export const articleAPI = {
    async getLatestArticles() {
        const res = await api.get<Article[]>("/article", 60)
        return res.data
    },
    async createArticle(article: MakeArticle) {
        return await api.pos("/article", article)
    },
    async GetRelatedArticles(article_id: string) {
        return await api.get<Article[]>(`/article/related/${article_id}`, "no-store")
    },
    async GetArticleDetail(article_id: string) {
        return await api.get<Article>(`/article/${article_id}`)
    },
    async GetRecommendArticles(token: string) {
        return await api.get<Article[]>('/account/article/recommend', 'no-store', token)
    },
    async ReadArticle(article_id: number, token: string) {
        return await api.put('/account/article/read', undefined, token, { "article_id": article_id })
    },
    async GetReadArticles(token?: string) {
        return await api.get<Article[]>('/account/article/read', "no-store", token)
    },
    async GetOverview(url?: string){
        return await api.get<string>(`/overview?url=${url}`,"reload")
    },
    async GetHeaders(url?: string) {
        return await api.get<Header[]>(`/scraper/header`, "reload")
    },
    async GetArticlesByPagination(page: number){
        const { data } = await api.get<Article[]>(`/article?page=${page} `, "no-store")
        return data
    },
    async GetFeedsByPagination(page: number, token?: string) {
        const { data, error } = await api.get<Article[]>(`/account/feeds?page=${page}`, "no-store", token)
        console.log(error)
        return data
    },
    async GetTopCommentForArticle(articleID: number){
        return await api.get<Comment>(`/comment/${articleID}`, 24 * 60 * 60)
    },
    async GetCommentsForArticle(articleID: number){
        return await api.get<Comment[]>(`/comment/${articleID}`, "no-store")
    },
    async GetrArticlesBySourceID(sourceID: number, page: number){
        return await api.get<Article[]>(`/article/source/${sourceID}?page=${page}`, "no-store")
    },
    async GetrArticlesByTopicID(topicID: number, page: number){
        return await api.get<Article[]>(`/article/topic/${topicID}?page=${page}`, "no-store")
    },
}
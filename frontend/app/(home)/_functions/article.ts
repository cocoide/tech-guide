import { api } from '@/app/_functions/API';
import { Article } from '@/types/model';

export async function fetchTrendArticles(page = 1): Promise<Article[]> {
    "use server"

    let result: Article[] = [];
    const { data } = await api.get<Article[]>(`/article/trend?page=${page}`)
    if (data) {
        result = data
    }
    return result
}

export async function fetchDiscussArticles(page = 1): Promise<Article[]> {
    "use server"

    let result: Article[] = [];
    const { data } = await api.get<Article[]>(`/article/discuss?page=${page}`)
    if (data) {
        result = data
    }
    return result
}

export async function fetchLatestArticles(page = 1): Promise<Article[]> {
    "use server"

    let result: Article[] = [];
    const { data } = await api.get<Article[]>(`/article?page=${page}`)
    if (data) {
        result = data
    }
    return result
}

export async function getArticlesByTopicID(page = 1, topicID?: number): Promise<Article[]> {
    "use server"

    if (!topicID) {
        return []
    }

    let result: Article[] = [];
    const { data } = await api.get<Article[]>(`/article/topic/${topicID}?page=${page}`)
    if (data) {
        result = data
    }
    return result
}

export async function getArticlesBySourceID(page = 1, sourceID?: number): Promise<Article[]> {
    "use server"

    if (!sourceID) {
        return []
    }

    let result: Article[] = [];
    const { data } = await api.get<Article[]>(`/article/source/${sourceID}?page=${page}`)
    if (data) {
        result = data
    }
    return result
}

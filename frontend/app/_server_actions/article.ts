import { Article } from '@/types/model';
import { articleAPI } from '../_functions/article';

export async function FetchTrendArticles(): Promise<Article[]> {
    "use server"

    const { data } = await articleAPI.GetTrendArticles()
    if (!data) {
        return []
    }
    return data
}
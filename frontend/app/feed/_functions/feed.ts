'use server'

import { api } from '@/app/_functions/API';
import { serverAuthFunc } from '@/app/_server_actions/auth';
import { Article } from '@/types/model';

export async function FetchFeed(page = 1): Promise<Article[]> {
    "use server"
    var articles: Article[] = [];

    const token = await serverAuthFunc.GetAccessToken()
    if (!token) {
        console.log("Unauthorized")
        return articles
    }
    const { data, error } = await api.get<Article[]>(`/account/feeds?page=${page}`, "no-store", token);
    console.log(error)
    if (data) {
        articles = data
    }
    return articles
}
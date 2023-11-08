'use server'

import { api } from '@/app/_functions/API';
import { Article } from '@/types/model';
import { cookies } from "next/headers";

export async function FetchFeed(page = 1): Promise<Article[]> {
    "use server"

    var articles: Article[] = [];

    var accessToken = cookies().get("accessToken")?.value
    if (!accessToken) {
        throw new Error(`Error getting token`)
    }
    const { data, error } = await api.get<Article[]>(`/account/feeds?page=${page}`, "no-store", accessToken);
    console.log(error)
    if (data) {
        articles = data
    }
    return articles
}
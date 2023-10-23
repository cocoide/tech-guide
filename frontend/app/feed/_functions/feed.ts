'use server'

import { api } from '@/app/_functions/API';
import { Article } from '@/types/model';
import { VerifyJwt } from '@/utils/jwt';
import { cookies } from 'next/headers';

export async function FetchFeed(page = 1): Promise<Article[]> {
    "use server"

    var token = cookies().get("accessToken")?.value
    if (token) {
        const resp = await VerifyJwt(token)
        if (resp?.updatedToken) {
            token = resp?.updatedToken
        }
    }
    var articles: Article[] = [];
    const { data, error } = await api.get<Article[]>(`/account/feeds?page=${page}`, "no-store", token);
    console.log(error)
    if (data) {
        articles = data
    }
    return articles
}
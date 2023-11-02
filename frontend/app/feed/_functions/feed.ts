'use server'

import { api } from '@/app/_functions/API';
import { serverAuthFunc } from '@/app/_server_actions/auth';
import { Article } from '@/types/model';
import {cookies} from "next/headers";
import {decodeJwt, refreshAccessToken} from "@/utils/jwt";

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
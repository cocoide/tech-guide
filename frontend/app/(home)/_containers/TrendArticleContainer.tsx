import ArticleItems from '@/app/feed/_components/ArticleItems'
import { fetchTrendArticles } from '../_functions/article'

export default async function TrendArticleContainer() {
    const articles = await fetchTrendArticles()
    return (
        <ArticleItems articles={articles} fetchFunc={fetchTrendArticles} />
    )
}
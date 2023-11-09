import ArticleItems from '@/app/feed/_components/ArticleItems'
import { fetchLatestArticles } from '../_functions/article'

export default async function LatestArticleContainer() {
    const articles = await fetchLatestArticles()
    return (
        <ArticleItems articles={articles} fetchFunc={fetchLatestArticles} />
    )
}
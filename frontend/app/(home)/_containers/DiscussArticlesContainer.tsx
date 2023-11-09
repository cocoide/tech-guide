import ArticleItems from '@/app/feed/_components/ArticleItems'
import { fetchDiscussArticles } from '../_functions/article'

export default async function DiscussArticleContainer() {
    const articles = await fetchDiscussArticles()
    return (
        <ArticleItems articles={articles} fetchFunc={fetchDiscussArticles} />
    )
}
import ArticleItems from '@/app/feed/_components/ArticleItems'
import { getArticlesBySourceID } from '../_functions/article'

export default async function ArticlesBySourceIDContainer({ sourceID }: { sourceID: number }) {
    const articles = await getArticlesBySourceID(undefined, sourceID)
    return (
        <ArticleItems articles={articles} fetchFunc={getArticlesBySourceID} />
    )
}
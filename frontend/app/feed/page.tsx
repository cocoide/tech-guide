
import ArticleItems from './_components/ArticleItems'
import { FetchFeed } from './_functions/feed'

export default async function FeedPage() {
    const feeds = await FetchFeed()
    return (
        <ArticleItems articles={feeds} fetchFunc={FetchFeed} />
    )
}
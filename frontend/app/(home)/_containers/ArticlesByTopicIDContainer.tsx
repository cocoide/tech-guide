import ArticleItems from '@/app/feed/_components/ArticleItems';
import { getArticlesByTopicID } from '../_functions/article';

export default async function ArticlesByTopicIDContainer({ topicID }: { topicID: number }) {
    const articles = await getArticlesByTopicID(undefined, topicID);
    return (
        <ArticleItems articles={articles} fetchFunc={getArticlesByTopicID} />
    )
}
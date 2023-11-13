import ArticleCard from '@/app/(home)/_components/ArticleCard';
import { articleAPI } from '@/app/_functions/article';
import { Article } from '@/types/model';

interface Props {
    origin?: number
    exclude?: number
    article_id: string
}

export default async function RelatedArticles({ origin, article_id, exclude }: Props) {
    const { data: relatedArticles } = await articleAPI.GetRelatedArticles(article_id);
    let renderArticles: Article[] | undefined = []
    if (exclude) {
        renderArticles = relatedArticles?.filter(article => article.id !== exclude);
    } else {
        renderArticles = relatedArticles
    }
    return (
        <div className='grid sm:grid-cols-2 lg:grid-cols-3 gap-3 p-3'>
            {renderArticles?.map(article => {
                return <ArticleCard key={article.id} article={article} origin={origin} />
            })}
        </div>
    )
}
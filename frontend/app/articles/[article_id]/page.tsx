import { articleAPI } from '../../_functions/article';
import ArticleDetail from './components/ArticleDetail';
import RelatedArticles from './components/RelatedArticles';

interface Props extends ArticleParams {
    searchParams: { "exclude": string }
}
export default async function ArticlePage({ params, searchParams }: Props) {
    const exclude = Number(searchParams.exclude)
    const { article_id } = params;
    const { data: relatedArticles } = await articleAPI.GetRelatedArticles(article_id);
    const selectedRelatedArticles = relatedArticles?.filter(article => article.id !== exclude);
    const { data: articleDetail } = await articleAPI.GetArticleDetail(article_id);

    return (
        <div className="flex flex-col md:p-5 space-y-5">
            {articleDetail &&
                <ArticleDetail article={articleDetail} />
            }
            {selectedRelatedArticles && selectedRelatedArticles.length > 0 &&
                <RelatedArticles articles={selectedRelatedArticles} origin={String(articleDetail?.id)} />
            }
        </div>
    )
}
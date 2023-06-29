import { articleAPI } from '../../_functions/article';
import ArticleDetail from './components/ArticleDetail';
import RelatedArticles from './components/RelatedArticles';

export default async function ArticlePage({ params }: ArticleParams) {
    const { article_id } = params;
    const { data: relatedArticles } = await articleAPI.GetRelatedArticles(article_id);
    const { data: articleDetail } = await articleAPI.GetArticleDetail(article_id);
    return (
        <div className="flex flex-col md:p-5 space-y-5">
            {articleDetail &&
                <ArticleDetail article={articleDetail} />
            }
            {relatedArticles && relatedArticles.length > 0 &&
                <RelatedArticles articles={relatedArticles} origin={String(articleDetail?.id)} />
            }
        </div>
    )
}
import { RelatedArticlesLoader } from '@/app/_components/loaders/RelatedArticlesLoader';
import { NewspaperIcon } from '@heroicons/react/24/outline';
import { Suspense } from 'react';
import { articleAPI } from '../../_functions/article';
import ArticleDetail from './components/ArticleDetail';
import RelatedArticles from './components/RelatedArticles';

interface Props extends ArticleParams {
    searchParams: { "origin": string }
}
export default async function ArticlePage({ params, searchParams }: Props) {
    const exclude = Number(searchParams?.origin)
    const { article_id } = params;
    const { data: articleDetail } = await articleAPI.GetArticleDetail(article_id);

    return (
        <div className="flex flex-col md:p-5 space-y-5">
            {articleDetail &&
                <ArticleDetail article={articleDetail} />
            }
            <div className="text-slate-600 custom-badge ml-3">
                <NewspaperIcon className="h-7 w-7" />
                <> 関連記事</>
            </div>
            <Suspense fallback={<RelatedArticlesLoader />}>
                <RelatedArticles origin={articleDetail?.id} article_id={params.article_id} exclude={exclude} />
            </Suspense>
        </div>
    )
}
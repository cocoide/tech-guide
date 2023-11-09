import { articleAPI } from '@/app/_functions/article'
import ArticleCard from '../_components/ArticleCard'

export default async function DiscussArticleContainer() {
    const { data: articles } = await articleAPI.GetDiscussArticles()
    return (
        <div className="min-h-screen w-full grid sm:grid-cols-2 xl:grid-cols-3  gap-[20px] py-[20px] px-[30px] sm:p-[20px]">
            {articles?.map((article, index) => (
                <ArticleCard key={article.title + index} article={article} />
            )
            )}
        </div>
    )
}
import ArticleCard from './_components/ArticleCard';
import { articleAPI } from './_functions/article';

export default async function ArticlePage() {
    const articles = await articleAPI.getLatestArticles()
    return (
        <div className="min-h-screen w-full pt-3 px-3 lg:px-[10%] gap-3 grid lg:grid-cols-2">
            {articles?.map(article => (
                <ArticleCard key={article.title} article={article} />
            )
            )}
        </div>
    )
}
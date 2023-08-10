import { articleAPI } from '@/app/_functions/article'
import { authAPI } from '@/app/_functions/auth'
import ArticleCard from '@/app/trend/_components/ArticleCard'

export default async function HistoryPage() {
    const { token } = await authAPI.GetAuthSession()
    const { data: articles } = await articleAPI.GetReadArticles(token)
    return (
        <div className="flex flex-col space-y-3 p-5">
            {articles && articles?.map(article => (
                <ArticleCard key={article.id} article={article} />
            ))}
        </div>
    )
}
import ArticleCard from '@/app/(timeline)/_components/ArticleCard'
import { articleAPI } from '@/app/(timeline)/_functions/article'
import { authAPI } from '@/app/_functions/auth'

export default async function HistoryPage() {
    const { token } = await authAPI.GetAuthSession()
    const { data: articles } = await articleAPI.GetReadArticles(token)
    return (
        <div className="flex flex-col space-y-3">
            {articles && articles?.map(article => (
                <ArticleCard key={article.id} article={article} />
            ))}
        </div>
    )
}
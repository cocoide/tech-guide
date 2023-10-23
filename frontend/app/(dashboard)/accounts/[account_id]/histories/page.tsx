import ArticleCard from '@/app/(home)/_components/ArticleCard'
import { articleAPI } from '@/app/_functions/article'
import { authServerFunc } from '@/app/_server_functions/auth'

export default async function HistoryPage() {
    const { token } = await authServerFunc.GetAuth()
    const { data: articles } = await articleAPI.GetReadArticles(token)
    return (
        <div className="flex flex-col space-y-3 p-5">
            {articles && articles?.map(article => (
                <ArticleCard key={article.id} article={article} />
            ))}
        </div>
    )
}
import ArticleCard from '@/app/(home)/_components/ArticleCard'
import { articleAPI } from '@/app/_functions/article'
import { serverAuthFunc } from '@/app/_server_actions/auth'

export default async function HistoryPage() {
    const token = await serverAuthFunc.GetAccessToken()
    const { data: articles } = await articleAPI.GetReadArticles(token)
    return (
        <div className="flex flex-col space-y-3 p-5">
            {articles && articles?.map(article => (
                <ArticleCard key={article.id} article={article} />
            ))}
        </div>
    )
}
import ArticleCard from '@/app/(home)/_components/ArticleCard'
import { Article } from '@/app/_models'

export default function RelatedArticles({ articles }: { articles: Article[] }) {
    return (
        <div className='flex flex-col space-y-3 bg-slate-50 rounded-xl p-3'>
            <div className="text-slate-600">関連記事</div>
            {articles.map(article => {
                return <ArticleCard key={article.id} article={article} />
            })}
        </div>
    )
}
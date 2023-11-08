import ArticleCard from '@/app/(home)/_components/ArticleCard'
import { Article } from '@/types/model'

interface Props {
    articles: Article[]
    origin?: string
}

export default function RelatedArticles({ articles, origin }: Props) {
    return (
        <div className='grid sm:grid-cols-2 gap-3 p-3'>
            <div className="text-slate-600">関連記事</div>
            {articles.map(article => {
                return <ArticleCard key={article.id} article={article} origin={origin} />
            })}
        </div>
    )
}
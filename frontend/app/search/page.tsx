import ArticleCard from '../(home)/_components/ArticleCard'
import { searchAPI } from './_funcitons/search'

interface Props {
    searchParams: { "q": string }
}

export default async function SearchPage({ searchParams }: Props) {
    const query = searchParams.q
    const { data: articles } = await searchAPI.findArticlesByTitle(query)
    return (
        <div className="w-full grid sm:grid-cols-2 xl:grid-cols-3 gap-2 p-[20px]">
            {articles?.map(article => (
                <ArticleCard article={article} key={article.id + "article_card"} />
            ))}
        </div>
    )
}
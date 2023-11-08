import LoaderArticleCard from '@/app/(home)/_components/LoaderArticleCard'

export const RelatedArticlesLoader = () => {
    return (
        <div className='grid sm:grid-cols-2 gap-3 p-3'>
            <div className="text-slate-600">関連記事</div>
            {Array(4).fill(null).map((_, index) => (
                <LoaderArticleCard key={index + "loader"} />
            ))}
        </div>
    )
}
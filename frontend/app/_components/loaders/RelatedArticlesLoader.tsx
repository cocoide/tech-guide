import LoaderArticleCard from '@/app/(home)/_components/LoaderArticleCard'

export const RelatedArticlesLoader = () => {
    return (
        <div className='grid sm:grid-cols-2 lg:grid-cols-3 gap-3 p-3'>
            {Array(4).fill(null).map((_, index) => (
                <LoaderArticleCard key={index + "loader"} />
            ))}
        </div>
    )
}
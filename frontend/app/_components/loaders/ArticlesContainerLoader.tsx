import LoaderArticleCard from '@/app/(home)/_components/LoaderArticleCard'

const ArticlesContainerLoader = () => {
    return (
        <div className="min-h-screen w-full grid sm:grid-cols-2 xl:grid-cols-3  gap-[20px] py-[20px] px-[30px] sm:p-[20px]">
            {Array(6).fill(null).map((_, index) => (
                <LoaderArticleCard key={index + "loader"} />
            ))}</div>
    )
}
export default ArticlesContainerLoader
import LoaderArticleCard from '../(home)/_components/LoaderArticleCard';

export default function Loading() {
    return (
        <div className="w-full grid sm:grid-cols-2 xl:grid-cols-3 gap-2 p-[20px]">
            {Array(10).fill(null).map((_, index) => (
                <LoaderArticleCard key={index + "loader"} />
            ))}
        </div>
    )
}
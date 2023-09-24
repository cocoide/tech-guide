import LoaderArticleCard from '@/app/(home)/_components/LoaderArticleCard';
import SectionHeader from '@/app/_components/layouts/desktop/SectionHeader';



export default function Loading() {
    return (
        <div className="flex flex-col w-full pb-10">
            <div className="h-10 bg-white dark:bg-black w-full border-b-[0.5px] custom-border-color flex-row items-center">
                <SectionHeader titleItem={<div className="h-6 w-14 rounded-full animate-pulse bg-slate-100 dark:bg-slate-800
                "></div>} rightItem={<div></div>} />
            </div>
            <div className="w-full grid sm:grid-cols-2 xl:grid-cols-3 gap-2 p-[20px]">
                {Array(10).fill(null).map((_, index) => (
                    <LoaderArticleCard key={index + "loader"} />
                ))}
            </div>

        </div>
    )
}
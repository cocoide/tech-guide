import SearchSection from './_components/SearchSectioin';

export default function Loading() {
    return (
        <div className="flex flex-col space-y-10 w-full p-10">
            <SearchSection />
            <div className="space-y-3">
                <div className="">トピック一覧</div>
                <div className="w-full flex flex-wrap gap-3">
                    {Array(20).fill(null)?.map((_, index) =>
                        <div key={index + "topic"}
                            className="h-[20px] w-[50px] flex-shrink-0 p-1 bg-slate-100 dark:bg-slate-700 text-slate-600 dark:text-slate-200 rounded-xl"
                        ></div>
                    )}
                </div>
            </div>


            <div className="space-y-3">
                <div className="">ドメイン一覧</div>
                <div className="w-full flex flex-wrap gap-3">
                    {Array(20).fill(null)?.map((_, index) =>
                        <div key={index + "source"}
                            className="h-[20px] w-[50px] flex-shrink-0 p-1 bg-slate-100 dark:bg-slate-700 text-slate-600 dark:text-slate-200 rounded-xl"
                        ></div>
                    )}
                </div>
            </div>
        </div>
    )
}
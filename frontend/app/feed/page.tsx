
import FeedItems from './_components/FeedItems'
import { FetchFeed } from './_functions/feed'

export default async function FeedPage() {
    const feeds = await FetchFeed()
    return (
        <FeedItems initialItems={feeds} fetchItems={FetchFeed}/>
    )
}


// export default function FeedPage() {
//     const myRef = useRef(null)
//     const { token } = useAuth()
//     const { data: articles, fetchNextPage, isFetchingNextPage, isLoading } = useInfiniteQuery({
//         queryKey: ['feeds_query'],
//         queryFn: async ({ pageParam = 1 }) => await articleAPI.GetFeedsByPagination(pageParam, token),
//     }
//     )
//     useEffect(() => {
//         const observer = new IntersectionObserver(
//             (entries) => {
//                 entries.forEach(e => fetchNextPage())
//             })
//         if (myRef.current) {
//             observer.observe(myRef.current)
//         }
//         // eslint-disable-next-line react-hooks/exhaustive-deps
//     }, [myRef])

//     return (
//         <div className="flex flex-col w-full pb-10 relative">
//             <div className="sticky top-0 h-12 bg-white/70 dark:bg-black/30 dark:text-slate-300 backdrop-blur-[5px] z-20">
//                 <SectionHeader
//                     titleItem={<div className='custom-badge text-gray-500'><NewspaperIcon className='h-5 w-5' /><div>フィード</div></div>}
//                     rightItem={
//                         <TopicDialogButton />
//                     }
//                 />
//             </div>
//             <div className="w-full grid sm:grid-cols-2 xl:grid-cols-3 gap-2 p-[20px]">
//                 {isLoading ?
//                     <>
//                         {Array(10).fill(null).map((_, index) => (
//                             <LoaderArticleCard key={index + "loader"} />
//                         ))}
//                     </>
//                     :
//                     <>
//                         {articles?.pages.map(page => (
//                             page?.map((article, index) => (
//                                 <ArticleCard key={article.title + index} article={article} />
//                             )
//                             ))
//                         )}
//                         {isFetchingNextPage &&
//                             Array(2).fill(null).map((_, index) => (
//                                 <LoaderArticleCard key={index + "loader"} />
//                             ))
//                         }
//                     </>
//                 }
//             </div>
//         </div>

//     )
// }
"use client"
import ArticleCard from '@/app/(home)/_components/ArticleCard'
import CircleLoading from '@/app/_components/animations/CircleLoading'
import SectionHeader from '@/app/_components/layouts/desktop/SectionHeader'
import { articleAPI } from '@/app/_functions/article'
import { NewspaperIcon } from '@heroicons/react/24/outline'
import { useInfiniteQuery } from "@tanstack/react-query"
import { useEffect, useRef } from "react"
import TopicDialogButton from '../../_components/layouts/button/FeedFileterDialogButton'

export default function FeedSection() {

    const myRef = useRef(null)
    const { data: articles, fetchNextPage, isFetchingNextPage } = useInfiniteQuery({
        queryKey: ['feeds_query'],
        queryFn: async ({ pageParam = 1 }) => await articleAPI.GetFeedsByPagination(pageParam),
    }
    )
    useEffect(() => {
        const observer = new IntersectionObserver(
            (entries) => {
                entries.forEach(e => fetchNextPage())
            })
        if (myRef.current) {
            observer.observe(myRef.current)
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [myRef])

    return (
        <div className="flex flex-col w-full pb-10 relative">
            <div className="sticky top-0 h-12 bg-white/70 dark:bg-black/30 dark:text-slate-300 backdrop-blur-[5px] z-20">
                <SectionHeader
                    titleItem={<div className='custom-badge text-gray-500'><NewspaperIcon className='h-5 w-5' /><div>フィード</div></div>}
                    rightItem={<TopicDialogButton />}
                />
            </div>
            <div className="w-full grid sm:grid-cols-2 xl:grid-cols-3 gap-2 p-[20px]">
                {articles?.pages.map(page => (
                    page?.map((article, index) => (
                        <ArticleCard key={article.title + index} article={article} />
                    )
                    ))
                )}
                <span ref={myRef}></span>
            </div>
            {isFetchingNextPage &&
                <div className="flex flex-row items-center justify-center w-full h-[200px]"
                ><CircleLoading /></div>
            }
        </div>
    )
}
"use client"
import { HomeIcon } from '@heroicons/react/24/outline'
import { useInfiniteQuery } from "@tanstack/react-query"
import { useEffect, useRef } from "react"
import CircleLoading from '../_components/animations/CircleLoading'
import ToggleDarkModeButton from '../_components/layouts/button/ToggleDarkModeButton'
import SectionHeader from '../_components/layouts/desktop/SectionHeader'
import { articleAPI } from '../_functions/article'
import ArticleCard from './_components/ArticleCard'

export default function ArticlePage() {
    const myRef = useRef(null)
    const { data: articles, fetchNextPage, isFetchingNextPage } = useInfiniteQuery(
        ['trend_query'],
        async ({ pageParam = 1 }) => await articleAPI.GetArticlesByPagination(pageParam),
        {
            getNextPageParam: (_, pages) => pages.length + 1
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
        <div className="flex flex-col w-full pb-10 bg-gray-50 dark:bg-black">
            <div className="sticky top-0 h-10 bg-white/80 dark:bg-black/30 dark:text-slate-300
             backdrop-blur-[5px] z-10">
                <SectionHeader
                    titleItem={<div className='custom-badge text-gray-500 dark:text-white'>
                        <HomeIcon className='h-5 w-5' /><div>最新の投稿</div>
                    </div>}
                    rightItem={<ToggleDarkModeButton />} />
            </div>
            <div className="min-h-screen w-full grid sm:grid-cols-2 xl:grid-cols-3 gap-2 p-3">
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
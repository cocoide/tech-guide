"use client"
import { useInfiniteQuery } from "@tanstack/react-query"
import { useEffect, useRef } from "react"
import CircleLoading from '../_components/animations/CircleLoading'
import TopicDialogButton from '../_components/layouts/components/FeedFileterDialogButton'
import SectionHeader from '../_components/layouts/desktop/SectionHeader'
import { articleAPI } from '../_functions/article'
import ArticleCard from '../trend/_components/ArticleCard'

export default function ArticlePage() {
    const myRef = useRef(null)
    const { data: articles, fetchNextPage, isFetchingNextPage } = useInfiniteQuery(
        ['feed_query'],
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
        <div className="flex flex-col w-full pb-10">
            <SectionHeader title={"フィード"} rightItem={<TopicDialogButton />} />
            <div className="min-h-screen w-full divide-y-[0.5px]">
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
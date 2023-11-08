"use client"

import { articleAPI } from '@/app/_functions/article'
import { useInfiniteQuery } from '@tanstack/react-query'
import { useEffect, useRef } from 'react'
import ArticleCard from '../_components/ArticleCard'
import LoaderArticleCard from '../_components/LoaderArticleCard'

const LatestArticleContainer = () => {
    const myRef = useRef(null)
    const { data: articles, fetchNextPage, isFetchingNextPage, isInitialLoading } = useInfiniteQuery(
        ['latest_query'],
        async ({ pageParam = 1 }) => await articleAPI.GetLatestArticlesByPagination(pageParam),
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
        <div className="min-h-screen w-full grid sm:grid-cols-2 xl:grid-cols-3  gap-[20px] py-[20px] px-[30px] sm:p-[20px]">
            {isInitialLoading ?
                <>
                    {Array(10).fill(null).map((_, index) => (
                        <LoaderArticleCard key={index + "loader"} />
                    ))}
                </>
                :
                <>
                    {articles?.pages.map(page => (
                        page?.map((article, index) => (
                            <ArticleCard key={article.title + index} article={article} />
                        )
                        ))
                    )}
                    {isFetchingNextPage &&
                        Array(2).fill(null).map((_, index) => (
                            <LoaderArticleCard key={index + "loader"} />
                        ))
                    }
                </>
            }
            <span ref={myRef}></span>
        </div>
    )
}
export default LatestArticleContainer
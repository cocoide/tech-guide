"use client"
import { useInfiniteQuery } from "@tanstack/react-query"
import { useEffect, useRef } from "react"
import ToggleDarkModeButton from '../_components/layouts/button/ToggleDarkModeButton'
import SectionHeader from '../_components/layouts/desktop/SectionHeader'
import { articleAPI } from '../_functions/article'
import ArticleCard from './_components/ArticleCard'
import LoaderArticleCard from './_components/LoaderArticleCard'
import SectionHeaderButtonGroup from './_components/SectionHeaderButtonGroup'

export default function ArticlePage() {
    const myRef = useRef(null)
    const { data: articles, fetchNextPage, isFetchingNextPage, isInitialLoading } = useInfiniteQuery(
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
        <div className="flex flex-col w-full pb-10 dark:bg-black">
            <div className="sticky top-0 h-10 bg-white/70 dark:bg-black/30 dark:text-slate-300
             backdrop-blur-[5px] z-20">
                <SectionHeader
                    titleItem={<SectionHeaderButtonGroup />}
                    rightItem={<ToggleDarkModeButton />} />
            </div>
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
        </div>
    )
}
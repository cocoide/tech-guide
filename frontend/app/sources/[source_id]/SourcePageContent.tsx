"use client"
import { sourceAPI } from '@/app/(dashboard)/settings/_functions/source'
import ArticleCard from '@/app/(home)/_components/ArticleCard'
import CircleLoading from '@/app/_components/animations/CircleLoading'
import SectionHeader from '@/app/_components/layouts/desktop/SectionHeader'
import { articleAPI } from '@/app/_functions/article'
import { useAuth } from '@/hooks/useAuth'
import { loginDialogAtom } from '@/stores/dialog'
import { Source } from '@/types/model'
import { useInfiniteQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
import Image from 'next/image'
import { useEffect, useRef } from 'react'
import toast from 'react-hot-toast'
import HeaderFollowButton from './HeaderFollowButton'

interface Props {
    params: {
        source_id: string
    },
    source?: Source
    isFollowing?: boolean
}
export default function SourcePageContent({ params, source, isFollowing }: Props) {
    const myRef = useRef(null)
    const { status, token } = useAuth()
    const [_, setLoginOpen] = useAtom(loginDialogAtom)
    const { data: articles, fetchNextPage, isFetchingNextPage } = useInfiniteQuery({
        queryKey: [`${params.source_id}.source_articles_query`],
        queryFn: async ({ pageParam = 1 }) => (await articleAPI.GetrArticlesBySourceID(Number(params.source_id), pageParam,)).data,
    })
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

    async function handleDoFollow() {
        if (status === "unauthenticated") {
            setLoginOpen(true)
            return
        }
        const { ok } = await sourceAPI.DoFollowSource(Number(params.source_id), token)
        if (!ok) {
            toast.error("エラーが発生")
        }
    }
    async function handleUnFollow() {
        if (status === "unauthenticated") {
            setLoginOpen(true)
            return
        }
        const { ok } = await sourceAPI.UnFollowSource(Number(params.source_id), token)
        if (!ok) {
            toast.error("エラーが発生")
        }
    }
    return (
        <div className="flex flex-col w-full pb-10">
            <div className="sticky top-0 h-12 bg-white/80 backdrop-blur-[5px] z-10">
                <SectionHeader
                    titleItem={<div className='custom-badge text-gray-500'>
                        <Image src={source?.icon_url as string} alt={source?.name as string} width={100} height={100}
                            className='h-7 w-7 rounded-full' />
                        <div>{source?.name}</div>
                    </div>}
                    rightItem={
                        <HeaderFollowButton
                            doFollowFunc={handleDoFollow}
                            unFollowFunc={handleUnFollow}
                            isFollowing={isFollowing}
                        />} />
            </div>
            <div className="min-h-screen w-full grid sm:grid-cols-2  xl:grid-cols-3 gap-2 p-3">
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
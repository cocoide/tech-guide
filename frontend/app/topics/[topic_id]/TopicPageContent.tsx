"use client"
import { topicAPI } from '@/app/(dashboard)/settings/_functions/topic'
import ArticleCard from '@/app/(home)/_components/ArticleCard'
import CircleLoading from '@/app/_components/animations/CircleLoading'
import SectionHeader from '@/app/_components/layouts/desktop/SectionHeader'
import { articleAPI } from '@/app/_functions/article'

import HeaderFollowButton from '@/app/sources/[source_id]/HeaderFollowButton'
import { useAuth } from '@/hooks/useAuth'
import { loginDialogAtom } from '@/stores/dialog'
import { Topic } from '@/types/model'
import { HashtagIcon } from '@heroicons/react/24/outline'
import { useInfiniteQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
import { useEffect, useRef } from 'react'
import toast from 'react-hot-toast'


interface Props {
    params: {
        topic_id: string
    },
    topic?: Topic
    isFollowing?: boolean
}
export default function TopicPageContent({ params, topic, isFollowing }: Props) {
    const myRef = useRef(null)
    const { status, token } = useAuth()
    const [_, setLoginOpen] = useAtom(loginDialogAtom)
    const { data: articles, fetchNextPage, isFetchingNextPage } = useInfiniteQuery({
        queryKey: [`${params.topic_id}.topic_articles_query`],
        queryFn: async ({ pageParam = 1 }) => (await articleAPI.GetrArticlesByTopicID(Number(params.topic_id), pageParam,)).data,
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
        const { ok } = await topicAPI.DoFollowTopic(Number(params.topic_id), token)
        if (!ok) {
            toast.error("エラーが発生")
        }
    }
    async function handleUnFollow() {
        if (status === "unauthenticated") {
            setLoginOpen(true)
            return
        }
        const { ok } = await topicAPI.UnFollowTopic(Number(params.topic_id), token)
        if (!ok) {
            toast.error("エラーが発生")
        }
    }
    return (
        <div className="flex flex-col w-full pb-10">
            <div className="sticky top-0 h-12 bg-white/80 backdrop-blur-[5px] z-10">
                <SectionHeader
                    titleItem={<div className='custom-badge text-gray-500'>
                        <HashtagIcon className="h-5 w-5"></HashtagIcon>
                        <div>{topic?.name}</div>
                    </div>}
                    rightItem={
                        <HeaderFollowButton
                            doFollowFunc={handleDoFollow}
                            unFollowFunc={handleUnFollow}
                            isFollowing={isFollowing}
                        />} />
            </div>
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
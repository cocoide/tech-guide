"use client"

import SectionHeader from '@/app/_components/layouts/desktop/SectionHeader'
import { topicAPI } from '@/app/_functions/topic'
import HeaderFollowButton from '@/app/sources/[source_id]/HeaderFollowButton'
import { useAuth } from '@/hooks/useAuth'
import { loginDialogAtom } from '@/stores/dialog'
import { Topic } from '@/types/model'
import { useQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
import Image from 'next/image'
import toast from 'react-hot-toast'


const TopicSectionHeader = ({ topic }: { topic: Topic }) => {
    const { data: isFollowing, refetch: refetchFollowingState } = useQuery({
        queryKey: [`${topic.id}.topic_query`],
        queryFn: async () => (await topicAPI.CheckFollow(topic.id, token)).data,
    })
    const { status, token } = useAuth()
    const [_, setLoginOpen] = useAtom(loginDialogAtom)
    async function handleDoFollow() {
        if (status === "unauthenticated") {
            setLoginOpen(true)
            return
        }
        const { ok } = await topicAPI.DoFollow(topic.id, token)
        if (!ok) {
            toast.error("エラーが発生")
        }
        await refetchFollowingState()
    }
    async function handleUnFollow() {
        if (status === "unauthenticated") {
            setLoginOpen(true)
            return
        }
        const { ok } = await topicAPI.UnFollow(topic.id, token)
        if (!ok) {
            toast.error("エラーが発生")
        }
        await refetchFollowingState()
    }


    useQuery
    return (
        <SectionHeader
            titleItem={<div className='custom-badge text-gray-500'>
                <Image src={topic?.icon_url as string} alt={topic?.name as string} width={100} height={100}
                    className='h-7 w-7 rounded-full' />
                <div>{topic?.name}</div>
            </div>}
            rightItem={<HeaderFollowButton
                doFollowFunc={handleDoFollow}
                unFollowFunc={handleUnFollow}
                isFollowing={isFollowing} />}
        />
    )
}
export default TopicSectionHeader
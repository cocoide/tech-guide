"use client"

import SectionHeader from '@/app/_components/layouts/desktop/SectionHeader'
import { topicAPI } from '@/app/_functions/topic'
import HeaderFollowButton from '@/app/sources/[source_id]/HeaderFollowButton'
import { useAuth } from '@/hooks/useAuth'
import { loginDialogAtom } from '@/stores/dialog'
import { Topic } from '@/types/model'
import { HashtagIcon } from '@heroicons/react/24/outline'
import { useQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
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
            titleItem={<div className='custom-badge custom-text'>
                <HashtagIcon className='h-7 w-7 custom-text' />
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
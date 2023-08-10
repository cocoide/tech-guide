"use client"
import { topicAPI } from '@/app/(dashboard)/settings/_functions/topic'
import { useAuth } from '@/hooks/useAuth'
import useCustomQuery from '@/hooks/useCustomQuery'
import { topicDialogAtom } from '@/stores/dialog'
import { XMarkIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import TopicFollowSection from './TopicFollowSection'

export default function DialogContent() {
    const { token } = useAuth();
    const [dialogOpen, setDialogOpen] = useAtom(topicDialogAtom)
    const { data: followingTopics } = useCustomQuery(topicAPI.GetFollowingTopics(token), dialogOpen, [dialogOpen])
    const { data: existingTopics } = useCustomQuery(topicAPI.GetAllTopics(), dialogOpen, [dialogOpen])
    return (
        <div className='flex flex-col w-full h-full overflow-y-scroll p-3 space-y-5 relative'>
            <button onClick={() => setDialogOpen(false)}
            ><XMarkIcon className='h-6 w-6 text-gray-500 absolute right-3 top-3' /></button>
            <div className="w-full text-center text-gray-500">トピックの設定</div>
            <TopicFollowSection followingTopics={followingTopics} existingTopics={existingTopics} />
        </div>
    )
}
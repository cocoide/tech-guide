"use client"
import { topicAPI } from '@/app/(dashboard)/settings/_functions/topic'
import FollowTopicsSection from '@/app/(dashboard)/settings/topics/components/FollowTopicsSection'
import RecommendTopicsSection from '@/app/(dashboard)/settings/topics/components/RecommendTopicsSection'
import { useAuth } from '@/hooks/useAuth'
import { topicDialogAtom } from '@/stores/dialog'
import { Topic } from '@/types/model'
import { useEffect, useState } from 'react'
import CustomDialog from '../elements/CustomDialog'

const TopicDialog = () => {
    const [followTopic, setFollowTopic] = useState<Topic[]>()
    const [existTopic, setExistTopic] = useState<Topic[]>()
    const { token } = useAuth()
    useEffect(() => {
        if (token) {
            (async () => {
                const { data: topics } = await topicAPI.GetFollowingTopics(token)
                setFollowTopic(topics)
            })()
        }
        (async () => {
            const { data: topics } = await topicAPI.GetAllTopics()
            setExistTopic(topics)
        })()
    }, [token])
    const unfollow_topics = existTopic?.filter(topic => {
        return !followTopic?.some(followTopic => followTopic.id === topic.id);
    });
    return (
        <CustomDialog
            openAtom={topicDialogAtom}
            layout='mt-[120px] sm:mb-[120px]  bg-white z-50 sm:mx-[15%] md:mx-[20%] lg:mx-[25%] rounded-xl'
            content={
                <div className='flex flex-col w-full h-full overflow-y-scroll p-3 space-y-3'>
                    {followTopic && followTopic.length ?
                        <FollowTopicsSection follow_topics={followTopic} />
                        :
                        <div className="text-slate-500">フォロー中のトピックがありません</div>
                    }
                    {unfollow_topics &&

                        <RecommendTopicsSection unfollow_topics={unfollow_topics} />
                    }
                </div>
            }
        />
    )
}
export default TopicDialog
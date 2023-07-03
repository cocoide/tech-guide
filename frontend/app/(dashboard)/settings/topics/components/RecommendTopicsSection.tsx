'use client'
import { useAuth } from '@/hooks/useAuth';
import { Topic } from '@/types/model';
import { PlusIcon } from '@heroicons/react/24/outline';
import { useEffect, useState } from 'react';
import { toast } from 'react-hot-toast';
import { topicAPI } from '../../_functions/topic';

const RecommendTopicsSection = ({ unfollow_topics }: { unfollow_topics: Topic[] }) => {
    const [isFollowing, setIsFollowing] = useState<{ [topic_id: number]: boolean }>({});
    const { token } = useAuth()
    useEffect(() => {
        const initialFollowStatus: { [topic_id: number]: boolean } = {};
        unfollow_topics.forEach(topic => {
            initialFollowStatus[topic.id] = false;
        });
        setIsFollowing(initialFollowStatus);
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])

    async function handleDoFollow(topic_id: number) {
        setIsFollowing(prevState => ({
            ...prevState,
            [topic_id]: true
        }));
        const { ok } = await topicAPI.DoFollowTopic(topic_id, token)
        if (!ok) {
            toast.error("エラーが発生")
        }
    }
    async function handleUnFollow(topic_id: number) {
        setIsFollowing(prevState => ({
            ...prevState,
            [topic_id]: false
        }));
        const { ok } = await topicAPI.UnFollowTopic(topic_id, token)
        if (!ok) {
            toast.error("エラーが発生")
        }
    }
    return (
        <>
            <div className='mt-5 text-center text-cyan-300 border-b border-cyan-300 mx-auto py-1 px-2'
            >おすすめのトピック</div>
            <div className="w-full flex flex-wrap gap-3">
                {unfollow_topics?.map(topic =>
                    <div className="" key={topic.id}>
                        {!isFollowing[topic.id] === false ?
                            <button onClick={() => handleUnFollow(topic.id)}
                                className="flex flex-row items-center space-x-[3px] text-white bg-cyan-300 rounded-full p-[6px]">
                                <div className="text-sm"> {topic.name}</div>
                                <PlusIcon className='h-4 w-4 text-cyan-white' />
                            </button>
                            :
                            <button onClick={() => handleDoFollow(topic.id)}
                                className="flex flex-row items-center space-x-[3px] text-slate-700 ring-1 ring-slate-200 rounded-full p-[6px]">
                                <div className="text-sm"> {topic.name}</div>
                                <PlusIcon className='h-4 w-4 text-cyan-300' />
                            </button>
                        }
                    </div>
                )}
            </div>
        </>
    )
}
export default RecommendTopicsSection
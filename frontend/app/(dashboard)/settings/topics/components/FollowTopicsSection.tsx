'use client'
import { useAuth } from '@/hooks/useAuth';
import { Topic } from '@/types/model';
import { ChatBubbleOvalLeftEllipsisIcon } from '@heroicons/react/24/outline';
import { useEffect, useState } from 'react';
import { toast } from 'react-hot-toast';
import { topicAPI } from '../../_functions/topic';

const FollowTopicsSection = ({ follow_topics }: { follow_topics: Topic[] }) => {
    const [isFollowing, setIsFollowing] = useState<{ [topic_id: number]: boolean }>({});
    const { token } = useAuth()
    useEffect(() => {
        const initialFollowStatus: { [topic_id: number]: boolean } = {};
        follow_topics.forEach(topic => {
            initialFollowStatus[topic.id] = true;
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
        <div className="flex flex-col space-y-[3px] w-full">
            <div className='text-center text-cyan-300 border-b border-cyan-300 mx-auto py-1 px-2'
            >フォロー中のトピック</div>
            {follow_topics?.map(topic => (
                <div key={topic.id} className="flex items-center p-[5px]  w-full justify-between">
                    <div className="flex items-center space-x-3">
                        <ChatBubbleOvalLeftEllipsisIcon className='h-5 w-5 text-cyan-300' />
                        <div className="text-slate-700"> {topic.name}</div>
                    </div>
                    {isFollowing[topic.id] ?
                        <button onClick={() => handleUnFollow(topic.id)} className='bg-slate-400 text-white rounded-full p-[5px] text-sm'
                        >フォロー中</button>
                        :
                        <button onClick={() => handleDoFollow(topic.id)} className='bg-slate-100 text-slate-400 ring-1 ring-slate-400  rounded-full p-[5px] text-sm'
                        >フォローする</button>
                    }
                </div>
            ))}
        </div>
    )
}
export default FollowTopicsSection
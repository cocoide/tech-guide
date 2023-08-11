import { topicAPI } from '@/app/(dashboard)/settings/_functions/topic';
import { Topic } from '@/types/model';
import { PlusIcon, XMarkIcon } from '@heroicons/react/24/outline';
import { useEffect, useState } from 'react';
import toast from 'react-hot-toast';

interface Props {
    token?: string
    followingTopics?: Topic[]
    existingTopics?: Topic[]
}
const TopicFollowSection = ({ token, followingTopics, existingTopics }: Props) => {

    const [topicMap, setTopicMap] = useState<{ [topic_id: number]: boolean }>({});
    useEffect(() => {
        const initialFollowStatus: { [topic_id: number]: boolean } = {};
        existingTopics?.forEach(topic => {
            initialFollowStatus[topic.id] = false;
        });
        followingTopics?.forEach(topic => {
            initialFollowStatus[topic.id] = true;
        });
        setTopicMap(initialFollowStatus);
    }, [existingTopics, followingTopics])

    async function handleDoFollow(topic_id: number) {
        setTopicMap(prevState => ({
            ...prevState,
            [topic_id]: true
        }));
        const { ok } = await topicAPI.DoFollowTopic(topic_id, token)
        if (!ok) {
            toast.error("エラーが発生")
        }
    }
    async function handleUnFollow(topic_id: number) {
        setTopicMap(prevState => ({
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
            <div className="w-full flex flex-wrap gap-3">
                {existingTopics?.map(topic =>
                    <div className="" key={topic.id}>
                        {!topicMap[topic.id] === false ?
                            <button onClick={() => handleUnFollow(topic.id)}
                                className="flex flex-row items-center space-x-[3px] text-white bg-gray-400 rounded-full p-[6px]">
                                <div className="text-sm"> {topic.name}</div>
                                <PlusIcon className='h-4 w-4' />
                            </button>
                            :
                            <button onClick={() => handleDoFollow(topic.id)}
                                className="flex flex-row items-center space-x-[3px] text-gray-700 ring-1 ring-gray-200 rounded-full p-[6px]">
                                <div className="text-sm"> {topic.name}</div>
                                <XMarkIcon className='h-4 w-4 text-gray-400' />
                            </button>
                        }
                    </div>
                )}
            </div>
        </>
    )
}
export default TopicFollowSection
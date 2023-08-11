import { sourceAPI } from '@/app/(dashboard)/settings/_functions/source';
import { Source } from '@/types/model';
import { PlusIcon, XMarkIcon } from '@heroicons/react/24/outline';
import Image from 'next/image';
import { useEffect, useState } from 'react';
import toast from 'react-hot-toast';

interface Props {
    existingSources?: Source[]
    followingSources?: Source[]
    token?: string
}
const DomainFollowSection = ({ existingSources, followingSources, token }: Props) => {
    const [sourceMap, setSourceMap] = useState<{ [source_id: number]: boolean }>({});
    useEffect(() => {
        const initialSourceStatus: { [topic_id: number]: boolean } = {};
        existingSources?.forEach(topic => {
            initialSourceStatus[topic.id] = false;
        });
        followingSources?.forEach(topic => {
            initialSourceStatus[topic.id] = true;
        });
        setSourceMap(initialSourceStatus);
    }, [existingSources, followingSources])

    async function handleDoFollow(topic_id: number) {
        setSourceMap(prevState => ({
            ...prevState,
            [topic_id]: true
        }));
        const { ok } = await sourceAPI.DoFollowSource(topic_id, token)
        if (!ok) {
            toast.error("エラーが発生")
        }
    }
    async function handleUnFollow(topic_id: number) {
        setSourceMap(prevState => ({
            ...prevState,
            [topic_id]: false
        }));
        const { ok } = await sourceAPI.UnFollowSource(topic_id, token)
        if (!ok) {
            toast.error("エラーが発生")
        }
    }
    return (
        <>
            <div className="w-full flex flex-wrap gap-3">
                {existingSources?.map(topic =>
                    <div className="" key={topic.id}>
                        {!sourceMap[topic.id] === false ?
                            <button onClick={() => handleUnFollow(topic.id)}
                                className="flex flex-row items-center space-x-[3px] text-white bg-gray-400 rounded-full p-[6px]">
                                <div className="custom-badge">
                                    <Image src={topic.icon_url} alt={topic.name} width={100} height={100} className='w-5 h-5 rounded-full' />
                                    <div className="">{topic.name}</div>
                                </div>
                                <PlusIcon className='h-4 w-4' />
                            </button>
                            :
                            <button onClick={() => handleDoFollow(topic.id)}
                                className="flex flex-row items-center space-x-[3px] text-gray-700 ring-1 ring-gray-200 rounded-full p-[6px]">
                                <div className="custom-badge">
                                    <Image src={topic.icon_url} alt={topic.name} width={100} height={100} className='w-5 h-5 rounded-full' />
                                    <div className="">{topic.name}</div>
                                </div>
                                <XMarkIcon className='h-4 w-4 text-gray-400' />
                            </button>
                        }
                    </div>
                )}
            </div>
        </>
    )
}
export default DomainFollowSection
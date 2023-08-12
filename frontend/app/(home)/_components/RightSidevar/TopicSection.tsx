import { HashtagIcon } from '@heroicons/react/24/outline';
import { popularAPI } from '../../../_functions/popular';
import Link from 'next/link';

export default async function TopicSection() {
    const { data: topics } = await popularAPI.GetPopularTopics()
    return (
        <div className="rounded-xl custom-border p-2 flex flex-col space-y-2 text-gray-500">
            <div className="flex flex-row items-center space-x-2">
                <HashtagIcon className='h-5 w-5' />
                <div className="">人気のトピック</div>
            </div>
            {topics?.map(topic => (
                <Link href={`/topics/${topic.id}`} key={topic.id} className="custom-border bg-gray-50 text-gray-500 mr-auto p-1 rounded-xl text-sm">
                    {topic.name}</Link>
            ))}
        </div>
    )
}
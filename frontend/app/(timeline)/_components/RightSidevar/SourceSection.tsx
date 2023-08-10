import { GlobeAltIcon } from '@heroicons/react/24/outline';
import { popularAPI } from '../../_functions/popular';

export default async function SourceSection() {
    const { data: topics } = await popularAPI.GetPopularSources()
    return (
        <div className="rounded-xl custom-border p-2 flex flex-col space-y-2 text-gray-500">
        <div className="flex flex-row items-center space-x-2">
            <GlobeAltIcon className='h-5 w-5' />
            <div className="">Popular Source</div>
        </div>
        {topics?.map(topic => (
                <div key={topic.id} className="">{topic.name}</div>
            ))}
    </div>
    )
}
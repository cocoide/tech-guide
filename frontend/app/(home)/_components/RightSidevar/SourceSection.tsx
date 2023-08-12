import { GlobeAltIcon } from '@heroicons/react/24/outline';
import { popularAPI } from '../../../_functions/popular';
import Link from 'next/link';
import Image from 'next/image';

export default async function SourceSection() {
    const { data: sources } = await popularAPI.GetPopularSources()
    return (
        <div className="rounded-xl custom-border p-2 flex flex-col space-y-2 text-gray-500">
            <div className="flex flex-row items-center space-x-2">
                <GlobeAltIcon className='h-5 w-5' />
                <div className="">人気のドメイン</div>
            </div>
            {sources?.map(source => (
                <Link href={`/sources/${source.id}`} key={source.id} className="custom-border  custom-badge bg-gray-50 text-gray-500 mr-auto p-1 rounded-xl text-sm">
                    <Image src={source.icon_url} alt={source.name} width={100} height={100} className='h-5 w-5 rounded-full' />
                            <div className="">{source.name}</div>
                    </Link>
            ))}
        </div>
    )
}
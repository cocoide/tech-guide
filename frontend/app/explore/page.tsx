import Image from 'next/image';
import { sourceAPI } from '../(dashboard)/settings/_functions/source';
import { topicAPI } from './_functions/topic';
import Link from 'next/link';

export default async function ExplorePage() {
    const { data: topics } = await topicAPI.GetAllTopics()
    const { data: sources } = await sourceAPI.GetAllSources()
    return (
        <div className="flex flex-col space-y-3 w-full">
            <div className="p-10 space-y-3">
                <div className="">トピック一覧</div>
                <div className="w-full flex flex-wrap gap-3">
            {topics?.map(topic =>
                <Link href={`/topics/${topic.id}`}  key={topic.id} className="flex-shrink-0 p-1 bg-slate-100 text-slate-600 rounded-xl"
                >{topic.name}</Link>
            )}
                </div>
            </div>


            <div className="p-10 space-y-3">
                <div className="">ドメイン一覧</div>
                <div className="w-full flex flex-wrap gap-3">
                    {sources?.map(source =>
                        <Link href={`/sources/${source.id}`} key={source.id} className="flex-shrink-0 p-1 bg-slate-100 text-slate-600 rounded-xl custom-badge"
                        ><Image src={source.icon_url} alt={source.name} width={100} height={100} className='h-5 w-5 rounded-full' />
                            <div className="">{source.name}</div>
                        </Link>
                    )}
                </div>
            </div>
        </div>
    )
}
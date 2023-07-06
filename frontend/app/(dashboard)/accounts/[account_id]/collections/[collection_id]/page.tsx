import ArticleCard from '@/app/(timeline)/_components/ArticleCard';
import { ChevronLeftIcon, FolderIcon } from '@heroicons/react/24/outline';
import Link from 'next/link';
import { authAPI } from '../../../../../_functions/auth';
import { collectionAPI } from './_functions/collection';

interface Porps {
    params: {
        collection_id: string
    }
}
export default async function CollectionPage({ params }: Porps) {
    const { collection_id } = params
    const { data: collection } = await collectionAPI.GetCollectionData(collection_id)
    const { user } = await authAPI.GetAuthSession()
    return (
        <div className="w-full lg:px-[10%] flex flex-col space-y-3 p-5">
            <div className="text-slate-600 text-center w-full flex flex-row items-center space-x-3">
                <Link href={`/accounts/${user.uid}/collections`}>
                    <ChevronLeftIcon className="h-5 w-5"></ChevronLeftIcon>
                </Link>
                <FolderIcon className='h-5 w-5' />
                <div>  {collection?.name}</div>
            </div>
            <div className="grid lg:grid-cols-2 w-full bg-cyan-50/70 rounded-xl gap-3">
                {collection?.articles.map(article => (
                    <ArticleCard key={article.title} article={article} />
                ))}
            </div>
        </div>
    )
}
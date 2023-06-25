'use client'

import { Article } from '@/app/_models'
import { collectionDialogAtom } from '@/stores/dialog'
import { ArrowTopRightOnSquareIcon, FolderPlusIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Link from 'next/link'

const ArticleCard = ({ article }: { article: Article }) => {
    const [_, setOpenCollectionDialog] = useAtom(collectionDialogAtom)

    return (
        <div className='p-3'>
            <div className='flex flex-col space-y-3'>
                <div className='flex flex-row justify-between'>
                    <div className="flex flex-col">
                        <div>{article.title}</div>
                        {article.topics?.length > 0 &&
                            <div className='flex flex-row space-x-3'>{article.topics.map(topic =>
                                (<div key={topic.name} className="">#{topic.name}</div>))}</div>
                        }
                    </div>
                    {article.thumbnail_url &&
                        // eslint-disable-next-line @next/next/no-img-element
                        <img src={article.thumbnail_url} alt={article.title} width={200} />
                    }
                </div>
                <div className='flex flex-row justify-between'>
                    <div className='flex flex-row space-x-3'>
                        <button onClick={() => setOpenCollectionDialog(article.id)}>
                            <FolderPlusIcon className='text-gray-500 h-5 w-5' />
                        </button>
                    </div>
                    <Link href={article.original_url} passHref>
                        <ArrowTopRightOnSquareIcon className='text-cyan-300 h-5 w-5' />
                    </Link>
                </div>
            </div>
        </div>
    )
}
export default ArticleCard
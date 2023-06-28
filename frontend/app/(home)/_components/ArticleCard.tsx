'use client'

import { Article } from '@/app/_models'
import { collectionDialogAtom } from '@/stores/dialog'
import { ArrowTopRightOnSquareIcon, FolderPlusIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Link from 'next/link'


interface Props {
    article: Article
    origin?: string
}
const ArticleCard = ({ article, origin }: Props) => {
    const [_, setOpenCollectionDialog] = useAtom(collectionDialogAtom)
    const queryParam = origin ? `?exclude=${origin}` : ''
    return (
        <div className='p-3'>
            <div className='flex flex-col space-y-3'>
                <Link href={`/articles/${article.id}${queryParam}`} className='flex flex-row justify-between'>
                    <div className="flex flex-col">
                        <div>{article.title}</div>
                    </div>
                    {article.thumbnail_url &&
                        // eslint-disable-next-line @next/next/no-img-element
                        <img src={article.thumbnail_url} alt={article.title} width={200} />
                    }
                </Link>
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
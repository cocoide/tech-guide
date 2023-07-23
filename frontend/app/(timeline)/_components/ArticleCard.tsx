'use client'

import { useAuth } from '@/hooks/useAuth'
import { collectionDialogAtom, loginDialogAtom } from '@/stores/dialog'
import { Article } from '@/types/model'
import { ArrowTopRightOnSquareIcon, FolderPlusIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Image from 'next/image'
import Link from 'next/link'
import { useEffect, useState } from 'react'
import { toast } from 'react-hot-toast'
import { articleAPI } from '../_functions/article'

interface Props {
    article: Article
    origin?: string
}
const ArticleCard = ({ article, origin }: Props) => {
    const { status, token } = useAuth()
    function handleCollectionDialog() {
        if (status === "authenticated") {
            setOpenCollectionDialog(article.id)
        }
        if (status === "unauthenticated") {
            setOpenLoginDialog(true)
        }
    }
    async function handleOnRead(article_id: number) {
        if (token) {
            const { ok } = await articleAPI.ReadArticle(article_id, token)
            if (!ok) {
                toast.error("エラーが発生")
            }
        }
    }
    const [_, setOpenCollectionDialog] = useAtom(collectionDialogAtom)
    const [__, setOpenLoginDialog] = useAtom(loginDialogAtom)
    const [queryParam, setQueryParam] = useState<string>("")
    useEffect(() => {
        if (origin) {
            let query = '';
            if (origin) {
                query += `?exclude=${origin}`;
            }
            setQueryParam(query);
        }
    }, [origin]);
    return (
        <div className='p-3 hover:bg-slate-50 duaration-500'>
            <div className='flex flex-col space-y-[5px]'>
                <Link href={`/articles/${article.id}${queryParam}`} className='flex flex-row justify-between'>
                    <div className="flex flex-col space-y-[3px]">
                        <div className="flex flex-row space-x-3">
                            {article.source.icon_url &&
                                <Image src={article.source.icon_url} alt={article.source.name} width={100} height={100}
                                    className='rounded-full h-7 w-7' />
                            }
                            <div className='text-slate-700'>{article.title}</div>
                        </div>
                        <div className="text-slate-500 mr-auto">{article.topics?.slice(0, 1).map(topic =>
                            (<div key={topic.id} className='flex flex-row items-center text-[10px] ring-1 rounded-md ring-slate-300 p-[2px]'>{topic.name}</div>)
                        )}</div>
                    </div>
                    {article.thumbnail_url &&
                        // eslint-disable-next-line @next/next/no-img-element
                        <img src={article.thumbnail_url} alt={article.title} width={200} className="rounded-md" />
                    }
                </Link>
                <div className='flex flex-row justify-between'>
                    <div className='flex flex-row space-x-3'>
                        <button onClick={handleCollectionDialog}>
                            <FolderPlusIcon className='text-gray-500 h-5 w-5' />
                        </button>
                    </div>
                    <Link onClick={() => handleOnRead(article.id)} href={article.original_url} passHref>
                        <ArrowTopRightOnSquareIcon className='text-cyan-300 h-5 w-5' />
                    </Link>
                </div>
            </div>
        </div>
    )
}
export default ArticleCard
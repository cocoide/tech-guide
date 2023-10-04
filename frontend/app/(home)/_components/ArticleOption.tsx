"use client"
import { useAuth } from '@/hooks/useAuth'
import { collectionDialogAtom, loginDialogAtom } from '@/stores/dialog'
import { Article } from '@/types/model'
import { BookmarkIcon, ChatBubbleOvalLeftEllipsisIcon, StarIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Link from 'next/link'
import toast from 'react-hot-toast'
import { articleAPI } from '../../_functions/article'

interface Props {
  article: Article
}
const ArticleOption = ({ article }: Props) => {
  const [_, setOpenCollectionDialog] = useAtom(collectionDialogAtom)
  const [__, setOpenLoginDialog] = useAtom(loginDialogAtom)
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
  const rating = article.rating
  const count = rating.hatena_stocks + rating.origin_stocks + rating.owned_stocks + rating.pocket_stocks
  return (
    <div className='flex flex-row items-center absolute bottom-[7px] px-[7px] justify-between w-full'>
      <div className="text-slate-400 dark:text-slate-200 flex flex-row items-center
      space-x-[5px]">
        {count != 0 &&
          <div className="text-slate-500 dark:text-slate-100 font-bold">{count}</div>
        }
      <Link href={`/sources/${article.source.id}`} className="text-gray-400 text-sm">
        {article.source.name}
      </Link>
      </div>
      <div className="text-slate-400 dark:text-slate-200 flex flex-row items-center
      space-x-[5px]">
        <button className='p-1 rounded-full
      hover:text-pink-300  hover:bg-pink-50 duration-500' onClick={handleCollectionDialog}>
          <StarIcon className='h-5 w-5' />
      </button>
        <Link href={`/articles/${article.id}?comment=true`} className='p-1 rounded-full
      hover:text-blue-300  hover:bg-blue-50'>
        <ChatBubbleOvalLeftEllipsisIcon className='h-5 w-5' />
      </Link>
        <button className='p-1 rounded-full 
      hover:text-green-300 hover:bg-green-50' onClick={handleCollectionDialog}>
        <BookmarkIcon className='h-5 w-5' />
        </button>
      </div>
    </div>
  )
}
export default ArticleOption
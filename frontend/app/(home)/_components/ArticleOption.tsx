"use client"
import { useAuth } from '@/hooks/useAuth'
import { collectionDialogAtom, loginDialogAtom } from '@/stores/dialog'
import { Article } from '@/types/model'
import { BookmarkIcon, ChatBubbleOvalLeftEllipsisIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Link from 'next/link'
import toast from 'react-hot-toast'
import { articleAPI } from '../../_functions/article'
import { RatingBadge } from './RatingBadge'

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
    <div className='flex flex-row items-center absolute bottom-[6px] px-[6px] justify-between w-full'>
      <div className="text-slate-400 dark:text-slate-200 flex flex-row items-center
      space-x-[8px]">
        {count != 0 &&
          <RatingBadge rating={article.rating} domain={article.source.domain} />
        }
      </div>
      <div className="text-slate-400 dark:text-slate-200 flex flex-row items-center
      space-x-[5px]">
        <Link href={`/articles/${article.id}?comment=true`} className='p-[5px] rounded-full
      hover:text-blue-300  hover:bg-blue-50'>
        <ChatBubbleOvalLeftEllipsisIcon className='h-6 w-6' />
      </Link>
        <button className='p-[5px] rounded-full 
      hover:text-green-300 hover:bg-green-50' onClick={handleCollectionDialog}>
        <BookmarkIcon className='h-6 w-6' />
        </button>
      </div>
    </div>
  )
}
export default ArticleOption
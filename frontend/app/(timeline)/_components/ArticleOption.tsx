"use client"
import { useAuth } from '@/hooks/useAuth'
import { collectionDialogAtom, loginDialogAtom } from '@/stores/dialog'
import { Article } from '@/types/model'
import { ArrowTopRightOnSquareIcon, BookmarkIcon, ChatBubbleOvalLeftEllipsisIcon, HandThumbUpIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Link from 'next/link'
import toast from 'react-hot-toast'
import { articleAPI } from '../_functions/article'

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
  return (
    <>
      <div className="">

      </div>
      <button className='absolute bottom-[12px] left-[40px]' onClick={handleCollectionDialog}>
        <HandThumbUpIcon className='text-slate-400 h-5 w-5' />
      </button>
      <button className='absolute bottom-[12px] left-[70px]' onClick={handleCollectionDialog}>
        <ChatBubbleOvalLeftEllipsisIcon className='text-slate-400 h-5 w-5' />
      </button>
      <button className='absolute bottom-[12px] left-[100px]' onClick={handleCollectionDialog}>
        <BookmarkIcon className='text-slate-400 h-5 w-5' />
      </button>
      <Link className='absolute bottom-[12px] right-[5px]' onClick={() => handleOnRead(article.id)} href={article.original_url} passHref>
        <ArrowTopRightOnSquareIcon className='text-cyan-300 h-5 w-5' />
      </Link>
    </>
  )
}
export default ArticleOption
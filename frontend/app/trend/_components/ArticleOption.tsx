"use client"
import { useAuth } from '@/hooks/useAuth'
import { collectionDialogAtom, loginDialogAtom } from '@/stores/dialog'
import { Article } from '@/types/model'
import { ArrowTopRightOnSquareIcon, BookmarkIcon, ChatBubbleOvalLeftEllipsisIcon, HandThumbUpIcon } from '@heroicons/react/24/outline'
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
  return (
    <>
      <div className="">

      </div>
      <button className='absolute bottom-[10px] left-[40px] p-1 rounded-full
      text-slate-400 hover:text-pink-300  hover:bg-pink-50 duration-500' onClick={handleCollectionDialog}>
        <HandThumbUpIcon className='h-5 w-5' />
      </button>
      <button className='absolute bottom-[10px] left-[70px] p-1 rounded-full
      text-slate-400 hover:text-blue-300  hover:bg-blue-50 duration-500' onClick={handleCollectionDialog}>
        <ChatBubbleOvalLeftEllipsisIcon className='h-5 w-5' />
      </button>
      <button className='absolute bottom-[10px] left-[100px] p-1 rounded-full 
      text-slate-400 hover:text-green-300 hover:bg-green-50  duration-500' onClick={handleCollectionDialog}>
        <BookmarkIcon className='h-5 w-5' />
      </button>
      <Link className='absolute bottom-[10px] right-[5px] rounded-full p-1
      text-slate-400 hover:text-cyan-300 hover:bg-cyan-50 duration-500' onClick={() => handleOnRead(article.id)} href={article.original_url} passHref>
        <ArrowTopRightOnSquareIcon className='h-5 w-5' />
      </Link>
    </>
  )
}
export default ArticleOption
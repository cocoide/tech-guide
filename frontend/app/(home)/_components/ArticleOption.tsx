"use client"
import { useAuth } from '@/hooks/useAuth'
import { collectionDialogAtom, commentDialogAtom, loginDialogAtom } from '@/stores/dialog'
import { Article } from '@/types/model'
import { ArrowTopRightOnSquareIcon, ChatBubbleOvalLeftEllipsisIcon, StarIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import toast from 'react-hot-toast'
import { articleAPI } from '../../_functions/article'
import PreviewButton from './PreviewButton'
import { RatingBadge } from './RatingBadge'

interface Props {
  article: Article
}
const ArticleOption = ({ article }: Props) => {
  const [_, setOpenCollectionDialog] = useAtom(collectionDialogAtom)
  const [__, setOpenLoginDialog] = useAtom(loginDialogAtom)
  const [___, setOpenCommentDialog] = useAtom(commentDialogAtom)
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
  function handleCommnetDialog(article: Article) {
    if (status === "authenticated") {
      setOpenCommentDialog(article)
    }
    if (status === "unauthenticated") {
      setOpenLoginDialog(true)
    }
  }
  const rating = article.rating
  const count = rating.hatena_stocks + rating.origin_stocks + rating.owned_stocks + rating.pocket_stocks
  const domain = extractDomain(article.original_url)
  return (
    <div className='flex flex-row items-center p-[6px] justify-between w-full'>
      <div className="text-slate-400 dark:text-slate-200 flex flex-row items-center
      space-x-[10px]">
        <PreviewButton url={article.original_url} domain={article.source.domain} summary={article.summary} />
        <button onClick={() => handleCommnetDialog(article)} className='p-[5px] rounded-full
      hover:text-blue-300  hover:bg-blue-50'>
          <ChatBubbleOvalLeftEllipsisIcon className='h-6 w-6' />
        </button>
        <button className='p-[5px] rounded-full 
      hover:text-green-300 hover:bg-green-50' onClick={handleCollectionDialog}>
          <StarIcon className='h-6 w-6' />
        </button>
      </div>
      <div className="text-slate-400 dark:text-slate-200 flex flex-row items-center
      space-x-[8px]">
        {count != 0 &&
          <RatingBadge rating={article.rating} domain={article.source.domain} />
        }
        <a href={article.original_url} target="_blank" className=" flex bg-gray-100 dark:bg-gray-700 cutom-outline
                            text-gray-500 dark:text-gray-200  rounded-xl shadow-sm p-[4px] text-sm animate-appear duration-100 z-10 custom-badge">
          <ArrowTopRightOnSquareIcon className="h-4 w-4"></ArrowTopRightOnSquareIcon>
          <div className="text-sm whitespace-nowrap overflow-x-hidden max-w-[100px]">{domain}</div>
        </a>
      </div>
    </div>
  )
}
export default ArticleOption

function extractDomain(url: string): string | null {
  try {
    const parsedUrl = new URL(url);
    let hostname = parsedUrl.hostname;

    if (hostname.startsWith('www.')) {
      hostname = hostname.slice(4);
    }

    return hostname;
  } catch (e) {
    console.error("Invalid URL provided:", e);
    return null;
  }
}
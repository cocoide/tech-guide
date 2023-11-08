"use client"
import YouTubeEmbed from '@/app/(home)/_components/YoutubeEmbed';
import { articleAPI } from '@/app/_functions/article';
import { useAuth } from '@/hooks/useAuth';
import { collectionDialogAtom, loginDialogAtom, outlineDialogAtom } from '@/stores/dialog';
import { Article } from '@/types/model';
import { useAtom } from 'jotai';
import Link from 'next/link';
import { toast } from 'react-hot-toast';

export default function ArticleDetail({ article }: { article: Article }) {
    const [_, setOpenCollectionDialog] = useAtom(collectionDialogAtom)
    const [__, setOpenLoginDialog] = useAtom(loginDialogAtom)
    const [___, setOpenOutlineDialog] = useAtom(outlineDialogAtom)
    function extractYoutubeID(url: string): string | null {
        const match = url.match(/^(?:https?:\/\/)?(?:www\.)?(?:m\.)?(?:youtu\.be\/|youtube\.com\/(?:watch\?(?:\S*?&)?v=|(?:embed|v|vi|user)\/))([\w-]+)/);
        return match ? match[1] : null;
    }
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
    const youtube_id = extractYoutubeID(String(article?.original_url))
    return (
        <div className="bg-gray-50/50 dark:bg-gray-900 rounded-xl custom-border
        shadow-[0_8px_30px_rgb(0,0,0,0.12)] 
        p-5  flex flex-col lg:flex-row items-center
        space-y-3 lg:space-x-3">
            {youtube_id ?
                <YouTubeEmbed youtube_id={youtube_id} />
                :
                article?.thumbnail_url &&
                // eslint-disable-next-line @next/next/no-img-element
                <img src={article.thumbnail_url} alt={article.title} width={500}
                    className='w-full h-auto lg:w-[50%] rounded-md custom-border' />
            }
            <div className="flex flex-col items-center w-full space-y-2">
                <div className="text-slate-700">{article.title}</div>
                <div className="text-slate-500 mr-auto flex flex-row items-center space-x-2 justify-center w-full">{article.topics?.slice(0, 3).map(topic =>
                    (<div key={topic.id} className='flex flex-row items-center text-[10px] ring-1 rounded-md ring-slate-300 p-[2px]'>{topic.name}</div>)
                )}</div>
                <div className="flex flex-row items-center space-x-3">
                    <Link href={article.original_url} onClick={() => handleOnRead(article.id)} className="bg-slate-100 text-slate-500 p-[5px] rounded-xl text-sm">読む</Link>
                    <button onClick={handleCollectionDialog} className="bg-cyan-300 text-white p-[5px] rounded-xl text-sm">保存</button>
                    <button onClick={() => setOpenOutlineDialog({ original_url: article.original_url })} className="bg-slate-100 text-white p-[5px] rounded-xl text-sm">概要</button>
                </div>
            </div>
        </div>
    )
}
"use client"
import YouTubeEmbed from '@/app/(timeline)/_components/YoutubeEmbed';
import { Article } from '@/app/_models';
import { collectionDialogAtom } from '@/stores/dialog';
import { useAtom } from 'jotai';
import Link from 'next/link';

export default function ArticleDetail({ article }: { article: Article }) {
    const [_, setDialogOpen] = useAtom(collectionDialogAtom)
    function extractYoutubeID(url: string): string | null {
        const match = url.match(/^(?:https?:\/\/)?(?:www\.)?(?:m\.)?(?:youtu\.be\/|youtube\.com\/(?:watch\?(?:\S*?&)?v=|(?:embed|v|vi|user)\/))([\w-]+)/);
        return match ? match[1] : null;
    }
    const youtube_id = extractYoutubeID(String(article?.original_url))
    return (
        <div className="bg-white rounded-xl 
        shadow-[0_8px_30px_rgb(0,0,0,0.12)] 
        p-5  flex flex-col md:flex-row items-center 
        space-y-3 md:space-x-3">
            {youtube_id ?
                <YouTubeEmbed youtube_id={youtube_id} />
                :
                article?.thumbnail_url &&
                // eslint-disable-next-line @next/next/no-img-element
                <img src={article.thumbnail_url} alt={article.title} width={500}
                    className='w-full h-auto md:w-[500px] rounded-xl' />
            }
            <div className="flex flex-col items-center w-full space-y-2">
            <div className="">{article.title}</div>
                <div className="flex flex-row items-center space-x-3">
                    <button onClick={() => setDialogOpen(article.id)} className="bg-cyan-300 text-white p-[5px] rounded-xl text-sm">保存</button>
                    <Link href={article.original_url} className="bg-slate-100 text-slate-500 p-[5px] rounded-xl text-sm">読む</Link>
                </div>
            </div>
        </div>
    )
}
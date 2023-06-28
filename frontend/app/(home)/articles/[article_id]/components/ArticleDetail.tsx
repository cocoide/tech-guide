import YouTubeEmbed from '@/app/(home)/_components/YoutubeEmbed';
import { Article } from '@/app/_models';

export default function ArticleDetail({ article }: { article: Article }) {
    function extractYoutubeID(url: string): string | null {
        const match = url.match(/^(?:https?:\/\/)?(?:www\.)?(?:m\.)?(?:youtu\.be\/|youtube\.com\/(?:watch\?(?:\S*?&)?v=|(?:embed|v|vi|user)\/))([\w-]+)/);
        return match ? match[1] : null;
    }
    const youtube_id = extractYoutubeID(String(article?.original_url))
    return (
        <div className="bg-white rounded-xl shadow-md p-5 ring-1 ring-slate-200 flex flex-col md:flex-row items-center 
        space-y-3 md:space-x-3">
            {youtube_id ?
                <YouTubeEmbed youtube_id={youtube_id} />
                :
                article?.thumbnail_url &&
                // eslint-disable-next-line @next/next/no-img-element
                <img src={article.thumbnail_url} alt={article.title} width={500}
                    className='w-full h-auto md:w-[500px] rounded-xl' />
            }
            <div className="">{article.title}</div>
        </div>
    )
}
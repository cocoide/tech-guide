import YouTubeEmbed from '@/app/(timeline)/_components/YoutubeEmbed';
import { Article } from '@/types/model';
import { extractYoutubeID } from '@/utils/regex';

export default function ArticlePreviewSection({ article }: { article: Article }) {
    const youtubeID = extractYoutubeID(article?.original_url)
    return (
        <>
            <div className="text-2xl text-gray-700 font-bold">{article?.title}</div>
            {youtubeID ?
                <YouTubeEmbed youtube_id={youtubeID} />
                :
                <>
                    {article.thumbnail_url &&
                        // eslint-disable-next-line @next/next/no-img-element
                        <img src={article.thumbnail_url} alt={article?.title} width={200} height={100}
                            className='w-[500px] h-auto rounded-xl ring-1 ring-gray-300' />
                    }
                </>
            }
            {article.summary?.length > 0 &&
                <div className="bg-gray-100 text-gray-400 text-sm p-3 rounded-md">
                    <div className="border-gray-500  border-l-2 pl-2">Sumamry: {article.summary}
                    </div>
                </div>
            }
            <div className="w-full flex flex-wrap gap-3">{article.topics.map((topic) => (
                <div key={topic.name} className="text-gray-400 ring-1 ring-gray-300 p-1 rounded-xl"># {topic.name}</div>
            ))}</div>
        </>
    )
}

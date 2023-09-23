import YouTubeEmbed from '@/app/(home)/_components/YoutubeEmbed';
import { Article } from '@/types/model';
import { extractYoutubeID } from '@/utils/regex';
import { Suspense } from 'react';
import SpeakerDeckEmbed from '../elements/SpeakerDeckEmbed';
import SpeakerDeckLoader from '../loader/SpeakerDeckLoader';

export default async function ArticlePreviewSection({ article }: { article: Article }) {
    const youtubeID = extractYoutubeID(article?.original_url)
    const isSpeakerDeck = article?.source.domain === "speakerdeck.com";
    return (
        <div className="flex flex-col space-y-2 w-full">
            <div className="text-xl text-slate-800 dark:text-slate-100 font-bold">{article?.title}</div>
            <>
                {youtubeID ?
                    <YouTubeEmbed youtube_id={youtubeID} />
                    :
                    isSpeakerDeck ?
                        <Suspense fallback={<SpeakerDeckLoader />}>
                            <SpeakerDeckEmbed url={article.original_url} />
                        </Suspense>
                        :
                        <>
                            {/* {article.thumbnail_url &&
                                // eslint-disable-next-line @next/next/no-img-element
                                <img src={article.thumbnail_url} alt={article?.title} width={200} height={100}
                                    className='w-[500px] h-auto rounded-xl ring-1 ring-gray-300' />
                            } */}
                        </>
                }
            </>
        </div>
    )
}

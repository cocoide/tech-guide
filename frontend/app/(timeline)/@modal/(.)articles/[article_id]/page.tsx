"use client"

import YouTubeEmbed from '@/app/(timeline)/_components/YoutubeEmbed'
import { API_URL } from '@/libs/constant'
import { Article } from '@/types/model'
import { extractYoutubeID } from '@/utils/regex'
import { ArrowTopRightOnSquareIcon, ChatBubbleOvalLeftEllipsisIcon, DocumentIcon, HandThumbDownIcon, HandThumbUpIcon, XMarkIcon } from '@heroicons/react/24/outline'
import { useQuery } from '@tanstack/react-query'
import Link from 'next/link'
import { useRouter } from 'next/navigation'

interface Props extends ArticleParams {
}

export default function Page({ params }: Props) {
    const router = useRouter()
    function handleClose() {
        router.back()
    }
    const getArticleDetail = async (articleId: string) => {
        const url = API_URL + `/article/${articleId}`
        const res = await fetch(url);
        return await res.json();
    };
    const { data: article } = useQuery<Article>({
        queryKey: ['articleDetail'],
        queryFn: () => getArticleDetail(params.article_id)
    })
    const youtube_id = extractYoutubeID(article?.original_url)
    console.log(youtube_id)
    return (
        <>
            <button onClick={handleClose}
                className="z-30 bg-black/40  fixed inset-0 backdrop-blur-[3px] animate-appear"></button>
            <button className="lg:hidden fixed left-0 top-3 py-1 pr-1 pl-5 rounded-r-full bg-black/30 z-50" onClick={handleClose}
            ><XMarkIcon className='h-6 w-6 text-white' /></button>

            <button className="hidden lg:flex fixed left-5 top-5 p-2 rounded-full bg-white z-50" onClick={handleClose}
            ><XMarkIcon className='h-5 w-5 text-gray-800' /></button>

            {/* modal content */}
            {article &&
                <div className="z-40 fixed bg-white inset-0 rounded-md ring-1 ring-gray-200
            sm:top-[100px] lg:left-[100px] lg:right-[100px] 
            overflow-y-scroll divide-x flex flex-col sm:flex-row">
                    <div className="flex flex-col p-5 lg:p-10 w-full space-y-5">
                        <div className="text-2xl text-gray-700 font-bold">{article?.title}</div>
                        {article.thumbnail_url &&
                            <>
                                {youtube_id ?
                                    <YouTubeEmbed youtube_id={youtube_id} />
                                    :
                                    // eslint-disable-next-line @next/next/no-img-element
                                    <img src={article.thumbnail_url} alt={article?.title} width={200} height={100}
                                        className='w-[500px] h-auto rounded-xl ring-1 ring-gray-300' />
                                }
                            </>
                        }
                        <div className="w-full flex flex-wrap gap-3">{article.topics.map((topic) => (
                            <div key={topic.name} className="text-gray-400 ring-1 ring-gray-300 p-1 rounded-xl"># {topic.name}</div>
                        ))}</div>
                        <div className="w-full rounded-xl bg-gray-50 p-2 flex flex-row items-center  space-x-3
                        ring-1 ring-gray-300"
                        ><HandThumbUpIcon className='h-7 w-7 text-gray-500 hover:text-pink-300 hover:bg-pink-100 p-[1px] rounded-full' />
                            <HandThumbDownIcon className='h-7 w-7 text-gray-500' />
                            <ChatBubbleOvalLeftEllipsisIcon className='h-7 w-7 text-gray-500' />
                            <DocumentIcon className='h-7 w-7 text-gray-500' />
                        </div>
                    </div>

                    <div className="flex w-[400px] flex-col p-5 lg:p-10 space-y-5">
                        <Link href={article.original_url} className="ring-1 ring-gray-300 rounded-md p-2 text-gray-400 mr-auto flex flex-row items-center space-x-2"
                        ><ArrowTopRightOnSquareIcon className='h-7 w-7' />
                            <div className=""> 記事を読む</div>
                        </Link>
                        <div className="flex flex-row items-center p-2 ring-1 ring-gray-300 rounded-md space-x-2">
                            {/* eslint-disable-next-line @next/next/no-img-element */}
                            <img src={article.source.icon_url} alt="" className="h-7 w-7 rounded-full " />
                            <div className="text-gray-500">{article.source.name}</div>
                        </div>
                    </div>
                </div>
            }
        </>
    )
}

"use client"
import YouTubeEmbed from '@/app/(timeline)/_components/YoutubeEmbed'
import { Article } from '@/types/model'
import { extractYoutubeID } from '@/utils/regex'
import { ArrowTopRightOnSquareIcon, ChatBubbleOvalLeftEllipsisIcon, ChevronLeftIcon, ChevronRightIcon, DocumentIcon, EllipsisVerticalIcon, HandThumbDownIcon, HandThumbUpIcon } from '@heroicons/react/24/outline'
import Link from 'next/link'
import CloseButton from './CloseButton'
interface Props {
    article?: Article
}
const ModalContent = ({ article }: Props) => {
    const youtubeID = extractYoutubeID(article?.original_url)
    return (
        <div>
            {article &&
                <div className="z-40 fixed bg-white inset-0 rounded-md ring-1 ring-gray-200
            sm:top-[100px] lg:left-[100px] lg:right-[100px] 
            overflow-y-scroll divide-x flex flex-col sm:flex-row">
                    <div className="flex flex-col p-5 lg:p-7 w-full space-y-5 lg:space-y-7">
                        <div className="flex flex-row items-center justify-between text-gray-500">
                            <div className="flex flex-row items-center space-x-5">
                                <ChevronLeftIcon className='w-7 h-7 p-[3px] hover:bg-gray-200 duration-500 rounded-md' />
                                <ChevronRightIcon className='w-7 h-7 p-[3px] hover:bg-gray-200 duration-500 rounded-md' />
                            </div>
                            <div className="flex flex-row items-center space-x-5">
                                <Link href={article.original_url}>
                                    <ArrowTopRightOnSquareIcon className='w-7 h-7 sm:hidden  p-[3px] hover:bg-gray-200 duration-500 rounded-md' />
                                </Link>
                                <EllipsisVerticalIcon className='w-7 h-7 sm:hidden  p-[3px] hover:bg-gray-200 duration-500 rounded-md' />
                                <div className="flex sm:hidden">
                                    <CloseButton />
                                </div>
                            </div>
                        </div>
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
                        <div className="w-full flex flex-wrap gap-3">{article.topics.map((topic) => (
                            <div key={topic.name} className="text-gray-400 ring-1 ring-gray-300 p-1 rounded-xl"># {topic.name}</div>
                        ))}</div>
                        <div className="w-full rounded-xl p-2 flex flex-row items-center  space-x-3
                        ring-1 ring-gray-300"
                        ><HandThumbUpIcon className='h-7 w-7 text-gray-500 hover:text-pink-300 hover:bg-pink-100 p-[1px] rounded-full' />
                            <HandThumbDownIcon className='h-7 w-7 text-gray-500' />
                            <ChatBubbleOvalLeftEllipsisIcon className='h-7 w-7 text-gray-500' />
                            <DocumentIcon className='h-7 w-7 text-gray-500' />
                        </div>
                    </div>

                    <div className="hidden sm:flex w-[400px] flex-col p-5 lg:p-7 space-y-5">
                        <div className="flex flex-row items-center justify-between">
                            <Link href={article.original_url} className="ring-1 ring-gray-300 rounded-md p-2 text-gray-400 mr-auto flex flex-row items-center space-x-2"
                            ><ArrowTopRightOnSquareIcon className='h-7 w-7' />
                                <div className=""> 記事を読む</div>
                            </Link>
                            <CloseButton />
                        </div>
                        <div className="flex flex-row items-center p-2 ring-1 ring-gray-300 rounded-md space-x-2">
                            {/* eslint-disable-next-line @next/next/no-img-element */}
                            <img src={article.source.icon_url} alt="" className="h-7 w-7 rounded-full " />
                            <div className="text-gray-500">{article.source.name}</div>
                        </div>
                    </div>
                </div>
            }
        </div>
    )
}
export default ModalContent
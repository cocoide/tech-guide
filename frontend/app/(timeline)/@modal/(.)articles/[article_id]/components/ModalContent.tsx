"use client"

import { Article } from '@/types/model'
import { ArrowTopRightOnSquareIcon, ChevronLeftIcon, ChevronRightIcon, EllipsisVerticalIcon } from '@heroicons/react/24/outline'
import Link from 'next/link'
import { Suspense } from 'react'
import OutlineLoader from '../loader/OutlineLoader'
import ActionSection from './ActionSection'
import ArticlePreviewSection from './ArticlePreviewSection'
import CloseButton from './CloseButton'
import CommnentSection from './CommentSection'
import OutlineSection from './OutlineSection'

interface Props {
    article?: Article
}
const ModalContent = ({ article }: Props) => {
    const domain = article?.source.domain
    const unShownOutline = domain === "youtube.com" || domain === "speakerdeck.com";
    return (
        <div>
            {article &&
                <div className="z-40 fixed bg-white inset-0 rounded-md ring-1 ring-gray-200
            sm:top-[100px] lg:left-[100px] lg:right-[100px] 
            overflow-y-scroll divide-x flex flex-col sm:flex-row">
                    <div className="flex flex-col p-5 lg:p-7 w-full space-y-3">
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
                        <ArticlePreviewSection article={article} />
                        <ActionSection articleId={article?.id} />
                        <CommnentSection />
                        <div className="w-full flex flex-wrap gap-3">{article.topics.map((topic) => (
                <div key={topic.name} className="text-gray-400 ring-1 ring-gray-300 p-1 rounded-xl"># {topic.name}</div>
            ))}</div>
                        <div className="sm:hidden">
                            {!unShownOutline &&
                            <Suspense fallback={<OutlineLoader />}>
                                <OutlineSection articleURL={article.original_url} />
                            </Suspense>
                            }
                        </div>
                        {article.summary?.length > 0 &&
                            <div className="bg-gray-100 text-gray-400 text-sm p-3 rounded-md">
                                <div className="border-gray-500  border-l-2 pl-2">Sumamry: {article.summary}
                                </div>
                            </div>
                        }
                    </div>

                    <div className="hidden sm:flex w-[400px] lg:w-[500px] flex-col p-5 lg:p-7 space-y-5">
                        <div className="flex flex-row items-center justify-between">
                            <Link href={article.original_url} className="ring-1 ring-gray-300 rounded-md p-2 text-gray-400 mr-auto flex flex-row items-center space-x-2"
                            ><ArrowTopRightOnSquareIcon className='h-7 w-7' />
                                <div className=""> 読む</div>
                            </Link>
                            <CloseButton />
                        </div>
                        <div className="flex flex-row items-center p-2 custom-border rounded-xl space-x-2">
                            {/* eslint-disable-next-line @next/next/no-img-element */}
                            <img src={article.source.icon_url} alt="" className="h-7 w-7 rounded-full " />
                            <div className="text-gray-500">{article.source.name}</div>
                        </div>
                        {!unShownOutline &&
                            <Suspense fallback={<OutlineLoader />}>
                                <OutlineSection articleURL={article.original_url} />
                            </Suspense>
                        }
                    </div>
                </div>
            }
        </div>
    )
}
export default ModalContent
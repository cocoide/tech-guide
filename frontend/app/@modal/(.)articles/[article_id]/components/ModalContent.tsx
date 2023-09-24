"use client"

import { articleAPI } from '@/app/_functions/article'
import { useAuth } from '@/hooks/useAuth'
import { Article } from '@/types/model'
import { StarIcon, XMarkIcon } from '@heroicons/react/24/outline'
import Image from 'next/image'
import { useRouter } from 'next/navigation'
import { Suspense } from 'react'
import OutlineLoader from '../loader/OutlineLoader'
import ArticlePreviewSection from './ArticlePreviewSection'
import OutlineSection from './OutlineSection'

interface Props {
    article?: Article
}
const ModalContent = ({ article }: Props) => {
    const { token } = useAuth()
    const router = useRouter()
    const domain = article?.source.domain
    const unShownOutline = domain === "youtube.com" || domain === "speakerdeck.com"||domain === "github.com"||domain==="producthunt.com/";
    async function handleOnRead(article_id: number) {
        if (token) {
            const { ok } = await articleAPI.ReadArticle(article_id, token)
        }
    }
    return (
        <div>
            {article &&
                <div className="z-30 bg-white dark:bg-black text-gray-500 dark:text-white
                fixed inset-0 rounded-xl
                custom-border
                md:left-[50px] md:right-[50px]
            sm:top-[100px] lg:left-[150px] lg:right-[150px] 
            overflow-y-scroll divide-x flex flex-col sm:flex-row">
                    <div className="flex flex-col p-5 lg:p-7 w-full space-y-3 relative md:px-[50px] lg:px-[100px]">
                        <button onClick={() => router.back()} className='absolute top-5 left-0 p-[5px] bg-black/30 dark:bg-white/50 rounded-r-full'>
                            <XMarkIcon className='h-7 w-7 text-white dark:text-slate-800 ml-3' />
                        </button>
                        <div className="flex flex-row items-center justify-end space-x-3">
                            <StarIcon className="h-7 w-7 text-slate-600 dark:text-slate-300" />
                            <a href={`/sources/${article.source.id}`} target="_blank" >
                                <Image src={article.source.icon_url} alt={article.source.name} width={200} height={200} className='h-7 w-7 rounded-full' />
                            </a>

                            <a href={article.original_url} target="_blank" onClick={() => handleOnRead(article.id)}
                                className="bg-slate-400/60 backdrop-blur-[5px] cutom-outline
                        text-white p-[3px] rounded-xl shadow-sm">元記事を読む</a>


                        </div>

                        <ArticlePreviewSection article={article} />
                        {article.summary?.length > 0 &&
                            <div className="bg-gray-50 text-slate-600 dark:text-slate-300 dark:bg-slate-800
                             text-sm p-3 rounded-md">
                                {article.summary}
                            </div>
                        }
                        {/* <ActionSection articleId={article?.id} /> */}
                        {/* <CommnentSection articleID={article?.id} /> */}
                        <div className="w-full flex flex-wrap gap-3">{article.topics.map((topic) => (
                            <a href={`/topics/${topic.id}`} target="_blank" key={topic.name} className="text-gray-400 ring-1 ring-gray-300 p-1 rounded-xl"># {topic.name}</a>
                        ))}</div>
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
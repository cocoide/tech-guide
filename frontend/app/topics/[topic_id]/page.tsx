
import ArticlesByTopicIDContainer from '@/app/(home)/_containers/ArticlesByTopicIDContainer'
import SectionHeader from '@/app/_components/layouts/desktop/SectionHeader'
import ArticlesContainerLoader from '@/app/_components/loaders/ArticlesContainerLoader'
import { topicAPI } from '@/app/_functions/topic'
import { Suspense } from 'react'
import TopicSectionHeader from './TopicSectionHeader'

interface Props {
    params: {
        topic_id: string
    }
}
export default async function TopicPage({ params }: Props) {
    const topicID = Number(params.topic_id)
    const { data: topic } = await topicAPI.GetTopicData(topicID)
    return (
        <>
            <div className="flex flex-col w-full dark:bg-black">
                <div className="sticky top-0 h-10 bg-white/70 dark:bg-black/30 dark:text-slate-300
             backdrop-blur-[5px] z-20">
                    {topic ?
                        <TopicSectionHeader topic={topic} />
                        :
                        <SectionHeader titleItem={<div className="h-6 w-14 rounded-full animate-pulse bg-slate-100 dark:bg-slate-800
                "></div>} rightItem={<div></div>} />
                    }
                </div>
            </div>
            <Suspense fallback={<ArticlesContainerLoader />}>
                <ArticlesByTopicIDContainer topicID={topicID} />
            </Suspense>
        </>
    )
}
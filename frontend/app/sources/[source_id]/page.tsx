import ArticlesBySourceIDContainer from '@/app/(home)/_containers/ArticlesBySourceIDContainer'
import SectionHeader from '@/app/_components/layouts/desktop/SectionHeader'
import ArticlesContainerLoader from '@/app/_components/loaders/ArticlesContainerLoader'
import { sourceAPI } from '@/app/_functions/source'
import { Suspense } from 'react'
import SourceSectionHeader from './SectionHeader'

interface Props {
    params: {
        source_id: string
    }
}
export default async function SourcePage({ params }: Props) {
    const sourceID = Number(params.source_id)
    const { data: source } = await sourceAPI.GetSourceData(sourceID)
    return (
        <>
            <div className="flex flex-col w-full pb-10 dark:bg-black">
                <div className="sticky top-0 h-10 bg-white/70 dark:bg-black/30 dark:text-slate-300
             backdrop-blur-[5px] z-20">
                    {source ?
                        <SourceSectionHeader source={source} />
                        :
                        <SectionHeader titleItem={<div className="h-6 w-14 rounded-full animate-pulse bg-slate-100 dark:bg-slate-800
                "></div>} rightItem={<div></div>} />
                    }
                </div>
            </div>
            <Suspense fallback={<ArticlesContainerLoader />}>
                <ArticlesBySourceIDContainer sourceID={sourceID} />
            </Suspense>
        </>
    )
}

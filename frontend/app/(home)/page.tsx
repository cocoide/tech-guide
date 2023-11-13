import DiscussContainer from '@/features/discuss/DiscussContainer'
import { Suspense } from 'react'
import ToggleDarkModeButton from '../_components/layouts/button/ToggleDarkModeButton'
import SectionHeader from '../_components/layouts/desktop/SectionHeader'
import ArticlesContainerLoader from '../_components/loaders/ArticlesContainerLoader'
import SectionHeaderButtonGroup from './_components/SectionHeaderButtonGroup'
import LatestArticleContainer from './_containers/LatestArticleContainer'
import TrendArticleContainer from './_containers/TrendArticleContainer'
interface Props {
    searchParams: { "order": "trend" | "discuss" }
}
export default async function ArticlePage({ searchParams }: Props) {
    const order = searchParams.order
    return (
        <div className="flex flex-col w-full pb-10 dark:bg-black">
            <div className="sticky top-0 h-10 bg-white/70 dark:bg-black/30 dark:text-slate-300
             backdrop-blur-[5px] z-20">
                <SectionHeader
                    titleItem={<SectionHeaderButtonGroup />}
                    rightItem={<ToggleDarkModeButton />} />
            </div>
            {order === "trend" ?
                <Suspense fallback={<ArticlesContainerLoader />}>
                    <TrendArticleContainer />
                </Suspense>
                :
                order === "discuss" ?
                    <Suspense fallback={<ArticlesContainerLoader />}>
                        <DiscussContainer />
                    </Suspense>
                    :
                <LatestArticleContainer />
            }
        </div>
    )
}
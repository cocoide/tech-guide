import { NewspaperIcon } from '@heroicons/react/24/outline'
import TopicDialogButton from '../_components/layouts/button/FeedFileterDialogButton'
import SectionHeader from '../_components/layouts/desktop/SectionHeader'

interface Props {
    children: React.ReactNode
}
export default function FeedLayout({ children }: Props) {
    return (

        <div className="flex flex-col w-full pb-10 relative">
            <div className="sticky top-0 h-12 bg-white/70 dark:bg-black/30 dark:text-slate-300 backdrop-blur-[5px] z-20">
                <SectionHeader
                    titleItem={<div className='custom-badge text-gray-500'><NewspaperIcon className='h-5 w-5' /><div>フィード</div></div>}
                    rightItem={
                        <TopicDialogButton />
                    }
                />
            </div>
                {children}
        </div>
    )
}
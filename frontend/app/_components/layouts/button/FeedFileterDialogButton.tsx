"use client"

import { topicDialogAtom } from '@/stores/dialog'
import { AdjustmentsHorizontalIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'

const FeedFileterDialogButton = () => {
    const [__, setTopicDialogOpen] = useAtom(topicDialogAtom)
    return (
        <button onClick={() => setTopicDialogOpen(true)} className="custom-text custom-badge">
            <AdjustmentsHorizontalIcon className='h-7 w-7' />
            <div className="text-sm">設定</div>
        </button>
    )
}
export default FeedFileterDialogButton
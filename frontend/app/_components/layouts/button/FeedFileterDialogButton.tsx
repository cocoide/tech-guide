"use client"

import { topicDialogAtom } from '@/stores/dialog'
import { AdjustmentsHorizontalIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'

const FeedFileterDialogButton = () => {
    const [__, setTopicDialogOpen] = useAtom(topicDialogAtom)
    return (
        <button onClick={() => setTopicDialogOpen(true)} className="custom-text custom-badge custom-border p-[3px] rounded-md shadow-sm">
            <AdjustmentsHorizontalIcon className='h-5 w-5' />
            <div className="text-sm">設定</div>
        </button>
    )
}
export default FeedFileterDialogButton
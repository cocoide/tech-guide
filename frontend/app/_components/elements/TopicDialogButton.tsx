"use client"

import { topicDialogAtom } from '@/stores/dialog'
import { AdjustmentsHorizontalIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'

const TopicDialogButton = () => {
    const [__, setTopicDialogOpen] = useAtom(topicDialogAtom)
    return (
        <button onClick={() => setTopicDialogOpen(true)} className="text-gray-500">
            <AdjustmentsHorizontalIcon className='h-7 w-7' />
        </button>
    )
}
export default TopicDialogButton
"use client"

import WStack from '@/app/_components/elements/ui/WStack'
import { clsx } from '@/utils/clsx'
import { ChatBubbleOvalLeftEllipsisIcon, FireIcon, HomeIcon } from '@heroicons/react/24/outline'
import { useRouter, useSearchParams } from 'next/navigation'

const SectionHeaderButtonGroup = () => {
    const searchParams = useSearchParams()
    const order = searchParams.get("order")
    const router = useRouter()
    return (
        <WStack centerX={true} className='space-x-3 w-full'>
            <button onClick={() => router.push('/')} className={clsx('custom-badge custom-text custom-border p-1 rounded-xl', !order ? 'bg-gray-100 dark:bg-gray-800' : '')}>
                <HomeIcon className='h-5 w-5' /><div className='text-sm'>Latest</div>
            </button>
            <button onClick={() => router.push('/?order=trend')} className={clsx('custom-badge custom-text custom-border p-1 rounded-xl', order === "trend" ? "bg-gray-100 dark:bg-gray-800" : "")}>
                <FireIcon className='h-5 w-5' /><div className='text-sm'>Trend</div>
            </button>
            <button onClick={() => router.push('/?order=discuss')} className={clsx('custom-badge custom-text custom-border p-1 rounded-xl', order === "discuss" ? "bg-gray-100 dark:bg-gray-800" : "")}>
                <ChatBubbleOvalLeftEllipsisIcon className='h-5 w-5' /><div className='text-sm'>Discuss</div>
            </button>
        </WStack>
    )
}
export default SectionHeaderButtonGroup
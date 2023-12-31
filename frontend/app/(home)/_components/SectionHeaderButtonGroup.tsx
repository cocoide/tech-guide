"use client"

import WStack from '@/app/_components/elements/ui/WStack'
import { useAuth } from '@/hooks/useAuth'
import { clsx } from '@/utils/clsx'
import { ChatBubbleOvalLeftEllipsisIcon, FireIcon, HomeIcon, NewspaperIcon } from '@heroicons/react/24/outline'
import { useRouter, useSearchParams } from 'next/navigation'

const SectionHeaderButtonGroup = () => {
    const searchParams = useSearchParams()
    const { status } = useAuth()
    const order = searchParams.get("order")
    const router = useRouter()
    function handleLatest() {
        router.push('/')
    }
    function handleTrend() {
        router.push('/?order=trend')
    }
    function handleDiscuss() {
        router.push('/?order=discuss')
    }
    function handleCustom() {
        router.push(`/feed`)
    }
    return (
        <WStack centerX={true} className='space-x-3 w-full'>
            <button onClick={handleLatest} className={clsx('custom-badge custom-text custom-border p-1 rounded-xl', !order ? 'bg-gray-100 dark:bg-gray-800' : '')}>
                <HomeIcon className='h-5 w-5' /><div className='text-sm'>Latest</div>
            </button>
            {status === "authenticated" &&
                <button onClick={handleCustom} className={clsx('custom-badge custom-text custom-border p-1 rounded-xl', order === "feed" ? "bg-gray-100 dark:bg-gray-800" : "")}>
                    <NewspaperIcon className='h-5 w-5' /><div className='text-sm'>Custom</div>
                </button>
            }
            <button onClick={handleTrend} className={clsx('custom-badge custom-text custom-border p-1 rounded-xl', order === "trend" ? "bg-gray-100 dark:bg-gray-800" : "")}>
                <FireIcon className='h-5 w-5' /><div className='text-sm'>Trend</div>
            </button>
            <button onClick={handleDiscuss} className={clsx('custom-badge custom-text custom-border p-1 rounded-xl', order === "discuss" ? "bg-gray-100 dark:bg-gray-800" : "")}>
                <ChatBubbleOvalLeftEllipsisIcon className='h-5 w-5' /><div className='text-sm'>Discuss</div>
            </button>
        </WStack>
    )
}
export default SectionHeaderButtonGroup
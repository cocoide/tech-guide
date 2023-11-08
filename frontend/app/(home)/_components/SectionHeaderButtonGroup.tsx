"use client"

import WStack from '@/app/_components/elements/ui/WStack'
import { ChatBubbleOvalLeftEllipsisIcon, FireIcon, HomeIcon } from '@heroicons/react/24/outline'
import Link from 'next/link'

const SectionHeaderButtonGroup = () => {
    return (
        <WStack centerX={true} className='space-x-3'>
            <Link href={'/'} className='custom-badge custom-text custom-border p-1 rounded-xl'>
                <HomeIcon className='h-5 w-5' /><div className='text-sm'>Latest</div>
            </Link>
            <Link href={'/?order=trend'} className='custom-badge custom-text custom-border p-1 rounded-xl'>
                <FireIcon className='h-5 w-5' /><div className='text-sm'>Trend</div>
            </Link>
            <Link href={'/?order=discuss'} className='custom-badge custom-text custom-border p-1 rounded-xl'>
                <ChatBubbleOvalLeftEllipsisIcon className='h-5 w-5' /><div className='text-sm'>Discuss</div>
            </Link>
        </WStack>
    )
}
export default SectionHeaderButtonGroup
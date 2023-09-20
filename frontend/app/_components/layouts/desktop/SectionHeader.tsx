"use client"
import { ChevronLeftIcon } from '@heroicons/react/24/outline'
import { useRouter } from 'next/navigation'
import { ReactNode } from 'react'
import ScrollButton from '../../elements/ScrollButton'

interface Props {
    titleItem: ReactNode
    rightItem: ReactNode
}
export default function SectionHeader({ titleItem, rightItem }: Props) {
    const router = useRouter()
    return (
        <div className="flex flex-row justify-between items-center p-2 w-full custom-border">
            <div className="flex flex-row items-center text-gray-500 space-x-1">
                <button onClick={() => router.back()} className="">
                    <ChevronLeftIcon className='h-5 w-5' />
                </button>
                <ScrollButton height={0}>{titleItem}</ScrollButton>
            </div>
            {rightItem}
        </div>
    )
}
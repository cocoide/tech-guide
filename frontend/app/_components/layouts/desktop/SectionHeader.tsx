"use client"
import { ChevronLeftIcon } from '@heroicons/react/24/outline'
import { useRouter } from 'next/navigation'
import { ReactNode } from 'react'

interface Props {
    title: string
    rightItem: ReactNode
}
export default function SectionHeader({ title, rightItem }: Props) {
    const router = useRouter()
    return (
        <div className="flex flex-row justify-between items-center p-2 w-full border-y-[0.5px]">
            <div className="flex flex-row items-center text-gray-500 space-x-1">
                <button onClick={() => router.back()} className="">
                    <ChevronLeftIcon className='h-5 w-5' />
                </button>
                <div className="">{title}</div>
            </div>
            {rightItem}
        </div>
    )
}
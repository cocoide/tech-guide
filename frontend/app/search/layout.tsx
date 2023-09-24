"use client"
import { MagnifyingGlassIcon } from '@heroicons/react/24/outline'
import { useSearchParams } from 'next/navigation'
import ToggleDarkModeButton from '../_components/layouts/button/ToggleDarkModeButton'
import SectionHeader from '../_components/layouts/desktop/SectionHeader'

interface Props {
    children: React.ReactNode
}

export default function SearchLayout({ children }: Props) {
    const searchParams = useSearchParams()
    const query = searchParams.get("q")
    return (
        <div className='flex flex-col w-full pb-10 dark:bg-black'>
            <div className="sticky top-0 h-10 bg-white/70 dark:bg-black/30 dark:text-slate-300
         backdrop-blur-[5px] z-20">
                <SectionHeader
                    titleItem={<div className='custom-badge text-gray-500 dark:text-white'>
                        <MagnifyingGlassIcon className='h-5 w-5' /><div>{query}: の検索結果</div>
                    </div>}
                    rightItem={<ToggleDarkModeButton />} />
            </div>
            {children}
        </div>
    )
}
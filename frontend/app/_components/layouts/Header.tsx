'use client'

import { loginDialogAtom } from '@/stores/dialog'
import { BellIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import { useSession } from 'next-auth/react'
import Image from 'next/image'
import Link from 'next/link'

export function Header() {
    const { data: session, status } = useSession()
    const [_, setLoginDialogOpen] = useAtom(loginDialogAtom)
    return (
        <div className="w-full p-[12px] bg-white/70  dark:bg-black
        backdrop-blur-[3px] flex flex-row items-center justify-between min-h-19 border-b-[0.5px] custom-border-color">
            <div className="flex items-center space-3">
                <Image src="/logo.svg" alt="" width={100} height={100} className='h-7 w-7' />
            <Link href={'/'} className="text-xl font-bold dark:text-white">Tech Guide</Link>

            </div>
            {status != "loading" && !session?.user &&
                <button onClick={() => setLoginDialogOpen(true)}
                    className="text-cyan-300 ring-1 bg-white dark:bg-black
                     ring-gray-200 dark:ring-gray-500 
                p-1 rounded-full font-medium
                flex flex-row items-center text-sm
                animate-appear duration-500"
                ><UserCircleIcon className='h-5 w-5' />ログイン</button>
            }
            {status != "loading" && session?.user &&
                <BellIcon className='h-7 w-7 text-gray-500' />
            }
        </div>
    )
}
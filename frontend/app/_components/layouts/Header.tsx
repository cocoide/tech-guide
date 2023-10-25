'use client'

import { useAuth } from "@/hooks/useAuth"
import { loginDialogAtom } from '@/stores/dialog'
import { BellIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Image from 'next/image'
import Link from 'next/link'

export function Header() {
    const { status } = useAuth()
    const [_, setLoginDialogOpen] = useAtom(loginDialogAtom)
    return (
        <div className="w-full p-[12px] bg-white/70  dark:bg-black
        backdrop-blur-[3px] flex flex-row items-center justify-between min-h-19 border-b-[0.5px] custom-border-color">
            <div className="flex items-center space-x-1">
                <Image src="/logo.svg" alt="" width={100} height={100} className='h-7 w-7' />
            <Link href={'/'} className="text-xl font-bold dark:text-white">Tech Guide</Link>

            </div>
            {status === "unauthenticated" &&
                <button onClick={() => setLoginDialogOpen(true)}
                    className="text-cyan-300 ring-1 bg-white dark:bg-black
                     ring-gray-200 dark:ring-gray-500 
                p-1 rounded-full font-medium
                flex flex-row items-center text-sm
                animate-appear duration-500"
                ><UserCircleIcon className='h-5 w-5' />ログイン</button>
            }
            {status == "loading"&&
                <div className="h-7 w-7 custom-loader rounded-full"></div>
            }
            {status === "authenticated" &&
                <BellIcon className='h-7 w-7 text-gray-500' />
            }
        </div>
    )
}
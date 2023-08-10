'use client'

import { loginDialogAtom } from '@/stores/dialog'
import { BellIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import { useSession } from 'next-auth/react'
import Link from 'next/link'

export function Header() {
    const { data: session, status } = useSession()
    const [_, setLoginDialogOpen] = useAtom(loginDialogAtom)
    return (
        <div className="w-full p-[12px] bg-white/70 backdrop-blur-[3px] flex flex-row items-center justify-between min-h-19">
            <Link href={'/'} className="text-xl font-bold">Tech Guide</Link>
            {status != "loading" && !session?.user &&
                <button onClick={() => setLoginDialogOpen(true)}
                    className="text-cyan-300 ring-1 bg-white ring-gray-200 
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
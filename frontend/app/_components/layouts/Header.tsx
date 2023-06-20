'use client'

import { postDialogAtom } from '@/stores/dialog'
import { DocumentTextIcon, UserCircleIcon, UserMinusIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import { signOut, useSession } from 'next-auth/react'
import Link from 'next/link'

export function Header() {
    const { data: session, status } = useSession()
    const [_, setOpen] = useAtom(postDialogAtom)
    return (
        <div className="w-full p-3 bg-white/70 backdrop-blur-[3px] flex flex-row items-center justify-between">
            <div className="text-xl font-bold">Tech Guide</div>
            {status != "loading" && !session?.user &&
                <Link href={"/login"}
                    className="text-cyan-300 ring-1 bg-white ring-gray-200 
                p-1 rounded-full font-medium
                flex flex-row items-center text-sm
                animate-appear duration-500"
                ><UserCircleIcon className='h-5 w-5' />ログイン</Link>
            }
            {status != "loading" && session?.user &&
                <div className='flex items-center space-x-3'>
                    <button className="text-cyan-300 ring-1 bg-white ring-gray-200
                p-1 rounded-full font-medium
                flex flex-row items-center text-sm
                animate-appear duration-500"
                    onClick={() => signOut()}><UserMinusIcon className='h-5 w-5' />ログアウト</button>
                    <button className="hidden text-cyan-300 ring-1 bg-white ring-gray-200
                p-1 rounded-full font-medium
                sm:flex flex-row items-center text-sm
                animate-appear duration-500"
                        onClick={() => setOpen(true)}><DocumentTextIcon className='h-5 w-5' />投稿する</button>
                </div>
            }
        </div>
    )
}
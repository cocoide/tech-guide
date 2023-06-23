'use client'

import { loginDialogAtom, postDialogAtom } from '@/stores/dialog'
import { DocumentTextIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import { useSession } from 'next-auth/react'
import Link from 'next/link'
import UserDropDown from './UserDropDown'

export function Header() {
    const { data: session, status } = useSession()
    const [_, setPostDialogOpen] = useAtom(postDialogAtom)
    const [__, setLoginDialogOpen] = useAtom(loginDialogAtom)
    return (
        <div className="w-full p-4 bg-white/70 backdrop-blur-[3px] flex flex-row items-center justify-between min-h-19">
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
                <div className='flex items-center space-x-3'>
                    <button className=" text-cyan-300 ring-1 bg-white ring-gray-200
                p-1 rounded-full font-medium
                flex flex-row items-center text-sm
                animate-appear duration-500"
                        onClick={() => setPostDialogOpen(true)}><DocumentTextIcon className='h-5 w-5' />投稿する</button>
                    <UserDropDown
                        name={session.user.name}
                        uid={session.user.uid}
                        img={session.user.image}
                    />
                </div>
            }
        </div>
    )
}
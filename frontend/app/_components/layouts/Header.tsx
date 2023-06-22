'use client'

import { loginDialogAtom, postDialogAtom } from '@/stores/dialog'
import { UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import { useSession } from 'next-auth/react'
import UserDropDown from './UserDropDown'

export function Header() {
    const { data: session, status } = useSession()
    const [_, setPostDialogOpen] = useAtom(postDialogAtom)
    const [__, setLoginDialogOpen] = useAtom(loginDialogAtom)
    return (
        <div className="w-full p-2 bg-white/70 backdrop-blur-[3px] flex flex-row items-center justify-between min-h-19">
            <div className="text-xl font-bold">Tech Guide</div>
            {status != "loading" && !session?.user &&
                <button onClick={() => setLoginDialogOpen(true)}
                    className="text-cyan-300 ring-1 bg-white ring-gray-200 
                p-1 rounded-full font-medium
                flex flex-row items-center text-sm
                animate-appear duration-500"
                ><UserCircleIcon className='h-5 w-5' />ログイン</button>
            }
            {status != "loading" && session?.user &&
                <UserDropDown
                    uid={session?.user.uid}
                    img={session?.user.image}
                />
            }
            {/* {status != "loading" && session?.user && */}
            {/* <div className='flex items-center space-x-3'> */}
            {/* <button className="text-cyan-300 ring-1 bg-white ring-gray-200
                p-1 rounded-full font-medium
                flex flex-row items-center text-sm
                animate-appear duration-500"
                    onClick={() => signOut()}><UserMinusIcon className='h-5 w-5' />ログアウト</button> */}
            {/* <button className="hidden text-cyan-300 ring-1 bg-white ring-gray-200
                p-1 rounded-full font-medium
                sm:flex flex-row items-center text-sm
                animate-appear duration-500"
                        onClick={() => setPostDialogOpen(true)}><DocumentTextIcon className='h-5 w-5' />投稿する</button>
                </div>
            } */}
        </div>
    )
}
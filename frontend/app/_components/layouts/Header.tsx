'use client'

import { useAuth } from "@/hooks/useAuth"
import { useSession } from '@/hooks/useSession'
import { loginDialogAtom, postDialogAtom } from '@/stores/dialog'
import { MagnifyingGlassIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Image from 'next/image'
import Link from 'next/link'
import WStack from '../elements/ui/WStack'

export function Header() {
    const { status } = useAuth()
    const session = useSession()
    const [_, setIsPostDialogOpen] = useAtom(postDialogAtom)
    const [__, setIsLoginDialogOpen] = useAtom(loginDialogAtom)
    function handleAuth() {
        if (status === "unauthenticated") {
            setIsLoginDialogOpen(true)
        }
    }
    function handlePost() {
        if (status === "unauthenticated") {
            setIsLoginDialogOpen(true)
        }
        if (status === "authenticated") {
            setIsPostDialogOpen(true)
        }
    }
    const AccountURL = session.account_id ? `/accounts/${session.account_id}` : "/"
    const [___, setLoginDialogOpen] = useAtom(loginDialogAtom)
    return (
        <div className="w-full p-[12px] bg-white/70  dark:bg-black
        backdrop-blur-[3px] flex flex-row items-center justify-between min-h-19 border-b-[0.5px] custom-border-color custom-text">
            <div className="flex items-center space-x-1">
                <Image src="/logo.svg" alt="" width={100} height={100} className='h-7 w-7' />
            <Link href={'/'} className="text-xl font-bold dark:text-white">Tech Guide</Link>
            </div>
            <WStack centerX={true} className='space-x-3'>
                <Link href={"/explore"} ><MagnifyingGlassIcon className='h-7 w-7' /></Link>
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
                <Link href={AccountURL} onClick={handleAuth}>
                    {
                        session.avatar_url.length > 0 ?
                            <Image src={session.avatar_url} alt={session.display_name} width={200} height={200}
                                    className="h-10 w-10 rounded-full custom-border"></Image>
                            :
                                <UserCircleIcon className='h-10 w-10' />}
                </Link>
            }
            {status === "authenticated" &&
                    <button onClick={handlePost} className='p-[5px] bg-cyan-300 rounded-xl shadow-sm text-white dark:text-black text-sm'>投稿</button>
            }
            </WStack>
        </div>
    )
}
'use client'

import { UserCircleIcon } from '@heroicons/react/24/outline'
import { useSession } from 'next-auth/react'
import Link from 'next/link'

export function Header() {
    const { data: session, status } = useSession()

    return (
        <div className="w-full p-3 bg-white/70 backdrop-blur-[3px] flex flex-row items-center justify-between">
            <div className="text-xl font-bold">Tech Guide</div>
            {status != "loading" && !session?.user &&
                <Link href={"/login"}
                    className="text-cyan-300 ring-1 bg-white ring-gray-200 
                p-1 rounded-full font-medium
                flex flex-row items-center text-sm"
                ><UserCircleIcon className='h-5 w-5' />ログイン</Link>
            }
        </div>
    )
}
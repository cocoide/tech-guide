"use client"
import { useAuth } from '@/hooks/useAuth'
import { loginDialogAtom, postDialogAtom } from '@/stores/dialog'
import { BookmarkIcon, HomeIcon, MagnifyingGlassIcon, NewspaperIcon, PlusCircleIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Image from 'next/image'
import Link from 'next/link'


const LeftSideVar = () => {
    const [_, openPostDialog] = useAtom(postDialogAtom)
    const [__, openLoginDialog] = useAtom(loginDialogAtom)
    const { status, user } = useAuth()
    const SideVarList = [
        { label: "ホーム", href: "/", icon: <HomeIcon className='h-7 w-7' />, isLogin: false },
        { label: "フィード", href: "/feed", icon: <NewspaperIcon className='h-7 w-7' />, isLogin: true },
        { label: "保存リスト", href: `/accounts/${user.uid}/collections`, icon: <BookmarkIcon className='h-7 w-7' />, isLogin: true },
        { label: "検索", href: "/explore", icon: <MagnifyingGlassIcon className='h-7 w-7' />, isLogin: false },
    ]
    return (
        <div className="hidden md:flex flex-col p-5 w-[190px] lg:w-[200px] justify-between  items-center h-[100%]">
            <div className="flex flex-row items-center space-x-3">
                <Image src="/logo.svg" alt="" width={100} height={100} className='h-7 w-7 lg:h-10 lg:w-10' />
            <Link href={"/"} className="text-2xl lg:text-3xl font-bold text-slate-700 dark:text-white w-full"
            >TechGuide</Link>
            </div>
            <div className="flex flex-col space-y-5 text-slate-400 mt-10 w-full">{
                SideVarList.map((item => (
                    (!item.isLogin || (item.isLogin && status === 'authenticated')) &&
                    <Link href={item.href} key={item.label} className="flex flex-row items-center space-x-2 animate-appear">
                        {item.icon}
                        <div className="">{item.label}</div>
                        </Link>
                )))
            }</div>
            <div className="flex flex-col space-y-3 w-full mt-auto">
                {status === 'authenticated' ?
                    <button onClick={() => openPostDialog(true)}
                        className="animate-appear bg-cyan-300  text-white py-2 w-[150px] rounded-xl shadow-sm flex flex-row items-center justify-center space-x-1"
                ><PlusCircleIcon className='h-5 w-5' />
                    <div className="">投稿する</div>
                </button>
                    :
                    <button onClick={() => openLoginDialog(true)}
                        className="animate-appear bg-white dark:bg-inherit
                          text-cyan-300 ring-1 ring-gray-200 dark:ring-gray-500 py-2 w-[150px] rounded-xl shadow-sm flex flex-row items-center justify-center space-x-1"
                    ><UserCircleIcon className='h-5 w-5' />
                        <div className="">ログイン</div>
                    </button>
                }
                {user.image &&
                    <Link href={`/accounts/${user.uid}`} className="animate-appear flex flex-row space-x-3 ring-[0.5px] w-[150px] items-center p-1 rounded-md ring-gray-200">
                        <Image src={user.image} alt={''} width={100} height={100} className='h-7 w-7 rounded-full' />
                        <div className="">{user.name}</div>
                    </Link>
                }
            </div>
        </div>
    )
}
export default LeftSideVar
"use client"
import { useAuth } from '@/hooks/useAuth'
import { loginDialogAtom, postDialogAtom } from '@/stores/dialog'
import { DocumentDuplicateIcon, FireIcon, HomeIcon, PlusCircleIcon, RssIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Image from 'next/image'
import Link from 'next/link'

const SideVarList = [
    { label: "ホーム", href: "/", icon: <HomeIcon className='h-7 w-7' />, isLogin: false },
    { label: "トレンド", href: "/trend", icon: <FireIcon className='h-7 w-7' />, isLogin: false },
    { label: "フィード", href: "/trend", icon: <RssIcon className='h-7 w-7' />, isLogin: true },
    { label: "保存リスト", href: "", icon: <DocumentDuplicateIcon className='h-7 w-7' />, isLogin: true },
]

const LeftSideVar = () => {
    const [_, openPostDialog] = useAtom(postDialogAtom)
    const [__, openLoginDialog] = useAtom(loginDialogAtom)
    const { status, user } = useAuth()
    function handleOpenPostDIalog() {
        if (status == "authenticated") {
            openPostDialog(true)
        }
        if (status == "unauthenticated") {
            openLoginDialog(true)
        }
    }
    return (
        <div className="hidden md:flex flex-col p-5 w-[220px] lg:w-[250px] justify-between  items-center h-[100%]">
            <Link href={"/"} className="text-3xl font-bold text-slate-700 w-full"
            >TechGuide</Link>
            <div className="flex flex-col space-y-5 text-slate-400 mt-10 w-full">{
                SideVarList.map((item => (
                    (!item.isLogin || (item.isLogin && status === 'authenticated')) &&
                    <div key={item.label} className="flex flex-row items-center space-x-2 animate-appear">
                        {item.icon}
                        <div className="">{item.label}</div>
                    </div>
                )))
            }</div>
            <div className="flex flex-col space-y-3 w-full mt-auto">
                <button onClick={handleOpenPostDIalog}
                    className="bg-cyan-300  text-white py-2 w-[150px] rounded-xl shadow-sm flex flex-row items-center justify-center space-x-1"
                ><PlusCircleIcon className='h-5 w-5' />
                    <div className="">投稿する</div>
                </button>
                {user.image &&
                    <Link href={`accounts/${user.uid}`} className="flex flex-row space-x-3 ring-[0.5px] w-[150px] items-center p-1 rounded-md ring-gray-200">
                        <Image src={user.image} alt={''} width={100} height={100} className='h-7 w-7 rounded-full' />
                        <div className="">{user.name}</div>
                    </Link>
                }
            </div>
        </div>
    )
}
export default LeftSideVar
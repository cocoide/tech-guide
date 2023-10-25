'use client'
import { useAuth } from '@/hooks/useAuth'
import { useSession } from '@/hooks/useSession'
import { loginDialogAtom, postDialogAtom } from '@/stores/dialog'
import { HomeIcon, MagnifyingGlassIcon, NewspaperIcon, PlusCircleIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/navigation'

export default function BottomNavigation() {
  const { status } = useAuth()
  const router = useRouter()
  const [_, setIsPostDialogOpen] = useAtom(postDialogAtom)
  const [__, setIsLoginDialogOpen] = useAtom(loginDialogAtom)
  const session = useSession()
  const AccountURL = session.account_id ? `/accounts/${session.account_id}` : "/"
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
  function handleOnFeedPage() {
    if (status === "unauthenticated") {
      setIsLoginDialogOpen(true)
    }
    if (status === "authenticated") {
      router.push("/feed")
    }
  }
  return (
    <div className="w-full flex flex-row items-center justify-between px-[10%] py-2 
    bg-white/80 dark:bg-black/30 backdrop-blur-[5px] text-slate-500 dark:text-slate-300">
      <Link href={"/"} ><HomeIcon className='h-7 w-7' /></Link>
      <button onClick={handleOnFeedPage} ><NewspaperIcon className='h-7 w-7' onClick={handleAuth} /></button>
      <button onClick={handlePost}><PlusCircleIcon className='h-7 w-7' /></button>
      <Link href={"/explore"} ><MagnifyingGlassIcon className='h-7 w-7' /></Link>
      <Link href={AccountURL} onClick={handleAuth}>
        {
          session.avatar_url.length > 0 ?
            <Image src={session.avatar_url} alt={session.display_name} width={200} height={200}
              className="h-7 w-7 rounded-full"></Image>
            :
            <UserCircleIcon className='h-7 w-7' />}</Link>
    </div>
  )
}
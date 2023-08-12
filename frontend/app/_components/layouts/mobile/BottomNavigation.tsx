'use client'
import { useAuth } from '@/hooks/useAuth'
import { loginDialogAtom, postDialogAtom } from '@/stores/dialog'
import { HomeIcon, MagnifyingGlassIcon, NewspaperIcon, PlusCircleIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/navigation'

export default function BottomNavigation() {
  const { user, status } = useAuth()
  const router = useRouter()
  const [_, setIsPostDialogOpen] = useAtom(postDialogAtom)
  const [__, setIsLoginDialogOpen] = useAtom(loginDialogAtom)
  const uid = user.uid
  const AccountURL = uid ? `/accounts/${uid}` : "/"
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
    <div className="w-full flex flex-row items-center justify-between px-[10%] py-2 bg-white/80 backdrop-blur-[5px] text-slate-500">
      <Link href={"/"} ><HomeIcon className='h-7 w-7' /></Link>
      <button onClick={handleOnFeedPage} ><NewspaperIcon className='h-7 w-7' onClick={handleAuth} /></button>
      <button onClick={handlePost}><PlusCircleIcon className='h-7 w-7' /></button>
      <Link href={"/explore"} ><MagnifyingGlassIcon className='h-7 w-7' /></Link>
      <Link href={AccountURL} onClick={handleAuth}>
        {
          user.uid && user.name && user.image ?
            <Image src={user.image} alt={user.name} width={200} height={200}
              className="h-7 w-7 rounded-full"></Image>
            :
            <UserCircleIcon className='h-7 w-7' />}</Link>
    </div>
  )
}
'use client'
import { useAuth } from '@/hooks/useAuth'
import { postDialogAtom } from '@/stores/dialog'
import { FireIcon, MagnifyingGlassIcon, NewspaperIcon, PlusCircleIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Image from 'next/image'
import Link from 'next/link'

export default function BottomNavigation() {
  const { user } = useAuth()
  const [_, setIsOpen] = useAtom(postDialogAtom)
  const uid = user.uid
  const AccountURL = uid ? `/accounts/${uid}` : "/login"
  return (
    <div className="w-full flex flex-row items-center justify-between px-[10%] py-2 bg-white/80 backdrop-blur-[5px] text-slate-500">
      <Link href={"/"} ><FireIcon className='h-7 w-7' /></Link>
      <Link href={"/"} ><NewspaperIcon className='h-7 w-7' /></Link>
      <button onClick={() => setIsOpen(true)}><PlusCircleIcon className='h-7 w-7' /></button>
      <Link href={"/"} ><MagnifyingGlassIcon className='h-7 w-7' /></Link>
      <Link href={AccountURL} >
        {
          user.uid && user.name && user.image ?
            <Image src={user.image} alt={user.name} width={200} height={200}
              className="h-7 w-7 rounded-full"></Image>
            :
            <UserCircleIcon className='h-7 w-7' />}</Link>
    </div>
  )
}
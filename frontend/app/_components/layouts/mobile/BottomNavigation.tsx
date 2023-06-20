'use client'
import { postDialogAtom } from '@/stores/dialog'
import { ChatBubbleOvalLeftEllipsisIcon, HomeIcon, MagnifyingGlassIcon, PlusCircleIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import { useSession } from 'next-auth/react'
import Link from 'next/link'

export default function BottomNavigation() {
  const { data: session } = useSession()
  const [_, setIsOpen] = useAtom(postDialogAtom)
  const uid = session?.user.uid
  const AccountURL = uid ? `/accounts/${uid}` : "/login"
  return (
    <div className="w-full flex flex-row items-center justify-between px-20 py-2">
      <Link href={"/"} ><HomeIcon className='h-7 w-7' /></Link>
      <Link href={"/"} ><ChatBubbleOvalLeftEllipsisIcon className='h-7 w-7' /></Link>
      <button onClick={() => setIsOpen(true)}><PlusCircleIcon className='h-7 w-7' /></button>
      <Link href={"/"} ><MagnifyingGlassIcon className='h-7 w-7' /></Link>
      <Link href={AccountURL} ><UserCircleIcon className='h-7 w-7' /></Link>
    </div>
  )
}
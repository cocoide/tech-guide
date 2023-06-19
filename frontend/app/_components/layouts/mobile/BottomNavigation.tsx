'use client'
import { ChatBubbleOvalLeftEllipsisIcon, HomeIcon, MagnifyingGlassIcon, PlusCircleIcon, UserCircleIcon } from '@heroicons/react/24/outline'
import { useSession } from 'next-auth/react'
import Link from 'next/link'

export default function BottomNavigation() {
  const { data: session } = useSession()
  const uid = session?.user.uid
  const nagigations = [
    { id: 'home', href: '/', icon: <HomeIcon className='h-7 w-7' /> },
    { id: 'discussion', href: '/', icon: <ChatBubbleOvalLeftEllipsisIcon className='h-7 w-7' /> },
    { id: 'post', href: '/', icon: <PlusCircleIcon className='h-7 w-7' /> },
    { id: 'search', href: '/', icon: <MagnifyingGlassIcon className='h-7 w-7' /> },
  ]
  const AccountURL = uid ? `/accounts/${uid}` : "/login"
  return (
    <div className="w-full flex flex-row items-center justify-between px-20 py-2">
      {nagigations.map(navi => (
        <Link href={navi.href} key={navi.id} className=""
        >{navi.icon}</Link>
      ))}
      <Link href={AccountURL} ><UserCircleIcon className='h-7 w-7' /></Link>
    </div>
  )
}
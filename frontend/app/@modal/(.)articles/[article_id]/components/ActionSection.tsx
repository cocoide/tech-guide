"use client"
import { useAuth } from '@/hooks/useAuth'
import { collectionDialogAtom, loginDialogAtom } from '@/stores/dialog'
import { BookmarkIcon, HandThumbDownIcon, HandThumbUpIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'

interface Props {
    articleId: number
}
export default function ActionSection({ articleId }: Props) {
    const [_, setLoginOpen] = useAtom(loginDialogAtom)
    const { status } = useAuth()
    const [__, setCollectionDialogOpen] = useAtom(collectionDialogAtom)
    function handleBookmark() {
        if (status === "unauthenticated") {
            setCollectionDialogOpen(articleId)
        }
        if (status === "authenticated") {
            setCollectionDialogOpen(articleId)
        }
    }
    return (
        <div className="w-full flex flex-row items-center  space-x-3">
            <div className="flex items-center rounded-full ring-1 ring-gray-100 bg-gray-100
            divide-x divide-gray-300  justify-center">
                <HandThumbUpIcon className='h-9 w-9 text-gray-500 hover:text-pink-300 hover:bg-pink-100 rounded-full p-1' />
                <HandThumbDownIcon className='h-9 w-9 text-gray-500 p-1' />
            </div>
            <button onClick={handleBookmark}
                className="flex items-center rounded-full ring-1 ring-gray-100 bg-gray-100
            justify-center p-1">
                <BookmarkIcon className='h-6 w-6 text-gray-500' />
                <div className="text-gray-600">保存</div>
            </button>
        </div>
    )
}
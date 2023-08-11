"use client"
import { articleAPI } from '@/app/_functions/article';
import { useAuth } from '@/hooks/useAuth';
import { loginDialogAtom } from '@/stores/dialog';
import { ChatBubbleOvalLeftEllipsisIcon, ChevronUpDownIcon } from '@heroicons/react/24/outline';
import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';
import { useState } from 'react';
import CommentView from './CommentView';

export default function CommentSection({ articleID }: { articleID: number }) {
    const { status } = useAuth()
    const [isViewOpen, setViewOpen] = useState(false)
    const [__, setLoginOpen] = useAtom(loginDialogAtom)
    const { data: comments } = useQuery({
        queryFn: async () => (await articleAPI.GetCommentsForArticle(articleID)).data,
        queryKey: ["comments_query"],
        enabled: isViewOpen
    })
    function handleOpenDialog() {
        if (status === "unauthenticated") {
            setLoginOpen(true)
            return
        }
        setViewOpen(!isViewOpen)
    }
    return (
        <>
            <div className="flex flex-col space-y-2 w-full p-2 border-y">
                <button onClick={handleOpenDialog} className=" w-full  flex flex-row justify-between items-center text-gray-500">
                    <div className="flex flex-row items-center space-x-1">
                <ChatBubbleOvalLeftEllipsisIcon className='h-5 w-5' />
                        <div className="text-sm">コメント</div>
            </div>
            <ChevronUpDownIcon className='h-7 w-7' />
                </button>
                {isViewOpen &&
                    <CommentView articleID={articleID}comments={comments} />
                }
            </div >
        </>
    )
}

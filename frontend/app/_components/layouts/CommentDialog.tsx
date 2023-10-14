"use client"
import ArticleCard from '@/app/(home)/_components/ArticleCard'
import { api } from '@/app/_functions/API'
import { useAuth } from '@/hooks/useAuth'
import { commentDialogAtom } from '@/stores/dialog'
import { ChatBubbleOvalLeftEllipsisIcon, XMarkIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Image from 'next/image'
import { ChangeEvent, useState } from 'react'
import toast from 'react-hot-toast'
import CustomDialog from '../elements/CustomDialog'

const CommentDialog = () => {
    const { user, token } = useAuth()
    const [comment, setComment] = useState("")
    const [commentDialogValue, setOpenCommentDialog] = useAtom(commentDialogAtom)
    function handleCommentChange(e: ChangeEvent<HTMLTextAreaElement>) {
        setComment(e.target.value)
    }
    async function handleSubmitComment(articleID: number) {
        if (comment.length < 1) {
            toast.error("コメントを入力して下さい")
            return
        }
        const { ok } = await api.pos(`/account/comment/${articleID}`, { "content": comment }, token)
        toast.loading("投稿中...")
        if (ok) {
            toast.dismiss()
            toast.success("投稿完了")
            setOpenCommentDialog(false)
        } else {
            toast.dismiss()
            toast.error("エラーが発生")
        }
    }
    return (
        <>
            {typeof commentDialogValue !== 'boolean' &&
                <CustomDialog
                    openAtom={commentDialogAtom}
                layout='mt-[150px] sm:my-[120px] z-50 sm:mx-[15%] md:mx-[20%] lg:mx-[25%] xl:mx-[27%] sm:rounded-xl'
                    content={
                        <div className='p-7 flex flex-col items-center w-full space-y-3'>
                            <div className="flex items-start w-[100%] space-x-3 justify-between">
                                <Image src={user?.image as string} width={70} height={70} alt={user?.name as string} className="custom-border h-[50px] w-[50px] rounded-full" />
                                <div className="flex flex-col w-[100%] justify-center items-center">
                                    <textarea onChange={handleCommentChange} value={comment} rows={1}
                                        className="w-[100%] min-h-auto   focus:ring-transparent ring-none border-none resize-none min-h-15 bg-gray-50 dark:bg-gray-700 rounded-xl dark:text-white" ></textarea>
                                </div>
                                <button onClick={() => setOpenCommentDialog(false)}>
                                    <XMarkIcon className='h-8 w-8 text-gray-400 p-[3px]  duration-500 rounded-md' />
                                </button>
                            </div>
                            <div className="min-w-[350px] lg:min-w-[500px]">
                                <ArticleCard article={commentDialogValue} />
                            </div>
                            <div className="flex flex-row justify-end w-[100%]">
                                <button onClick={() => handleSubmitComment(commentDialogValue.id)} className=" bg-cyan-300 bg-primary rounded-xl shadow-sm p-2 text-bold text-white flex items-center"
                                ><ChatBubbleOvalLeftEllipsisIcon className="mr-1 h-5 w-5 text-white text-sm font-bold" />投稿</button>
                            </div>
                                <div className="w-[100%] h-[70px] sm:hidden"></div>
                        </div>
                    }
                />
            }
        </>
    )
}
export default CommentDialog
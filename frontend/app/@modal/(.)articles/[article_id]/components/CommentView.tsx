"use client"
import { api } from '@/app/_functions/API';
import { useAuth } from '@/hooks/useAuth';
import { useSession } from '@/hooks/useSession';
import { Comment } from '@/types/model';
import Image from 'next/image';
import { ChangeEvent, RefObject, useState } from 'react';
import toast from 'react-hot-toast';


interface Props {
    articleID: number
    comments?: Comment[]
    inputRef: RefObject<HTMLInputElement>
}
const CommentView = ({ articleID, comments, inputRef }: Props) => {
    const { token } = useAuth()
    const session = useSession()
    const [comment, setComment] = useState("")
    function handleCommentChange(e: ChangeEvent<HTMLInputElement>) {
        setComment(e.target.value)
    }
    async function handleSubmit() {
        if (comment.length < 1) {
            toast.error("コメントを入力して下さい")
            return
        }
        const { ok } = await api.pos(`/account/comment/${articleID}`, { "content": comment }, token)
        toast.loading("投稿中...")
        if (ok) {
            toast.dismiss()
            toast.success("投稿完了")
        } else {
            toast.dismiss()
            toast.error("エラーが発生")
        }
    }
    return (
        <div className='flex flex-col space-y-2 w-full'>
            <div className="flex flex-row items-center w-full space-x-3">
                <Image src={session.avatar_url} width={70} height={70} alt={session.display_name} className="h-7 w-7 rounded-full bg-shadow" />
                <input ref={inputRef} onChange={handleCommentChange} className="ring-none border-none w-full p-1 text-sm text-gray-500 focus:ring-transparent " placeholder="コメントを入力" />
                <button onClick={handleSubmit} className="text-sm bg-cyan-300 w-20 h-auto text-white p-1 rounded-md">投稿</button>
            </div>
            {comments&&comments?.length>0&&
            <div className="flex flex-col space-y-3 bg-gray-100 p-2 rounded-md">
                {comments?.map(comment => (
                    <div key={comment.id + "comment"} className="flex flex-row space-x-2 items-center">
                        {comment.account?.avatar_url ?
                            <Image src={comment?.account?.avatar_url} alt={"comment" + articleID} className="h-5 w-5 rounded-full" width={200} height={200} />
                            :
                            <div className="h-7 w-7 bg-gray-200 rounded-full"></div>
                        }
                        <div className="text-gray-500 text-sm ">{comment?.content}</div>
                    </div>
                ))}
            </div>
            }
        </div>
    )
}
export default CommentView
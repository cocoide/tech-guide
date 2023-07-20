'use client'
import { api } from '@/app/_functions/API'
import { useAuth } from '@/hooks/useAuth'
import useAutosizeTextArea from '@/hooks/useAutosizeTextArea'
import { postDialogAtom } from '@/stores/dialog'
import { PlusCircleIcon, XMarkIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import Image from 'next/image'
import { ChangeEvent, useEffect, useRef, useState } from 'react'
import { toast } from 'react-hot-toast'
import CircleLoading from '../animations/CircleLoading'
import CustomDialog from '../elements/CustomDialog'

const PostDialog = () => {
    const [_, setDialogOpen] = useAtom(postDialogAtom)
    const [ogp, setOGP] = useState<OGP | null>(null)
    const [url, setUrl] = useState('')
    const [typingTimeout, setTypingTimeout] = useState<NodeJS.Timeout | null>(null)
    const [isFetching, setIsFetching] = useState(false)
    const { user, token } = useAuth()

    const [comment, setComment] = useState("")
    const textAreaRef = useRef<HTMLTextAreaElement>(null);
    useAutosizeTextArea(textAreaRef.current, comment);
    function handleCommentChange(e: ChangeEvent<HTMLTextAreaElement>) {
        setComment(e.target.value)
    }
    async function handleSubmitPost(url: string, content: string) {
        if (ogp) {
            const { ok } = await api.pos("/account/comment", { "original_url": url, "content": content }, token)
            if (!ok) {
                toast.error("エラーが発生")
            } else {
                toast.success("投稿完了")
                setDialogOpen(false)
                setOGP(null)
            }
        }
    }
    type OGP = {
        title: string
        thumbnail: string
        sitename: string
        description: string
    }
    useEffect(() => {
        const fetchOGP = async () => {
            const { data: ogp, ok } = await api.get<OGP>(`/ogp?url=${url}`)
            if (ok && ogp) {
                setOGP(ogp)
                setIsFetching(false)
            }
            if (!ok) {
                setIsFetching(false)
            }
        }
        if (typingTimeout) {
            clearTimeout(typingTimeout)
        }
        if (url.trim() !== '') {
            setIsFetching(true)
            const timeout = setTimeout(() => {
                fetchOGP()
            }, 1500)
            setTypingTimeout(timeout)
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [url])
    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setUrl(e.target.value)
        setOGP(null)
    }
    return (
        <CustomDialog
            closeFunc={() => setUrl("")}
            openAtom={postDialogAtom}
            layout='overflow-y-auto  rounded-xl sm:mx-[10%] sm:my-20 md:mx-[15%] lg:mx-[25%] md:my-[100px] mt-[150px]'
            content={
                <div className="h-full w-full p-10 flex flex-col justify-between">
                    <div className="flex flex-col space-y-7 text-gray-600">
                        <div className="flex flex-row items-center justify-between">
                            <button className='flex flex-row items-center' onClick={() => setDialogOpen(false)}>
                                <XMarkIcon className='h-5 w-5 text-slate-600' />
                            </button>
                            <button onClick={async () => await handleSubmitPost(url, comment)} className=" bg-cyan-300 bg-primary rounded-xl shadow-sm p-2 text-bold text-white flex items-center"
                            ><PlusCircleIcon className="mr-1 h-5 w-5 text-white" />投稿する</button>
                        </div>
                        <input
                            onChange={handleInputChange}
                            value={url}
                            className="w-[100%] bg-gray-50 rounded-md p-[5px] ring-1 ring-gray-300"
                            placeholder='https://'></input>
                    </div>
                    {url && isFetching &&
                        <div className="w-full flex justify-center">
                            <CircleLoading />
                        </div>
                    }
                    {ogp?.title &&
                        <div className="text-center text-gray-600">{ogp?.title}</div>
                    }
                    {ogp?.thumbnail &&
                        // eslint-disable-next-line @next/next/no-img-element 
                        <img src={ogp?.thumbnail} alt={ogp?.thumbnail}
                            className='rounded-md  h-[150px] mx-auto' />
                    }
                    <div className="flex items-start w-[100%] space-x-3 pt-8">
                        <Image src={user?.image as string} width={70} height={70} alt={user?.name as string} className="h-[50px] w-[50px] rounded-full bg-shadow" />
                        <div className="flex flex-col w-[100%] justify-center items-center">
                            <textarea ref={textAreaRef} onChange={handleCommentChange} value={comment} rows={1}
                                className="w-[100%] min-h-auto   focus:ring-transparent ring-none border-none resize-none min-h-15" placeholder="コメントを入力"></textarea>
                            <div className="border w-full border-shadow mb-5"></div>
                        </div>
                    </div>
                </div>
            } />
    )
}
export default PostDialog

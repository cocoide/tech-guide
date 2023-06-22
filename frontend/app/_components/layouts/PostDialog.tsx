'use client'
import { api } from '@/app/_functions/API'
import { postDialogAtom } from '@/stores/dialog'
import { DocumentTextIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import { useEffect, useState } from 'react'
import { toast } from 'react-hot-toast'
import CircleLoading from '../animations/CircleLoading'
import CustomDialog from '../elements/CustomDialog'

const PostDialog = () => {
    const [_, setDialogOpen] = useAtom(postDialogAtom)
    const [ogp, setOGP] = useState<OGP | null>(null)
    const [url, setUrl] = useState('')
    const [typingTimeout, setTypingTimeout] = useState<NodeJS.Timeout | null>(null)
    const [isFetching, setIsFetching] = useState(false)
    async function SubmitArticle(url: string) {
        if (ogp) {
            const { ok } = await api.pos("/article", { "original_url": url })
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
            layout='rounded-xl sm:mx-[15%] sm:my-20 md:mx-[20%] lg:mx-[30%] md:my-[100px] mt-[150px]'
            content={
                <div className="h-full w-full py-10 px-12 flex flex-col justify-between">
                    <div className="flex flex-col space-y-5 text-gray-600">
                        <div className="text-center flex items-center space-x-3 justify-center">
                            <DocumentTextIcon className='h-5 w-5' />
                            <div className=""> 投稿する</div>
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
                    <button disabled={isFetching} onClick={async () => await SubmitArticle(url)}
                        className='text-cyan-300 ring-1 ring-cyan-400  p-[5px] rounded-md flex items-center 
                        justify-center space-x-2'
                    >投稿完了</button>
                </div>
            } />
    )
}
export default PostDialog

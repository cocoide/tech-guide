'use client'
import { collectionAPI } from '@/app/(dashboard)/accounts/[account_id]/_functions/collection'
import { api } from '@/app/_functions/API'
import { Collection } from '@/app/_models'
import { collectionDialogAtom } from '@/stores/dialog'
import { ChevronLeftIcon } from '@heroicons/react/24/outline'
import { PlusCircleIcon } from '@heroicons/react/24/solid'
import { useAtom } from 'jotai'
import { useSession } from 'next-auth/react'
import { useEffect, useState } from 'react'
import { useForm } from "react-hook-form"
import { toast } from 'react-hot-toast'
import CustomDialog from '../elements/CustomDialog'

export type NewCollectionRequest = {
    name: string,
    description: string,
    visibility: number,
    article_id: number,
}

const CollectionDialog = () => {
    const [dialogAtom, setDialogAtom] = useAtom(collectionDialogAtom)
    const [isNewCollection, setIsNewCollection] = useState(false)
    const [collections, setCollections] = useState<Collection[] | undefined>()
    const { data: session } = useSession()
    useEffect(() => {
        if (dialogAtom !== false) {
            (async () => {
                const { data: collections, ok } = await collectionAPI.getCollectionForBookmark(session?.token)
                if (!ok) {
                    toast.error("エラーが発生")
                }
                setCollections(collections)
            })()
        }
    }, [dialogAtom, session?.token])
    type BookmarkRequest = { collection_id: number, article_id: number }
    async function handleBookmark(articleId: number, collectionId: number) {
        toast.loading("保存中...")
        const { ok, status } = await api.pos<BookmarkRequest>("/account/bookmark", { article_id: articleId, collection_id: collectionId }, session?.token)
        if (!ok) {
            if (status == 409) {
                return toast.error("すでに保存されてあります")
            }
            toast.error("エラーが発生")
        }
        toast.dismiss()
        toast.success("保存完了")
        setDialogAtom(false)
    }
    const { register, handleSubmit }
        = useForm<NewCollectionRequest>();

    async function handleNewCollectionSubmit(data: NewCollectionRequest) {
        if (typeof dialogAtom == "number") {
            toast.loading("作成中...")
            const { ok } = await collectionAPI.makeCollectionWithBookmark(data, dialogAtom)
            if (!ok) {
                return toast.error("エラーが発生")
            }
            toast.dismiss()
            toast.success("作成完了")
        }
    }

    return (
        <CustomDialog
            openAtom={collectionDialogAtom}
            closeFunc={() => setIsNewCollection(false)}
            layout='my-[120px] bg-white z-50 mx-[5%] sm:mx-[15%] md:mx-[20%] lg:mx-[25%] rounded-xl'
            content={
                <div className='flex flex-col h-full p-10 space-y-5 w-full text-slate-700'>
                    {isNewCollection ?
                        <>
                            <div className="flex flex-row items-center w-full justify-center">
                                <button onClick={() => setIsNewCollection(false)} className="absolute left-10">
                                    <ChevronLeftIcon className='h-5 w-5 text-gray-600' />
                                </button>
                                <div className="text-center">コレクションを作成</div>
                            </div>
                            <form onSubmit={handleSubmit(handleNewCollectionSubmit)} className='flex flex-col p-5 space-y-5 h-full text-slate-500'>
                                <div className="w-full">
                                    <div className="">コレクション名</div>
                                    <input {...register("name", { required: true })} className='p-1 ring-1 ring-slate-200 rounded-md w-full' />
                                </div>
                                <div className="w-full">
                                    <div className="">詳細</div>
                                    <input {...register("description")} className='p-1 ring-1 ring-slate-200 rounded-md w-full' />
                                </div>
                                <p className='h-full'></p>
                                <button type="submit" className='ring-1 ring-cyan-300 rounded-xl p-[3px] text-cyan-300'>作成完了</button>
                            </form>
                        </>
                        :
                        <>
                            <div className="text-center">保存先</div>
                    <div className="overflow-y-auto h-full flex flex-col space-y-3">
                        {typeof dialogAtom == 'number' && collections?.map((c => (
                            <button onClick={async () => handleBookmark(dialogAtom, c.id)}
                                key={c.id} className="flex flex-row items-center space-x-5">
                                {c.articles[0]?.thumbnail_url ?
                                    // eslint-disable-next-line @next/next/no-img-element
                                    <img src={c.articles[0].thumbnail_url} alt={c.articles[0].title} width={200} className='rounded-md w-[200px] h-[100px] shadow-md' />
                                    :
                                    <div className="rounded-md w-[100px] h-[60px] bg-slate-200 shadow-[3px]"></div>
                                }
                                <div className="">{c.name}</div>
                            </button>
                    )))}
                    </div>
                            <button onClick={() => setIsNewCollection(true)} className="flex flex-row items-center space-x-2 ring-1 ring-slate-200 rounded-full mx-auto p-[3px]">
                                <PlusCircleIcon className='h-7 w-7 text-cyan-300' /><div>新しいコレクションを作成</div></button>
                        </>
                    }
                </div>
            }
        />
    )
}
export default CollectionDialog
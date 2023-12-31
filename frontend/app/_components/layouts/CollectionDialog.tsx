'use client'
import { collectionAPI } from '@/app/(dashboard)/accounts/[account_id]/_functions/collection'
import { api } from '@/app/_functions/API'
import { collectionDialogAtom } from '@/stores/dialog'
import { Collection } from '@/types/model'
import { ChevronLeftIcon, NewspaperIcon, PlusIcon, XMarkIcon } from '@heroicons/react/24/outline'
import { useAtom } from 'jotai'
import { useEffect, useState } from 'react'
import { useForm } from "react-hook-form"
import { toast } from 'react-hot-toast'
import CustomDialog from '../elements/CustomDialog'
import {useAuth} from "@/hooks/useAuth";

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
    const { token } = useAuth()
    useEffect(() => {
        if (dialogAtom !== false) {
            (async () => {
                const { data: collections, ok } = await collectionAPI.getCollectionForBookmark(token)
                if (!ok) {
                    toast.error("エラーが発生")
                }
                setCollections(collections)
            })()
        }
    }, [dialogAtom, token])
    type BookmarkRequest = { collection_id: number, article_id: number }
    async function handleBookmark(articleId: number, collectionId: number) {
        toast.loading("保存中...")
        const { ok, status } = await api.pos<BookmarkRequest>("/account/bookmark", { article_id: articleId, collection_id: collectionId }, token)
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
            layout='mt-[120px] sm:mb-[120px]  bg-white z-50 sm:mx-[15%] md:mx-[20%] lg:mx-[25%] rounded-xl'
            content={
                <div className='flex flex-col h-full p-10 space-y-5 w-full text-gray-600 dark:text-gray-100'>
                    {isNewCollection ?
                        <>
                            <div className="flex flex-row items-center w-full justify-center">
                                <button onClick={() => setIsNewCollection(false)} className="absolute left-10">
                                    <ChevronLeftIcon className='h-5 w-5' />
                                </button>
                                <div className="text-center">コレクションを作成</div>
                            </div>
                            <form onSubmit={handleSubmit(handleNewCollectionSubmit)} className='flex flex-col p-5 space-y-5 h-full'>
                                <div className="w-full">
                                    <div className="">コレクション名</div>
                                    <input {...register("name", { required: true })} className='p-1 appearance-none outline outline-slate-200 rounded-md w-full' />
                                </div>
                                <div className="w-full">
                                    <div className="">詳細</div>
                                    <input {...register("description")} className='p-1 appearance-none outline outline-slate-200 rounded-md w-full' />
                                </div>
                                <p className='h-full'></p>
                                <button type="submit" className='ring-1 ring-cyan-300 rounded-xl p-[3px] text-cyan-300'>作成完了</button>
                            </form>
                        </>
                        :
                        <>
                            <div className="relative">
                                <button onClick={() => setDialogAtom(false)}>
                                    <XMarkIcon className='h-8 w-8 text-gray-400 p-[3px]  duration-500 rounded-md absolute right-0' />
                                </button>
                            <div className="text-center">コレクションを選ぶ</div>
                            </div>
                    <div className="overflow-y-auto h-full flex flex-col space-y-3">
                        {typeof dialogAtom == 'number' && collections?.map((c => (
                            <button onClick={async () => handleBookmark(dialogAtom, c.id)} key={c.name + c.id}
                                className="rounded-md flex flex-row justify-between  custom-border
                                             shadow-sm w-full h-[120px] p-2 items-center
                                            ">
                                <div className="flex flex-col space-y-3 w-full h-full p-2">
                                    <div className="text-md  w-full custom-badge">
                                        <div> {c.name}</div>
                                        <div className="bg-cyan-50 text-cyan-300 border-cyan-300 border-[0.5px] rounded-full h-5 w-5 flex justify-center items-center"
                                        >{collections.length}</div>
                                    </div>
                                </div>
                                {/* image section */}
                                {c.articles[0]?.thumbnail_url ?
                                    // eslint-disable-next-line @next/next/no-img-element
                                    <img src={c.articles[0].thumbnail_url} alt={c.articles[0].title} width={200}
                                        className='rounded-md w-[50%] h-[100px] shadow-md overflow-hidden' />
                                    :
                                    <div className="flex items-center justify-center rounded-md w-[50%] h-[100px] bg-gray-100 shadow-[2px] custom-border">
                                        <NewspaperIcon className="h-7 w-7 text-gray-500" />
                                    </div>
                                }
                            </button>

                    )))}
                    </div>
                            <button onClick={() => setIsNewCollection(true)} className="flex flex-row items-center space-x-2
                            custom-border rounded-md justify-center
                             w-full p-[3px]">
                                <PlusIcon className='h-5 w-5' /><div>コレクションを作成</div></button>
                        </>
                    }
                </div>
            }
        />
    )
}
export default CollectionDialog
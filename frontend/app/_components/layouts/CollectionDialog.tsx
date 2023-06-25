'use client'
import { collectionAPI } from '@/app/(dashboard)/accounts/[account_id]/_functions/collection'
import { api } from '@/app/_functions/API'
import { Collection } from '@/app/_models'
import { collectionDialogAtom } from '@/stores/dialog'
import { PlusCircleIcon } from '@heroicons/react/24/solid'
import { useAtom } from 'jotai'
import { useSession } from 'next-auth/react'
import { useSearchParams } from 'next/navigation'
import { useEffect, useState } from 'react'
import { toast } from 'react-hot-toast'
import CustomDialog from '../elements/CustomDialog'

const CollectionDialog = () => {
    const [dialogAtom, setDialogAtom] = useAtom(collectionDialogAtom)
    const queryParam = useSearchParams()
    const d = queryParam.get("test")
    const [collections, setCollections] = useState<Collection[] | undefined>()
    const { data: session } = useSession()
    useEffect(() => {
        if (dialogAtom != false) {
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

    return (
        <CustomDialog
            openAtom={collectionDialogAtom}
            layout='my-[120px] bg-white z-50 mx-[5%] sm:mx-[15%] md:mx-[20%] lg:mx-[25%] rounded-xl'
            content={
                <div className='flex flex-col h-full p-10 space-y-5 w-full text-slate-700'>
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
                    <div className="flex flex-row items-center space-x-2 ring-1 ring-slate-200 rounded-full mx-auto p-[3px]">
                        <PlusCircleIcon className='h-7 w-7 text-cyan-300' /><div>新しいコレクションを作成</div></div>
                </div>
            }
        />
    )
}
export default CollectionDialog
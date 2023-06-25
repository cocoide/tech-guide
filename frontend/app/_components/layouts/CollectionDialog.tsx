'use client'
import { collectionAPI } from '@/app/(dashboard)/accounts/[account_id]/_functions/collection'
import { Collection } from '@/app/_models'
import { collectionDialogAtom } from '@/stores/dialog'
import { useAtom } from 'jotai'
import { useSession } from 'next-auth/react'
import { useEffect, useState } from 'react'
import CustomDialog from '../elements/CustomDialog'

const CollectionDialog = () => {
    const [dialogAtom, setDialogAtom] = useAtom(collectionDialogAtom)
    const [collections, setCollections] = useState<Collection[] | undefined>()
    const { data: session } = useSession()
    useEffect(() => {
        if (dialogAtom != false) {
            (async () => {
                const { data: collections, error } = await collectionAPI.getCollectionForBookmark(session?.token)
                console.log(error)
                setCollections(collections)
            })()
        }
    }, [dialogAtom, session?.token])
    return (
        <CustomDialog
            openAtom={collectionDialogAtom}
            layout='my-[150px] bg-white z-50 sm:mx-[15%] md:mx-[20%] lg:mx-[25%] sm:rounded-xl'
            content={
                <div className='flex flex-col h-full p-10'>
                    <div className="">コレクションに保存する</div>
                    {collections?.map((c => (
                        <div key={c.id} className="">{c.name}</div>
                    )))}
                </div>
            }
        />
    )
}
export default CollectionDialog
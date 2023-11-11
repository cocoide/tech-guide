"use client"

import SectionHeader from '@/app/_components/layouts/desktop/SectionHeader'
import { sourceAPI } from '@/app/_functions/source'
import { useAuth } from '@/hooks/useAuth'
import { loginDialogAtom } from '@/stores/dialog'
import { Source } from '@/types/model'
import { useQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
import Image from 'next/image'
import toast from 'react-hot-toast'
import HeaderFollowButton from './HeaderFollowButton'

const SourceSectionHeader = ({ source }: { source: Source }) => {

    const { data: isFollowing, refetch: refetchFollowingState } = useQuery({
        queryKey: [`${source.id}.source_query`],
        queryFn: async () => (await sourceAPI.CheckFollow(source.id, token)).data,
    })
    const { status, token } = useAuth()
    const [_, setLoginOpen] = useAtom(loginDialogAtom)
    async function handleDoFollow() {
        if (status === "unauthenticated") {
            setLoginOpen(true)
            return
        }
        const { ok } = await sourceAPI.DoFollow(source.id, token)
        if (!ok) {
            toast.error("エラーが発生")
        }
        await refetchFollowingState()
    }
    async function handleUnFollow() {
        if (status === "unauthenticated") {
            setLoginOpen(true)
            return
        }
        const { ok } = await sourceAPI.UnFollow(source.id, token)
        if (!ok) {
            toast.error("エラーが発生")
        }
        await refetchFollowingState()
    }


    useQuery
    return (
        <SectionHeader
            titleItem={<div className='custom-badge text-gray-500'>
                <Image src={source?.icon_url as string} alt={source?.name as string} width={100} height={100}
                    className='h-7 w-7 rounded-full' />
                <div>{source?.name}</div>
            </div>}
            rightItem={<HeaderFollowButton
                doFollowFunc={handleDoFollow}
                unFollowFunc={handleUnFollow}
                isFollowing={isFollowing} />}
        />
    )
}
export default SourceSectionHeader
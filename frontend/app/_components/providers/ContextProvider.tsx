"use client"
import { api } from '@/app/_functions/API'
import { useAuth } from '@/hooks/useAuth'
import { topicDialogAtom } from '@/stores/dialog'
import { UserSession } from '@/types/model'
import { useAtom } from 'jotai'
import { ReactNode, useEffect } from 'react'

const ContextProvider = ({ children }: { children: ReactNode }) => {
    const { status, token } = useAuth()
    const [_, open] = useAtom(topicDialogAtom)
    useEffect(() => {
        (async () => {
            const { data: session } = await api.get<UserSession>("/account/session", "no-store", token)
            console.log(session)
            open(true)
        })()
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [status, token])
    return (
        <>{children}</>
    )
}
export default ContextProvider
'use client'
import { clsx } from '@/utils/clsx'
import { PrimitiveAtom, useAtom } from 'jotai'
import { ReactNode } from 'react'
import '../../../styles/animation.css'

interface Props {
    content: ReactNode
    button?: ReactNode
    openAtom: PrimitiveAtom<boolean>
    layout?: string
    openFunc?: () => void
    closeFunc?: () => void
}
export default function CustomDialog({ content, button, openAtom, layout, openFunc, closeFunc }: Props) {
    const [isOpen, setisOpen] = useAtom(openAtom)
    function handleClose() {
        closeFunc
        setisOpen(false)
    }
    function handleOpen() {
        openFunc
        setisOpen(true)
    }
    return (
        <>
            {isOpen ?
                <>
                    <button onClick={handleClose}
                        className="z-40 bg-gray-500/30  fixed inset-0 backdrop-blur-[3px] animate-appear"></button>
                    <div className={clsx(layout ? layout : "",
                        "z-50 fixed inset-0 sm:animate-scale animate-slideUp bg-white duration-700 shadow-[5px] flex items-center justify-center")}
                    >{content}</div>
                </>
                :
                <button onClick={handleOpen}>{button}</button>
            }
        </>
    )
}
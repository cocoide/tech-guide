'use client'
import { clsx } from '@/utils/clsx'
import { PrimitiveAtom, useAtom } from 'jotai'
import { ReactNode } from 'react'
import '../../../styles/animation.css'

interface Props<T> {
    content: ReactNode
    button?: ReactNode
    openAtom: PrimitiveAtom<boolean | T>
    layout?: string
    openFunc?: () => void
    closeFunc?: () => void
}
export default function CustomDialog<T>({ content, button, openAtom, layout, openFunc, closeFunc }: Props<T>) {
    const [isOpen, setIsOpen] = useAtom(openAtom)
    function handleClose() {
        if (closeFunc) {
            closeFunc()
        }
        setIsOpen(false)
    }
    function handleOpen() {
        if (openFunc) {
            openFunc()
        }
        setIsOpen(true)
    }
    return (
        <>
            {isOpen ?
                <>
                    <button onClick={handleClose}
                        className="z-40 bg-gray-500/30  fixed inset-0 backdrop-blur-[3px] animate-appear"></button>
                    <div className={clsx(layout ? layout : "",
                        "z-50 fixed inset-0 animate-appear bg-white dark:bg-black duration-700 shadow-[5px] flex items-center justify-center")}
                    >{content}</div>
                </>
                :
                <button onClick={handleOpen}>{button}</button>
            }
        </>
    )
}
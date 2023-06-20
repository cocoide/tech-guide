'use client'
import { PrimitiveAtom, useAtom } from 'jotai'
import { ReactNode } from 'react'
import '../../../styles/animation.css'

interface Props {
    content: ReactNode
    button?: ReactNode
    openAtom: PrimitiveAtom<boolean>
}
export default function CustomDialog({ content, button, openAtom }: Props) {
    const [isOpen, setisOpen] = useAtom(openAtom)
    function handleClose() {
        setisOpen(false)
    }
    function handleOpen() {
        setisOpen(true)
    }
    return (
        <>
            {isOpen ?
                <>
                    <button onClick={handleClose}
                        className="z-40 bg-gray-500/30  fixed inset-0 backdrop-blur-[3px] animate-appear"></button>
                    <div className="
                    z-50 fixed inset-0 sm:animate-scale animate-slideUp 
                    bg-white duration-700 rounded-xl shadow-[5px]
                    sm:mx-[15%] sm:my-20 md:mx-[20%] lg:mx-[30%] md:my-[100px] mt-[150px] 
                    flex items-center justify-center"
                    >{content}</div>
                </>
                :
                <button onClick={handleOpen}>{button}</button>
            }
        </>
    )
}
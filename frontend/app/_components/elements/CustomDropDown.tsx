'use client'

import { ReactNode, useState } from 'react'
import '../../../styles/animation.css'

interface Props {
    button?: ReactNode
    layout?: string
    menues: Menue[]
    option?: ReactNode
}
export type Menue = {
    icon?: ReactNode
    href?: string
    label: string
}

export default function CustomDropDown({ button, menues, option }: Props) {
    const [isOpen, setIsOpen] = useState(false)
    function toggleOpen() {
        setIsOpen(!isOpen)
    }
    return (
        <div>
            <div onClick={toggleOpen}>{button}</div>
            {isOpen &&
                <>
                    <div className="absolute right-2 z-10 p-3
                mt-2 min-w-[200px] origin-top-right bg-white rounded-xl ring-1 ring-slate-100
                flex flex-col space-y-3
                shadow-[rgba(0,_0,_0,_0.24)_0px_3px_8px]
                animate-slideTinyUp duration-500
                text-slate-500
                ">
                        {menues.map(menue => (
                            <div key={menue.label} className="flex flex-row items-center p-1 space-x-2">
                                <div>{menue.icon}</div>
                                <div> {menue.label}</div>
                            </div>
                        ))}
                        {option}
                    </div>
                </>
            }
        </div >
    )
}
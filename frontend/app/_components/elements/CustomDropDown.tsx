'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { ReactNode, useEffect, useState } from 'react'
import '../../../styles/animation.css'

interface Props {
    header?: ReactNode
    button?: ReactNode
    layout?: string
    menues: Menue[]
    footer?: ReactNode
}
export type Menue = {
    icon?: ReactNode
    href: string
    label: string
}

export default function CustomDropDown({ header, button, menues, footer }: Props) {
    const [isOpen, setIsOpen] = useState(false)
    const pathname = usePathname();
    function toggleOpen() {
        setIsOpen(!isOpen)
    }
    useEffect(() => {
        setIsOpen(false);
    }, [pathname]);
    return (
        <div>
            <div onClick={toggleOpen}>{button}</div>
            {isOpen &&
                <>
                <div className="absolute right-5 z-10 p-3
                mt-2 min-w-[250px] origin-top-right bg-white rounded-md
                ring-1 ring-slate-100
                flex flex-col space-y-3
                shadow-[rgba(0,_0,_0,_0.24)_0px_3px_8px]
                animate-slideTinyUp duration-500
                text-slate-500
                ">
                    {header}
                        {menues.map(menue => (
                            <Link href={menue.href} key={menue.label} className="flex flex-row items-center p-1 space-x-2">
                                <div>{menue.icon}</div>
                                <div> {menue.label}</div>
                            </Link>
                        ))}
                    {footer}
                    </div>
                </>
            }
        </div >
    )
}
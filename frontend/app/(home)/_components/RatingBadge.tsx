"use client"
import { ArticleRating } from '@/types/model';
import { clsx } from '@/utils/clsx';
import { FireIcon, TrophyIcon } from '@heroicons/react/24/outline';
import Image from 'next/image';
import { ReactNode, useState } from 'react';

export const RatingBadge = ({ rating, domain }: { rating: ArticleRating, domain: string }) => {
    const [isOpen, setIsOpen] = useState(false)
    function handleOpen() {
        setIsOpen(!isOpen)
    }
    var text_style: string
    var icon: ReactNode = null
    const count = rating.hatena_stocks + rating.origin_stocks + rating.owned_stocks + rating.pocket_stocks
    const hatena = rating.hatena_stocks
    const origin = rating.origin_stocks
    const owned = rating.owned_stocks
    const pocket = rating.pocket_stocks
    switch (true) {
        case count >= 1000:
            text_style = "text-yellow-400  dark:text-yellow-300 font-bold"
            icon = <TrophyIcon className='h-5 w-5' />
            break;
        case 999 >= count && count >= 100:
            text_style = "text-red-400 dark:text-red-300 font-bold"
            icon = <FireIcon className='h-5 w-5' />
            break;
        default:
            text_style = "text-slate-300 dark:text-slate-100 font-bold"
    }
    return (
        <button onClick={handleOpen} className={clsx("relative flex flex-row items-center space-2 p-[2px]", text_style)}>
            {icon != null && icon}
            <div className="">{count} pt</div>
            {isOpen &&
                <button onClick={handleOpen} className="z-30 bg-black/10  fixed inset-0"></button>
            }
            {isOpen &&
                <div className="absolute -bottom-8 left-0 z-40 p-2 rounded-xl bg-white/80 backdrop-blur-sm
                flex flex-row items-center space-x-7 w-[200px]">
                    <UnitRateBadge
                        icon={<Image src={"/pocket.png"} width={50} height={50} alt='pocket' className='h-5 w-5' />}
                        count={pocket}
                    />
                    <UnitRateBadge
                        icon={<Image src={"/hatena.png"} width={50} height={50} alt='hatena' className='h-5 w-5' />}
                        count={hatena}
                    />
                    <UnitRateBadge
                        icon={<Image src={"/logo.svg"} width={50} height={50} alt='origin' className='h-6 w-6' />}
                        count={owned}
                    />
                    {domain === "qiita.com" ?
                        <UnitRateBadge
                            icon={<Image src={"/source/qiita.png"} width={50} height={50} alt='origin' className='h-5 w-5' />}
                            count={origin} />
                        :
                        domain === "github.com" ?
                            <UnitRateBadge
                                icon={<Image src={"/source/github.webp"} width={50} height={50} alt='origin' className='h-5 w-5' />}
                                count={origin} />
                            :
                            <></>
                    }
                </div>
            }
        </button>
    )
}

interface UnitRateBadgeProps {
    icon: ReactNode
    count: number
}

const UnitRateBadge = ({ icon, count }: UnitRateBadgeProps) => {
    return (
        <>
            {count !== 0 &&
                <div className="custom-badge text-gray-500">
                    {icon}
                    <div className="text-sm">{count}</div>
                </div>
}
        </>
    )
}
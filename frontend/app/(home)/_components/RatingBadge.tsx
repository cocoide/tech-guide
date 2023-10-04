import { clsx } from '@/utils/clsx';
import { FireIcon, TrophyIcon } from '@heroicons/react/24/outline';
import { ReactNode } from 'react';

export const RatingBadge = ({ count }: { count: number }) => {
    var text_style: string
    var icon: ReactNode = null

    switch (true) {
        case count >= 1000:
            text_style = "text-white bg-yellow-500 font-bold rounded-xl"
            icon = <TrophyIcon className='h-3 w-3' />
        case 999 >= count && count >= 100:
            text_style = "text-red-500 font-bold p-[2px]"
            icon = <FireIcon className='h-3 w-3' />
        case 99 >= count && count >= 50:
            text_style = "text-slate-500 dark:text-slate-100 font-bold"
        default:
            text_style = "text-slate-400 dark:text-slate-100 font-semibold"
    }
    return (
        <div className={clsx("flex flex-row items-center space-1 p-[2px]", text_style)}>
            <div className=""> {count} pt</div>
            {icon != null && icon}
        </div>
    )
}
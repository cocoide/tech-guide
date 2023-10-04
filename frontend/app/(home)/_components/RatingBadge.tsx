import { clsx } from '@/utils/clsx';
import { FireIcon, TrophyIcon } from '@heroicons/react/24/outline';
import { ReactNode } from 'react';

export const RatingBadge = ({ count }: { count: number }) => {
    var text_style: string
    var icon: ReactNode = null

    switch (true) {
        case count >= 1000:
            text_style = "text-yellow-300  dark:text-yellow-100 font-bold"
            icon = <TrophyIcon className='h-5 w-5' />
            break;
        case 999 >= count && count >= 100:
            text_style = "text-red-300 dark:text-red-200 font-bold"
            icon = <FireIcon className='h-5 w-5' />
            break;
        case 99 >= count && count >= 50:
            text_style = "text-slate-300 dark:text-slate-100 font-bold"
            break;
        default:
            text_style = "text-slate-300 dark:text-slate-100 font-bold"
    }
    return (
        <div className={clsx("flex flex-row items-center space-2 p-[2px]", text_style)}>
            {icon != null && icon}
            <div className=""> {count} pt</div>
        </div>
    )
}
import { clsx } from '@/utils/clsx';

const ProgressBar = ({ step }: { step?: number }) => {
    let width_style;
    switch (step) {
        case 1:
            width_style = "w-2/5";
            break;
        case 2:
            width_style = "w-3/5";
            break;
        case 3:
            width_style = "w-4/5";
            break;
        case 4:
            width_style = "w-5/5";
            break;
        default:
            width_style = "w-1/5";
            break;
    }
    return (
        <div className="mt-4 overflow-hidden rounded-full bg-gray-200 w-full">
            <div className={clsx(width_style, "h-2 rounded-full bg-cyan-400")}></div>
        </div>
    )
}
export default ProgressBar
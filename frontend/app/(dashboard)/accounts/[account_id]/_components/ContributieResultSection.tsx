import { Contribution } from '@/types/model';
import { clsx } from '@/utils/clsx';

const ContributieResultSection = ({ contributions }: { contributions?: Contribution[] }) => {
    const points = Array.from({ length: 60 }, () => 0);

    contributions?.forEach((contribute) => {
        const today = new Date();
        today.setHours(0, 0, 0, 0);
        const contributeDate = new Date(contribute.date);
        contributeDate.setHours(0, 0, 0, 0);
        const dateDiff = Math.floor((today.getTime() - contributeDate.getTime()) / (1000 * 60 * 60 * 24));
        if (dateDiff >= 0 && dateDiff < 90) {
            points[89 - dateDiff + 1] = contribute.points;
        }
    });
    function checkPoint(point: number): string {
        if (point === 0) {
            return "bg-gray-100";
        } else if (point >= 1 && point <= 3) {
            return "bg-cyan-50";
        } else if (point >= 4 && point <= 6) {
            return "bg-cyan-100";
        } else if (point >= 7 && point <= 9) {
            return "bg-cyan-150";
        } else if (point >= 10 && point <= 15) {
            return "bg-cyan-200";
        } else if (point >= 16) {
            return "bg-cyan-200";
        } else {
            return "bg-gray-100";
        }
    }
    return (
        <div className="w-full flex flex-col space-y-2 sm:w-[500px] items-center">
            <div className='grid grid-rows-4 grid-flow-col gap-[5px] overflow-x-scroll sm:w-[500px]'
            >{points.map(((point, index) => (
                <div key={index + "." + point} className={clsx("h-5 w-5 rounded-md", checkPoint(point),
                )}></div>
            )))
                }</div>
            <div className="text-slate-400">contributions</div>
        </div>
    )
}
export default ContributieResultSection
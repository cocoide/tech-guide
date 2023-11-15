import { ChildrenProps } from '@/types/props'
import Image from 'next/image'

const OnboardingLayout = ({ children }: ChildrenProps) => {
    return (
        <div className="flex lg:flex-row text-custom">
            <div className="hidden lg:flex flex-col space-y-3 lg:w-[50%]">
                <div className="text-2xl font-bold">Tech Guideへようこそ</div>
                <Image src={"/about/view.png"} alt='view' width={400} height={200} className="rounded-md h-[200px] w-[350px]  custom-border"></Image>
            </div>
            <div className='border-[0.5px] border-cyan-300 bg-cyan-50 rounded-xl shadow-sm m-10'
            >{children}</div>
        </div>
    )
}
export default OnboardingLayout
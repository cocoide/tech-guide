import { ChildrenProps } from '@/types/props'
import Image from 'next/image'

const OnboardingLayout = ({ children }: ChildrenProps) => {
    return (
        <div className="flex lg:flex-row text-custom justify-center w-full p-10 space-x-10">
            <div className="hidden lg:flex flex-col space-y-3 lg:w-[50%] items-center">
                <div className="text-2xl font-bold">Tech Guideへようこそ</div>
                <Image src={"/about/view.png"} alt='view' width={400} height={200} className="rounded-md h-[200px] w-[350px]  custom-border"></Image>
            </div>
            <div className='custom-boarder bg-gray-50 rounded-xl shadow-sm'
            >{children}</div>
        </div>
    )
}
export default OnboardingLayout
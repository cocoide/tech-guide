import HStack from '@/app/_components/elements/ui/HStack'
import WStack from '@/app/_components/elements/ui/WStack'

const LoaderArticleCard = () => {
    return (
        <HStack className='w-full space-y-5
            bg-gray-50/50 border-[0.5px] border-gray-200/70 dark:border-gray-400 dark:bg-black
            relative rounded-xl p-3'>
            <WStack centerY={true} className="space-x-5 w-full">
                <div className="rounded-full h-7 w-7 custom-loader"></div>
                <HStack className='space-y-2 w-full' centerY={true}>
                    <div className="rounded-full h-5 w-[250px] custom-loader"></div>
                    <div className="rounded-full h-5 w-[200px] custom-loader"></div>
                </HStack>
            </WStack>
            <div className="h-[150px] w-[270px] mx-auto bg-gray-100 dark:bg-gray-500 rounded-xl"></div>
            <WStack className='justify-end w-full'>
                <div className='h-6  rounded-xl w-[100px] custom-loader'></div>
            </WStack>
        </HStack>
    )
}
export default LoaderArticleCard
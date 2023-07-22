import { ChatBubbleOvalLeftEllipsisIcon, DocumentIcon, HandThumbDownIcon, HandThumbUpIcon, ShareIcon } from '@heroicons/react/24/outline'

const MobileFooter = () => {
    return (
        <div className="z-50 bg-white/70 sm:hidden fixed bottom-0 right-0 left-0  w-full py-2 px-5 flex flex-row items-center  justify-between"
        ><HandThumbUpIcon className='h-7 w-7 text-gray-500 hover:text-pink-300 hover:bg-pink-100 p-[1px] rounded-full' />
            <HandThumbDownIcon className='h-7 w-7 text-gray-500' />
            <ChatBubbleOvalLeftEllipsisIcon className='h-7 w-7 text-gray-500' />
            <DocumentIcon className='h-7 w-7 text-gray-500' />
            <ShareIcon className='h-7 w-7 text-gray-500' />
        </div>
    )
}
export default MobileFooter
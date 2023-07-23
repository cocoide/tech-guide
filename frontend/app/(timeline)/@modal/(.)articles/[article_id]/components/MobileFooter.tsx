import { ChatBubbleOvalLeftEllipsisIcon, FolderPlusIcon, HandThumbUpIcon, ShareIcon } from '@heroicons/react/24/outline'

const MobileFooter = () => {
    return (
        <div className="z-50 bg-white sm:hidden fixed bottom-2 left-1/2 transform -translate-x-1/2
        py-2 px-5 flex flex-row items-center  justify-center space-x-5
        ring-[0.5px] ring-gray-300 mx-auto rounded-full"
        ><HandThumbUpIcon className='h-7 w-7 text-gray-500 hover:text-pink-300 hover:bg-pink-100 p-[1px] rounded-full' />
            <ChatBubbleOvalLeftEllipsisIcon className='h-7 w-7 text-gray-500' />
            <FolderPlusIcon className='h-7 w-7 text-gray-500' />
            <ShareIcon className='h-7 w-7 text-gray-500' />
        </div>
    )
}
export default MobileFooter
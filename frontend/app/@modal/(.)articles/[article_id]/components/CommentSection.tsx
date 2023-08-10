import { ChatBubbleOvalLeftEllipsisIcon, ChevronUpDownIcon } from '@heroicons/react/24/outline';
const CommentSection = () => {
    return (
        <div className="border-y w-full p-2 flex flex-row justify-between items-center text-gray-500">
            <div className="flex flex-row items-center space-x-1 bg-gray-100 rounded-full p-1">
                <ChatBubbleOvalLeftEllipsisIcon className='h-5 w-5' />
                <div className="">コメント</div>
            </div>
            <ChevronUpDownIcon className='h-7 w-7' />
        </div>
    )
}
export default CommentSection

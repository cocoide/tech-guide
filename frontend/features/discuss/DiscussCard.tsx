import HStack from '@/app/_components/elements/ui/HStack'
import WStack from '@/app/_components/elements/ui/WStack'
import { Comment } from '@/types/model'
import Image from 'next/image'
import Link from 'next/link'

interface Props {
    comment: Comment
}
const DiscussCard = ({ comment }: Props) => {
    return (
        <WStack className="custom-text space-x-1 p-5  sm:pl-10 border-b-[0.5px] sm:border-b-0 sm:custom-boarder sm:rounded-xl  border-gray-200 dark:border-gray-400">
            <Image src={comment.account.avatar_url} alt={comment.account.display_name} width={100} height={100} className='w-10 h-10 rounded-full custom-border' />
            <HStack className='space-y-2'>
                <div className="font-medium">{comment.account.display_name}</div>
                <div className="font-medium text-sm">{comment.content}</div>
                <Link href={`/articles/${comment.article.id}`} className="overflow-hidden h-[180px] w-[300px] relative rounded-xl custom-border">
                    {/* eslint-disable-next-line @next/next/no-img-element */}
                    <img src={comment.article.thumbnail_url} alt={comment.article.title} className='min-h-[180px] w-[300px]' />
                </Link>
            </HStack>
        </WStack>
    )
}
export default DiscussCard
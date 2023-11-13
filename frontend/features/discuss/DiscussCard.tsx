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
        <WStack className="custom-text space-x-1">
            <Image src={comment.account.avatar_url} alt={comment.account.display_name} width={100} height={100} className='w-10 h-10 rounded-full' />
            <HStack className='space-y-2'>
                <div className="font-medium">{comment.account.display_name}</div>
                <div className="font-medium text-sm">{comment.content}</div>
                <Link href={`/articles/${comment.article.id}`} className="overflow-hidden h-[150px] w-[270px] relative rounded-xl custom-border">
                    {/* eslint-disable-next-line @next/next/no-img-element */}
                    <img src={comment.article.thumbnail_url} alt={comment.article.title} width={200} height={100} className='min-h-[150px] w-[280px] lg:w-[400px]' />
                </Link>
                <div className="">{comment.article.title}</div>
            </HStack>
        </WStack>
    )
}
export default DiscussCard
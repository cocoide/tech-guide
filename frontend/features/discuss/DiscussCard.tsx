import ArticleCard from '@/app/(home)/_components/ArticleCard'
import HStack from '@/app/_components/elements/ui/HStack'
import WStack from '@/app/_components/elements/ui/WStack'
import { Comment } from '@/types/model'
import Image from 'next/image'

interface Props {
    comment: Comment
}
const DiscussCard = ({ comment }: Props) => {
    return (
        <WStack className="custom-text">
            <Image src={comment.account.avatar_url} alt={comment.account.display_name} width={100} height={100} className='w-10 h-10 rounded-full' />
            <HStack className='space-y-3'>
                <div className="text-sm">{comment.account.display_name}</div>
                <div className="text-md font-bold">{comment.content}</div>
                <ArticleCard article={comment.article} />
            </HStack>
        </WStack>
    )
}
export default DiscussCard
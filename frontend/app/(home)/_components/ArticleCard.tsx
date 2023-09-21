import { Article } from '@/types/model'
import Link from 'next/link'
import ArticleOption from './ArticleOption'

interface Props {
    article: Article
    origin?: string
}

const ArticleCard = ({ article, origin }: Props) => {
    const queryParam = origin ? `?origin=${origin}` : ""
    return (
        <div className='bg-white dark:bg-black
        relative rounded-xl custom-shadow dark:shadow-gray-700 min-h-[30px]'>
            <Link href={`/articles/${article.id}${queryParam}`} className='hover:bg-slate-100 h-full w-full duaration-500'>
                <div className='flex flex-col space-y-[5px]'>
                    <div className="flex flex-row space-x-3">
                        <div className='text-slate-700 dark:text-white  mt-[12px] mx-[12px] mb-8'>{article.title}</div>
                    </div>

                    <div className='flex flex-row justify-between'>
                        <div className='flex flex-row space-x-3'>
                        </div>
                    </div>
                </div>
            </Link>
            <ArticleOption article={article} />
        </div>
    )
}
export default ArticleCard
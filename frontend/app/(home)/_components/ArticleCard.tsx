import { Article } from '@/types/model'
import Image from 'next/image'
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
        relative rounded-md shadow-sm dark:shadow-gray-700 custom-border min-h-[40px]'>
            {article.source.icon_url &&
                <Link href={`/sources/${article.source.id}`} className=''>
                <Image src={article.source.icon_url} alt={article.source.name} width={100} height={100}
                    className='rounded-full h-[24px] w-[24px] absolute top-[12px] left-[12px]' />
                               </Link>
            }
            <Link href={`/articles/${article.id}${queryParam}`} className='hover:bg-slate-100 h-full w-full duaration-500'>
                <div className='flex flex-col space-y-[5px]'>
                    <div className="flex flex-row space-x-3">
                        <div className='text-slate-700 dark:text-white  mt-[12px] ml-10 mb-7 pr-5'>{article.title}</div>
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
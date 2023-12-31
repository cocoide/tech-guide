import { Article } from '@/types/model'
import { NewspaperIcon } from '@heroicons/react/24/outline'
import Image from 'next/image'
import Link from 'next/link'
import ArticleOption from './ArticleOption'

interface Props {
    article: Article
    origin?: Number
}

const ArticleCard = ({ article, origin }: Props) => {
    const queryParam = origin ? `?origin=${origin}` : ""
    return (
        <div className='bg-gray-50/50 dark:bg-gray-900
        relative rounded-2xl min-h-[25px] custom-border shadow-sm flex flex-col'>
            <Link href={`/sources/${article.source.id}`}>
                <Image src={article.source.icon_url} alt={article.source.name} width={200} height={200} className='absolute top-3 left-3
                            h-7 w-7 rounded-full z-10' />
            </Link>
            <Link href={`/articles/${article.id}${queryParam}`} className='h-full w-full duaration-500'>
                <div className='flex flex-col p-4 space-y-2 h-full'>
                    <div className='text-slate-800 dark:text-white  font-bold min-h-[40px] pl-8 text-sm'>{article.title}</div>
                    <p className="h-auto w-full"></p>
                    <div className="overflow-hidden h-[150px] w-[270px] mx-auto relative flex flex-row justify-center rounded-md custom-border">

                        {article.thumbnail_url ?
                            // eslint-disable-next-line @next/next/no-img-element, jsx-a11y/alt-text
                            <img src={article.thumbnail_url} className='min-h-[160px] w-[280px] lg:w-[400px]' />
                            :
                            <div className="flex items-center justify-center h-[160px] w-[280px] lg:w-[400px] bg-slate-200 dark:slate-800"
                            ><NewspaperIcon className='h-10 w-10' /></div>
                        }
                    </div>
                </div>
            </Link>
            <ArticleOption article={article} />
        </div>
    )
}
export default ArticleCard



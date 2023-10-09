import { Article } from '@/types/model'
import { NewspaperIcon } from '@heroicons/react/24/outline'
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
        <div className='group bg-white dark:bg-gray-900
        relative rounded-2xl custom-shadow min-h-[25px] custom-border'>
            <Link href={`/articles/${article.id}${queryParam}`} className='hover:bg-slate-100 h-full w-full duaration-500'>
                <div className='flex flex-col p-4 space-y-2'>
                    <div className="flex  flex-row items-center justify-between">
                        <Link href={`/sources/${article.source.id}`} target="_blank">
                            <Image src={article.source.icon_url} alt={article.source.name} width={200} height={200} className='h-7 w-7 rounded-full z-10' />
                        </Link>

                        <a href={article.original_url} target="_blank" className="hidden group-hover:flex bg-slate-700  cutom-outline
                        text-white  rounded-xl shadow-sm p-[3px] text-sm animate-appear duration-100 z-10">元記事を読む</a>
                    </div>
                    <div className='text-slate-800 dark:text-white  font-bold text-md min-h-[60px]'>{article.title}</div>
                    <div className="overflow-hidden h-[200px] relative flex flex-row justify-center">
                        {article.thumbnail_url ?
                            // eslint-disable-next-line @next/next/no-img-element, jsx-a11y/alt-text
                            <img src={article.thumbnail_url} className='h-40 w-100 rounded-xl' />
                            :
                            <div className="flex items-center justify-center h-[160px] w-[400px] bg-slate-200 dark:slate-800 rounded-xl"
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
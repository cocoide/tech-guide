import { Article } from '@/types/model'
import { NewspaperIcon } from '@heroicons/react/24/outline'
import Link from 'next/link'
import ArticleOption from './ArticleOption'

interface Props {
    article: Article
    origin?: string
}

const ArticleCard = ({ article, origin }: Props) => {
    const queryParam = origin ? `?origin=${origin}` : ""
    return (
        <div className='bg-white dark:bg-black overflow-hidden
        relative rounded-xl custom-shadow dark:shadow-slate-600 min-h-[20px]'>
            <a href={article.original_url} target="_blank" className="bg-slate-400/60 backdrop-blur-[5px] absolute top-3 right-3 cutom-outline
                        text-white p-[3px] rounded-xl shadow-sm z-10">元記事を読む</a>
            <Link href={`/articles/${article.id}${queryParam}`} className='hover:bg-slate-100 h-full w-full duaration-500'>
                <div className='flex flex-col space-y-[5px]'>
                    <div className="overflow-y-hidden h-[200px] relative">
                        {article.thumbnail_url ?
                            // eslint-disable-next-line @next/next/no-img-element, jsx-a11y/alt-text
                            <img src={article.thumbnail_url} className='h-50 w-100' />
                            :
                            <div className="flex items-center justify-center h-full w-full bg-slate-200 dark:slate-800"
                            ><NewspaperIcon className='h-10 w-10' /></div>
                        }
                    </div>
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
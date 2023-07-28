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
        <div className='relative'>
            {article.source.icon_url &&
                <Image src={article.source.icon_url} alt={article.source.name} width={100} height={100}
                    className='rounded-full h-[24px] w-[24px] absolute top-[12px] left-[12px]' />
            }
            <Link href={`/articles/${article.id}${queryParam}`} className='hover:bg-slate-100 h-full w-full duaration-500'>
                <div className='flex flex-col space-y-[5px]'>

                    <div className="flex flex-row space-x-3">
                        <div className='text-slate-700 mt-[12px] ml-10 mb-7'>{article.title}</div>
                        </div>
                        <div className="text-slate-500 mr-auto">{article.topics?.slice(0, 1).map(topic =>
                            (<div key={topic.id} className='flex flex-row items-center text-[10px] ring-1 rounded-md ring-slate-300 p-[2px]'>{topic.name}</div>)
                        )}</div>

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
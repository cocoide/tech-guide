import { Article } from '@/types/model'
import { ArrowTopRightOnSquareIcon, NewspaperIcon } from '@heroicons/react/24/outline'
import Image from 'next/image'
import Link from 'next/link'
import ArticleOption from './ArticleOption'

interface Props {
    article: Article
    origin?: string
}

const ArticleCard = ({ article, origin }: Props) => {
    const queryParam = origin ? `?origin=${origin}` : ""
    const domain = extractDomain(article.original_url)
    return (
        <div className='bg-white dark:bg-gray-900
        relative rounded-2xl custom-shadow min-h-[25px] custom-border'>
            <Link href={`/sources/${article.source.id}`}>
                <Image src={article.source.icon_url} alt={article.source.name} width={200} height={200} className='absolute top-3 left-3
                            h-7 w-7 rounded-full z-10' />
            </Link>
            <a href={article.original_url} target="_blank" className="flex bg-gray-100 dark:bg-gray-700 cutom-outline absolute top-3 right-3
                        text-gray-500 dark:text-gray-200  rounded-xl shadow-sm p-[4px] text-sm animate-appear duration-100 z-10 custom-badge">
                <div className="text-sm">{domain}</div>
                <ArrowTopRightOnSquareIcon className="h-4 w-4"></ArrowTopRightOnSquareIcon>
            </a>
            <Link href={`/articles/${article.id}${queryParam}`} className='hover:bg-slate-100 h-full w-full duaration-500'>
                <div className='flex flex-col p-4 space-y-2'>
                    <p className="w-[100%] h-6"></p>
                    <div className='text-slate-800 dark:text-white  font-bold text-md min-h-[60px]'>{article.title}</div>
                    <div className="overflow-hidden h-[200px] relative flex flex-row justify-center">
                        {article.thumbnail_url ?
                            // eslint-disable-next-line @next/next/no-img-element, jsx-a11y/alt-text
                            <img src={article.thumbnail_url} className='h-40 w-70 lg:w-100 rounded-xl' />
                            :
                            <div className="flex items-center justify-center h-[160px] w-[280px] lg:w-[400px] bg-slate-200 dark:slate-800 rounded-xl"
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

function extractDomain(url: string): string | null {
    try {
        const parsedUrl = new URL(url);
        return parsedUrl.hostname;
    } catch (e) {
        console.error("Invalid URL provided:", e);
        return null;
    }
}

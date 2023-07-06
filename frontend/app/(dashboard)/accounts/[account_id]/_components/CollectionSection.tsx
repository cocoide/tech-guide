import { Collection } from '@/types/model'
import Link from 'next/link'

interface Props {
    collections: Collection[]
    account_id: string
}
const CollectionSection = ({ collections, account_id }: Props) => {
    return (
        <div className="">
            <div className="gap-5 grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 p-5 place-items-center">
                {collections?.map(c => (
                    <Link href={`/accounts/${account_id}/collections/${c.id}/`} key={c.id} className="flex flex-col space-y-3 w-auto">
                        {c.articles[0]?.thumbnail_url ?
                            // eslint-disable-next-line @next/next/no-img-element
                            <img src={c.articles[0].thumbnail_url} alt={c.articles[0].title} width={100} className='rounded-md w-auto shadow-md' />
                            :
                            <div className="rounded-md w-[200px] h-[100px] bg-slate-200 shadow-[3px]"></div>
                        }
                        <div className="flex flex-row items-center justify-between">
                            <div className="text-md">{c.name}</div>
                            <div className="text-gray-400">{c.articles.length} picks</div>
                        </div>
                    </Link>
                ))}
            </div>
        </div>
    )
}
export default CollectionSection
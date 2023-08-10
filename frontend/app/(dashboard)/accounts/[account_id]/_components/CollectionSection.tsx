import { Collection } from '@/types/model'
import Link from 'next/link'

interface Props {
    collections: Collection[]
    account_id: string
}
const CollectionSection = ({ collections, account_id }: Props) => {
    return (
        <div className="flex flex-col space-y-3 p-3 items-center px-[5%] md:px-[7%] lg:px-[15%]">
                {collections?.map(c => (
                    <Link href={`/accounts/${account_id}/collections/${c.id}/`} key={c.name + c.id}
                        className="rounded-md flex flex-row justify-between  bg-gray-50 ring-1 ring-gray-100
                         shadow-sm w-full h-[120px] p-1 items-center
                        ">
                        <div className="flex flex-row items-center justify-center space-x-3 w-full">
                            <div className="text-md">{c.name}</div>
                        </div>
                        {/* image section */}
                        <div className="overflow-hidden h-full w-[200px] flex flex-row space-x-1">
                            {c.articles[0]?.thumbnail_url ?
                                // eslint-disable-next-line @next/next/no-img-element
                                <img src={c.articles[0].thumbnail_url} alt={c.articles[0].title} width={100}
                                    className='w-[50%]  shadow-md overflow-hidden' />
                                :
                                <div className="w-[50%] h-full bg-slate-200 shadow-[3px]"></div>
                            }
                        {c.articles[0]?.thumbnail_url ?
                            // eslint-disable-next-line @next/next/no-img-element
                                <img src={c.articles[0].thumbnail_url} alt={c.articles[0].title} width={100}
                                    className=' w-[30%]  overflow-hidden shadow-md' />
                            :
                                <div className="w-[30%] h-full bg-slate-200 shadow-[3px]"></div>
                            }
                            <div className="w-[20%] h-full bg-slate-200 shadow-[3px]"></div>
                        </div>
                    </Link>
                ))}
        </div>
    )
}
export default CollectionSection
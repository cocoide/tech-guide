import { Collection } from '@/types/model'
import { NewspaperIcon } from '@heroicons/react/24/outline'
import Link from 'next/link'

interface Props {
    collections: Collection[]
    account_id: string
}
export default async function CollectionSection({ collections, account_id }: Props) {
    return (
        <div className="flex flex-col space-y-3 p-3 items-center px-[5%] md:px-[7%] lg:px-[15%]">
                {collections?.map(c => (
                    <Link href={`/accounts/${account_id}/collections/${c.id}/`} key={c.name + c.id}
                        className="rounded-md flex flex-row justify-between  custom-border
                         shadow-sm w-full h-[120px] p-2 items-center
                        ">
                        <div className="flex flex-col space-y-3 w-full h-full p-2">
                            <div className="text-md text-gray-600 w-full custom-badge">
                                <div> {c.name}</div>
                                <div className="bg-cyan-50 text-cyan-300 border-cyan-300 border-[0.5px] rounded-full h-5 w-5 flex justify-center items-center"
                                >{collections.length}</div>
                            </div>
                        </div>
                        {/* image section */}
                            {c.articles[0]?.thumbnail_url ?
                            // eslint-disable-next-line @next/next/no-img-element
                            <img src={c.articles[0].thumbnail_url} alt={c.articles[0].title} width={200}
                                className='rounded-md w-[50%] h-[100px] shadow-md overflow-hidden' />
                                :
                            <div className="flex items-center justify-center rounded-md w-[50%] h-[100px] bg-gray-100 shadow-[2px] custom-border">
                                <NewspaperIcon className="h-7 w-7 text-gray-500" />
                            </div>
                        }
                    </Link>
                ))}
        </div>
    )
}

import { Collection } from '@/app/_models'

const CollectionSection = ({ collections }: { collections: Collection[] }) => {
    return (
        <div className="">
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 p-5 place-items-center">
                {collections?.map(c => (
                    <div key={c.id} className="flex flex-col space-y-3 w-[200px]">
                        {c.articles[0]?.thumbnail_url ?
                            // eslint-disable-next-line @next/next/no-img-element
                            <img src={c.articles[0].thumbnail_url} alt={c.articles[0].title} width={200} className='rounded-md w-[200px] h-[100px] shadow-md' />
                            :
                            <div className="rounded-md w-[200px] h-[100px] bg-slate-200 shadow-[3px]"></div>
                        }
                        <div className="flex flex-row items-center justify-between">
                            <div className="text-md">{c.name}</div>
                            <div className="text-gray-400">{c.articles.length} picks</div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )
}
export default CollectionSection
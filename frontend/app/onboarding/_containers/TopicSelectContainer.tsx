import { topicAPI } from '@/app/(dashboard)/settings/_functions/topic'
import HStack from '@/app/_components/elements/ui/HStack'

export default async function TopicSelectContainer() {
    const { data: categories } = await topicAPI.GetAllCategories()
    return (
        <HStack className='custom-text space-y-3 items-center w-full p-5 md:p-10'>
            <div className="">フォローするカテゴリを選ぶ</div>
            <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-5">
                {categories?.map(category => (
                    <div key={category.id} className="aspect-square w-full custom-border bg-cyan-50 text-cyan-300 rounded-xl flex items-center justify-center p-5"
                    >{category.name}</div>
                ))}
            </div>
        </HStack>
    )
}
import { topicAPI } from '@/app/(dashboard)/settings/_functions/topic';
import HStack from '@/app/_components/elements/ui/HStack';
import SelectTopicGroup from '../_components/SelectTopicGroup';

export default async function TopicSelectContainer() {
    const { data: categories } = await topicAPI.GetAllCategories()
    return (
        <HStack className='custom-text space-y-3 items-center w-full p-5 md:p-10'>
            <div className="">フォローするカテゴリを選ぶ</div>
            <SelectTopicGroup categories={categories} />
        </HStack>
    )
}
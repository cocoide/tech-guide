"use client"

import { Category, Topic } from '@/types/model';
import { clsx } from '@/utils/clsx';

const SelectTopicGroup = ({ categories }: { categories?: Category[] }) => {
    const followingTopics: Record<number, Topic[]> = {};
    function handleFollowCategory(category: Category) {
        followingTopics[category.id] = category.topics
    }
    return (
        <div className="grid grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-5">
            {categories?.map(category => (
                <button key={category.id} onClick={() => handleFollowCategory(category)}
                    className={clsx(followingTopics[category.id] ? "bg-cyan-300 text-white" : "bg-cyan-50 text-cyan-300 ", "relative custom-border aspect-square w-full rounded-xl flex items-center justify-center p-5")}
                >{category.name}
                    {followingTopics[category.id].length > 0}
                    <div className="absolute bottom-3 right-3">{followingTopics[category.id].length}件フォロー中</div>
                </button>
            ))}
        </div>
    )
}
export default SelectTopicGroup
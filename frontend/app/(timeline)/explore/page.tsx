import { topicAPI } from './_functions/topic';

export default async function ExplorePage() {
    const { data: topics } = await topicAPI.GetAllTopics()
    return (
        <div className="w-full p-10 flex flex-wrap gap-3">
            {topics?.map(topic =>
                <div key={topic.id} className="flex-shrink-0 p-1 bg-slate-100 text-slate-600 rounded-xl"
                >{topic.name}</div>
            )}
        </div>
    )
}
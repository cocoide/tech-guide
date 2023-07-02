import { authAPI } from '@/app/_functions/auth'
import { topicAPI } from '../_functions/topic'
import FollowTopicsSection from './components/FollowTopicsSection'
import RecommendTopicsSection from './components/RecommendTopicsSection'

export default async function TopicsSettingPage() {
    const session = await authAPI.GetAuthSession()
    const { data: follow_topics } = await topicAPI.GetFollowingTopics(session?.token)
    const { data: topics } = await topicAPI.GetAllTopics()
    const unfollow_topics = topics?.filter(topic => {
        return !follow_topics?.some(followTopic => followTopic.id === topic.id);
    });
    return (
        <div className="flex flex-col space-y-3 justify-center w-full">
            {follow_topics ?
                <FollowTopicsSection follow_topics={follow_topics} />
                :
                <div className="text-slate-500">フォロー中のトピックがありません</div>
            }
            {unfollow_topics &&
                <RecommendTopicsSection unfollow_topics={unfollow_topics} />
            }
        </div>
    )
}
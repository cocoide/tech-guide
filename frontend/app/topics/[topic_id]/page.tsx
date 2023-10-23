import { topicAPI } from '@/app/_functions/topic'
import { authServerFunc } from '@/app/_server_functions/auth'
import TopicPageContent from './TopicPageContent'

interface Props {
    params: {
        topic_id: string
    }
}
export default async function TopicPage({ params }: Props) {
    const { token } = await authServerFunc.GetAuth()
    const { data: topic } = await topicAPI.GetTopicData(Number(params.topic_id))
    const { data: isFollowing } = await topicAPI.CheckFollow(Number(params.topic_id), token)
    return (
        <TopicPageContent params={params} topic={topic} isFollowing={isFollowing} />
    )
}

import { sourceAPI } from '@/app/_functions/source'
import { authServerFunc } from '@/app/_server_functions/auth'
import SourcePageContent from './SourcePageContent'

interface Props {
    params: {
        source_id: string
    }
}
export default async function SourcePage({ params }: Props) {
    const { token } = await authServerFunc.GetAuth()
    const { data: source } = await sourceAPI.GetSourceData(Number(params.source_id))
    const { data: isFollowing, error } = await sourceAPI.CheckFollow(Number(params.source_id), token)
    return (
        <SourcePageContent params={params} source={source} isFollowing={isFollowing} />
    )
}

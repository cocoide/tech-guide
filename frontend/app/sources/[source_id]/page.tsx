import { authAPI } from '@/app/_functions/auth'
import { sourceAPI } from '@/app/_functions/source'
import SourcePageContent from './SourcePageContent'

interface Props {
    params: {
        source_id: string
    }
}
export default async function SourcePage({ params }: Props) {
    const { token } = await authAPI.GetAuthSession()
    const { data: source } = await sourceAPI.GetSourceData(Number(params.source_id))
    const { data: isFollowing, error } = await sourceAPI.CheckFollow(Number(params.source_id), token)
    return (
        <SourcePageContent params={params} source={source} isFollowing={isFollowing} />
    )
}

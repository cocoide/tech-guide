import ContributieResultSection from './_components/ContributieResultSection';
import { activityAPI } from './_functions/activity';

interface Props extends AccountParams {
}

export default async function AccountPage({ params }: Props) {
    const { data: contributions } = await activityAPI.GetContributions(params.account_id)
    return (
        <div className="p-5 w-full flex flex-col jutity-center items-center">
            <ContributieResultSection contributions={contributions} />
        </div>
)
}
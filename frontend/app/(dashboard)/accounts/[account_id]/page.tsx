import CollectionSection from './_components/CollectionSection'
import ProfileSection from './_components/ProfileSection'
import { accountAPI } from './_functions/account'
import { collectionAPI } from './_functions/collection'

interface Props extends AccountParams {
}

export default async function page({ params }: Props) {
    const accountId = Number(params.account_id)
    const { data: collections } = await collectionAPI.getCollections(accountId)
    const { data: profile } = await accountAPI.getProfile(accountId)
    return (
        <div className="flex flex-col space-y-3">
            {profile &&
                <ProfileSection profile={profile} />
            }
            {collections &&
                <CollectionSection collections={collections} />
            }
        </div>
    )
}
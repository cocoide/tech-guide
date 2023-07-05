import CollectionSection from '../_components/CollectionSection'
import { collectionAPI } from '../_functions/collection'

interface Props extends AccountParams {
}

export default async function page({ params }: Props) {
    const accountId = Number(params.account_id)
    const { data: collections } = await collectionAPI.getCollections(accountId)
    return (
        <div className="flex flex-col space-y-3">
            {collections &&
                <CollectionSection collections={collections} />
            }
        </div>
    )
}
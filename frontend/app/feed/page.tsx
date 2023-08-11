import { authAPI } from '../_functions/auth';
import FeedSection from './_components/FeedSection';


export default async function FeedPage() {
    const { token } = await authAPI.GetAuthSession()
    return (
        <FeedSection token={token} />
    )
}
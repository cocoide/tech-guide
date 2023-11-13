import DiscussItems from './DiscussItems';
import { getLatestComment } from './_function/comment';

export default async function DiscussContainer() {
    const comments = await getLatestComment()
    return <DiscussItems Discusss={comments} fetchFunc={getLatestComment} />
}
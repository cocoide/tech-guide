import { api } from '@/app/_functions/API';
import { Comment } from '@/types/model';

export async function getLatestComment(page = 1) {
    "use server"

    let comments: Comment[] = []
    const { data } = await api.get<Comment[]>(`/comment?page=${page}`, 60 * 60 * 24)
    if (data) {
        comments = data
    }
    return comments
}
import { api } from '@/app/_functions/API';
import { Source, Topic } from '@/types/model';

export const popularAPI = {
    async GetPopularTopics() {
        return await api.get<Topic[]>("/topic/popular", "no-store")
    },
    async GetPopularSources() {
        return await api.get<Source[]>("/source/popular", "no-store")
    }
}
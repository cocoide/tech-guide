import { Topic } from '@/types/model';
import { api } from './API';

export const topicAPI={
    async GetTopicData(topicID: number){
        return api.get<Topic>(`/topic/${topicID}`,"no-store");
    },
    async CheckFollow(topicID: number,token?:string){
        return api.get<boolean>(`/account/topic/follow/${topicID}`,"no-store",token);
    },
    async DoFollow(topicID: number,token?:string){
        return api.put(`/account/topic/follow/${topicID}`,undefined,token);
    },
    async UnFollow(topicID: number,token?:string){
        return api.del(`/account/topic/follow/${topicID}`,token);
    },
}
import { Source } from '@/types/model';
import { api } from './API';

export const sourceAPI ={
    async GetSourceData(sourceID: number){
        return api.get<Source>(`/source/${sourceID}`,"no-store");
    },
    async CheckFollow(sourceID: number,token?:string){
        return api.get<boolean>(`/account/source/follow/${sourceID}`,"no-store",token);
    },
    async DoFollow(sourceID: number,token?:string){
        return api.put(`/account/source/follow/${sourceID}`,token);
    },
    async UnFollow(sourceID: number,token?:string){
        return api.del(`/account/source/follow/${sourceID}`,token);
    },
}
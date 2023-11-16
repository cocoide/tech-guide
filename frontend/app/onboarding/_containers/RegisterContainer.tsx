"use client"

import { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';

import CircleLoading from '@/app/_components/animations/CircleLoading';
import HStack from '@/app/_components/elements/ui/HStack';
import { api } from '@/app/_functions/API';
import Image from 'next/image';
import { useRouter } from 'next/navigation';

type Inputs = {
    display_name: string
    avatar_url: string
}
interface Props {
    sessionId?: string
    display_name?: string
    avatar_url?: string
}
const RegisterContainer = ({ sessionId, avatar_url, display_name }: Props) => {
    const router = useRouter()
    const [isLoading, setIsLoading] = useState(false);
    const {
        register,
        handleSubmit,
        setValue,
        watch,
    } = useForm<Inputs>();
    const onSubmit = async (data: Inputs) => {
        setIsLoading(true);
        const { ok } = await api.pos("/onboarding/register", data, undefined, undefined, sessionId);
        if (ok) {
            router.push(`/onboarding`);
        } else {
            router.push("/onboarding?status=error");
        }
        setIsLoading(false);
    };
    useEffect(() => {
        if (avatar_url) {
            setValue("avatar_url", avatar_url);
        }
        if (display_name) {
            setValue("display_name", display_name);
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])
    return (
        <HStack className='custom-text space-y-3 items-center w-full p-10 h-'>
            <>{isLoading && (
                <div className="z-10 bg-white/30 backdrop-blur-[2px]  fixed inset-0 flex justify-center items-center">
                    <CircleLoading />
                </div>
            )}</>
            <div className="text-xl">プロフィールを登録</div>
            <form onSubmit={handleSubmit(onSubmit)} className='flex flex-col items-center w-full space-y-3'>
                <div className="text-sm">表示名</div>
                <input {...register('display_name', { "required": true })} className='custom-boarder bg-slate-50 rounded-xl p-[3px] w-[300px]' />
                <div className="text-sm">アイコン</div>
                {watch("avatar_url") && watch("avatar_url").length > 0 ?
                    <Image src={watch("avatar_url")} alt='avatar' width={100} height={100} className='rounded-md h-[100px] w-[100px]' />
                    :
                    <div className="custom-loader rounded-md h-[100px] w-[100px]"></div>
                }
                <button className="bg-cyan-300 text-white p-[6px] rounded-xl shadow-sm">登録完了</button>
            </form>
        </HStack>
    )
}

export default RegisterContainer
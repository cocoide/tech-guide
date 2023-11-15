"use client"

import { useState } from 'react';
import { useForm } from 'react-hook-form';

import CircleLoading from '@/app/_components/animations/CircleLoading';
import HStack from '@/app/_components/elements/ui/HStack';
import { api } from '@/app/_functions/API';
import { useRouter } from 'next/navigation';

type Inputs = {
    display_name: string
    avatar_url: string
}
export default function RegisterContainer() {
    const router = useRouter()
    const [isLoading, setIsLoading] = useState(false);
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<Inputs>();
    const onSubmit = async (data: Inputs) => {
        setIsLoading(true);
        const { ok } = await api.pos("/auth/signup", data)
        if (ok) {
            router.push(`/onboarding?step=3`);
        } else {
            router.push("/onboarding?step=2&status=error");
        }
        setIsLoading(false);
    };
    return (
        <HStack className='custom-text space-y-3 items-center w-full'>
            {isLoading && (
                <div className="z-10 bg-white/30 backdrop-blur-[2px]  fixed inset-0 flex justify-center items-center">
                    <CircleLoading />
                </div>
            )}
            <div className="text-xl">プロフィールを登録</div>
            <form onSubmit={handleSubmit(onSubmit)}>
            <input {...register('display_name')} />
            <input {...register('avatar_url')} />
                <button className="bg-cyan-300 text-white p-[6px] rounded-xl text-sm">登録</button>
        </form>
        </HStack>
    )
}
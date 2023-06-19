'use client'

import CircleLoading from '@/app/_components/animations/CircleLoading';
import { LockClosedIcon } from '@heroicons/react/24/outline';
import { signIn, useSession } from 'next-auth/react';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import { useForm } from "react-hook-form";
import { toast } from 'react-hot-toast';
import { LoginRequest } from "../_function/auth";

export default function LoginPage() {
    const { register, handleSubmit } = useForm<LoginRequest>();
    const [isLoading, setIsLoading] = useState(false);
    const router = useRouter()
    const { data: session } = useSession()
    useEffect(() => {
        if (session?.user.uid) {
            router.push(`accounts/${session?.user.uid}`)
        }
    }, [router, session?.user])
    const onSubmit = async (body: LoginRequest) => {
        setIsLoading(true)
        const { email, password } = body
        try {
            const res = await signIn("credentials", {
                redirect: false,
                email,
                password,
            })
            if (res?.error != null) {
                setIsLoading(false)
                return toast.error("passwordまたはemailが間違ってます")
            }
        } catch (error) {
            console.log(error)
            setIsLoading(false)
            return toast.error("エラーが発生")
        }
        toast.success("ログインに成功")
        return router.push(`/accounts/${session?.user.uid}`)
    };

    return (
        <>
            {isLoading && <div className="z-10 bg-white/30 backdrop-blur-[2px]  fixed inset-0 flex justify-center items-center"
            ><CircleLoading /></div>}
            <div className="flex justify-center w-full">
                <div className='min-h-screen space-y-5 items-center w-[400px] md:w-[500px]'>
                    <form onSubmit={handleSubmit(onSubmit)}
                        className='flex flex-col items-center space-y-10 bg-white rounded-xl shadow-[5px] p-7 w-[100%] mt-20'>
                        <div className="flex flex-col w-[100%]">
                            <div className="text-sm text-slate-500">メールアドレス</div>
                            <input {...register("email", { required: true })}
                                className='bg-slate-50 ring-1 ring-gray-200 p-[5px] rounded-md w-[100%]' />
                        </div>
                        <div className="flex flex-col w-full">
                            <div className="text-sm text-slate-500">パスワード</div>
                            <input type="password" {...register("password", { required: true })}
                                className='bg-slate-50 ring-1 ring-gray-200 p-[5px] rounded-md' />
                        </div>
                        <button type="submit"
                            className='text-cyan-300 ring-1 ring-cyan-300  p-[5px] rounded-md w-full flex items-center 
                        justify-center space-x-2'
                        ><LockClosedIcon className='h-5 w-5 text-cyan-300' />
                            <div className=""> Login</div>
                        </button>
                    </form>
                </div>
            </div>
        </>
    );
};

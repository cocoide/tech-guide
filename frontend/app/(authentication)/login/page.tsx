'use client'

import CircleLoading from '@/app/_components/animations/CircleLoading';
import { signIn, useSession } from 'next-auth/react';
import Image from 'next/image';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import { toast } from 'react-hot-toast';

export default function LoginPage() {
    const [isLoading, setIsLoading] = useState(false);
    const router = useRouter()
    const { data: session } = useSession()
    useEffect(() => {
        if (session?.user.uid) {
            router.push(`accounts/${session?.user.uid}`)
        }
    }, [router, session?.user])
    async function signInWithGoogle() {
        setIsLoading(true)
        try {
            await signIn("google", {
                callbackUrl: window.location.href,
            })
        } catch (error) {
            toast.error("エラーが発生")
        }
    }

    return (
        <>
            {isLoading && <div className="z-10 bg-white/30 backdrop-blur-[2px]  fixed inset-0 flex justify-center items-center"
            ><CircleLoading /></div>}
            {session?.user.image}
            <div className="flex justify-center w-full">
                <div className='min-h-screen space-y-5 items-center w-[400px] md:w-[500px]'>
                    <form className='flex flex-col items-center space-y-10 bg-white rounded-xl shadow-[5px] p-7 w-[100%] mt-20'>
                        <button className='ring-1 ring-gray-200 rounded-md p-2 text-gray-700 flex flex-row items-center space-x-2' type="button" onClick={signInWithGoogle}>
                            <Image src={"/google.png"} alt='google' width={50} height={50} className='h-5 w-5' />
                            <div className="">
                                Googleアカウントでログイン
                            </div>
                        </button>
                    </form>
                </div>
            </div>
        </>
    );
};

"use client"
import { loginDialogAtom } from '@/stores/dialog'
import { signIn } from 'next-auth/react'
import Image from 'next/image'

import { useSetAtom } from 'jotai'
import { toast } from 'react-hot-toast'
import CustomDialog from '../elements/CustomDialog'

const LoginDialog = () => {
    const setLoginDialog = useSetAtom(loginDialogAtom)
    async function signInWithGoogle() {
        toast.loading("ログイン中...");
        try {
            await signIn("google", {
                callbackUrl: window.location.href,
            })
        } catch (error) {
            toast.error("エラーが発生")
        }
        setLoginDialog(false)
        toast.dismiss();
        toast.success("ログイン完了");
    }
    return (
        <CustomDialog
            openAtom={loginDialogAtom}
            layout='my-[150px] bg-white z-50 sm:mx-[15%] md:mx-[20%] lg:mx-[25%] sm:rounded-xl'
            content={
                <div className='flex flex-col justify-items-stretch space-y-5'>
                    <div className="text-gray-400 w-[300px] bg-cyan-50 p-3 rounded-xl">
                        『Tech Guide』にログインすると
                        <p></p>
                        投稿機能、保存機能、おすすめ機能などがご利用可能になります。
                    </div>
                    <button className='ring-1 ring-gray-200 rounded-md p-2 text-gray-700 flex flex-row items-center justify-center space-x-2' type="button"
                        onClick={signInWithGoogle}>
                        <Image src={"/google.png"} alt='google' width={50} height={50} className='h-5 w-5' />
                        <div className="text-center">
                            Googleアカウントでログイン
                        </div>
                    </button>
                </div>
            }
        />
    )
}
export default LoginDialog
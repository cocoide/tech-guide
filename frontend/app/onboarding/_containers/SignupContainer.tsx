"use client"

import HStack from '@/app/_components/elements/ui/HStack';
import Image from 'next/image';
import { useRouter } from 'next/navigation';
import toast from 'react-hot-toast';

const SignupContainer = () => {
    const router = useRouter()
    function handleSignup(type: "google" | "github") {
        toast.loading("ログイン中...");
        try {
            router.push(process.env.NEXT_PUBLIC_API_BASE_URL + `/oauth/login?type=${type}`)
        } catch (error) {
            toast.error("エラーが発生")
        }
    }

    return (
        <div className="flex lg:flex-row flex-col text-custom lg:justify-center items-center w-full p-10 lg:space-x-10 space-y-10">
            <div className="flex flex-col space-y-3 lg:w-[50%] items-center">
                <div className="text-2xl font-bold">Tech Guideへようこそ</div>
                <Image src={"/about/view.png"} alt='view' width={400} height={200} className="rounded-md h-[200px] w-[350px]  custom-border"></Image>
            </div>
            <div className='custom-boarder bg-gray-50 rounded-xl shadow-sm'
            >
                <HStack className="w-full items-center space-y-5 custom-text p-5">
            <div className="">TechGuideにサインアップ</div>
            <button className='ring-1 ring-gray-200 w-[300px] rounded-md p-2 text-gray-700 flex flex-row items-center space-x-3 bg-white' type="button"
                onClick={() => handleSignup("google")}>
                <Image src={"/google.png"} alt='google' width={50} height={50} className='h-5 w-5' />
                <div className="text-center">
                    Googleアカウントで登録
                </div>
            </button>
            <button className='ring-1 ring-gray-200 w-[300px] rounded-md p-2 text-gray-700 flex flex-row items-center space-x-3 bg-white' type="button"
                onClick={() => handleSignup("google")}>
                <Image src={"/github.svg"} alt='github' width={50} height={50} className='h-5 w-5' />
                <div className="text-center">
                    Githubアカウントで登録
                </div>
            </button>
            <HStack className="list-disc text-gray-400 w-[300px] bg-gray-50 p-3 rounded-xl custom-border space-y-2">
                <div className="font-bold text-gray-500">サインアップ後に楽しめる機能</div>
                <li>カスタムフィード</li>
                <li>投稿機能</li>
                <li>コントリビューション機能</li>
            </HStack>
        </HStack>
            </div>
        </div>
    )
}
export default SignupContainer
import HStack from '@/app/_components/elements/ui/HStack'
import Image from 'next/image'

const SignupContainer = () => {
    function handleSignup(oauthType: "google" | "github") {

    }

    return (
        <HStack className="w-full justify-center space-y-5 custom-text">
            <div className="">TechGuideにサインアップ</div>
            <button className='ring-1 ring-gray-200 rounded-md p-2 text-gray-700 custom-badge' type="button"
                onClick={() => handleSignup("google")}>
                <Image src={"/google.png"} alt='google' width={50} height={50} className='h-5 w-5' />
                <div className="text-center">
                    Googleアカウントで登録
                </div>
            </button>
            <button className='ring-1 ring-gray-200 rounded-md p-2 text-gray-700 custom-badge' type="button"
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
    )
}
export default SignupContainer
'use client'

import { useRouter } from 'next/navigation'
import { useAuth } from "@/hooks/useAuth";

export default function SettingPageLayout({ children, }: { children: React.ReactNode }) {
    const router = useRouter()
    const { status } = useAuth()
    if (status === "unauthenticated") {
        router.push("/")
    }
    return (
        <div className='p-5 w-full md:w-[700px] flex flex-col justify-center items-center mx-auto space-y-3'>
            <div className="w-full text-start text-xl text-slate-500 font-bold">設定</div>
            {children}
        </div>
    )
}
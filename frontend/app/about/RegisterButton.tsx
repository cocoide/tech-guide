"use client"

import { useAuth } from '@/hooks/useAuth'
import { useRouter } from 'next/navigation'

const RegisterButton = () => {
    const router = useRouter()
    const { status } = useAuth()
    function handleRegister() {
        if (status === "authenticated") {
            router.push("/")
        }
        if (status === "unauthenticated") {
            router.push("/onboarding")
        }
    }
    return (
        <button onClick={handleRegister} className="text-white bg-cyan-300 p-[7px] rounded-xl text-md shadow-sm"
        >はじめる</button>
    )
}

export default RegisterButton
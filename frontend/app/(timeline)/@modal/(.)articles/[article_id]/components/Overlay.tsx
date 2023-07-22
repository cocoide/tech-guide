"use client"

import { useRouter } from 'next/navigation'

const Overlay = () => {
    const router = useRouter()
    function handleClose() {
        router.back()
    }
    return (
        <button onClick={handleClose}
            className="z-30 bg-black/40  fixed inset-0 backdrop-blur-[3px] animate-appear"></button>
    )
}
export default Overlay
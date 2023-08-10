"use client"

import { XMarkIcon } from '@heroicons/react/24/outline'
import { useRouter } from 'next/navigation'

const CloseButton = () => {
    const router = useRouter()
    function handleClose() {
        router.back()
    }
    return (
        <button onClick={handleClose}>
            <XMarkIcon className='h-7 w-7 text-gray-600 p-[3px] hover:bg-gray-200 duration-500 rounded-md' />
        </button>
    )
}
export default CloseButton
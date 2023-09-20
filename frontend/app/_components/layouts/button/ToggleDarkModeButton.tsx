"use client"

import { useCacheDarkMode } from '@/hooks/useCacheDarkMode'
import { MoonIcon, SunIcon } from '@heroicons/react/24/outline'

const ToggleDarkModeButton = () => {
    const { isDarkMode, toggle } = useCacheDarkMode()
    return (
        <button onClick={() => toggle(!isDarkMode)}
            className='text-gray-700 dark:text-gray-100'>
            {isDarkMode ?
                <MoonIcon className="h-5 w-5" />
                :
                <SunIcon className="h-6 w-6" />
            }
        </button>
    )
}
export default ToggleDarkModeButton
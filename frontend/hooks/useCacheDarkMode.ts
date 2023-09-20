import { useEffect } from 'react';

import { useLocalStorage } from 'react-use';
import { useToggleDarkMode } from './useToggleDarkMode';

const Theme = {
    Dark: 'dark',
    Light: 'light',
} as const

type UseDarkMode = () => {
    isDarkMode: boolean
    toggle: (isDark: boolean) => void
}

export const useCacheDarkMode: UseDarkMode = () => {
    const [value, setValue] = useLocalStorage<typeof Theme['Dark' | 'Light']>('theme')
    const { isDarkMode, toggle } = useToggleDarkMode()

    const persistToggle = (isDark: boolean) => {
        toggle(isDark)
        setValue(isDark ? Theme.Dark : Theme.Light)
    }

    useEffect(() => {
        if (
            value === Theme.Dark ||
            (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)
        ) {
            toggle(true)
            setValue(Theme.Dark)
        } else {
            toggle(false)
            setValue(Theme.Light)
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [value])

    return { isDarkMode, toggle: persistToggle }
}
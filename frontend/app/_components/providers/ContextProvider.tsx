"use client"
import { ReactNode } from 'react'

const ContextProvider = ({ children }: { children: ReactNode }) => {
    return (
        <>{children}</>
    )
}
export default ContextProvider
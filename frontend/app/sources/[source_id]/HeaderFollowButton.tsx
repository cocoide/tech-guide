"use client"

import { useEffect, useState } from 'react'

interface Props {
    doFollowFunc: () => void
    unFollowFunc: () => void
    isFollowing: boolean | undefined
}
const HeaderFollowButton = ({ isFollowing, doFollowFunc, unFollowFunc }: Props) => {
    const [onClicked, setOnClicked] = useState(false)
    useEffect(() => {
        if (isFollowing) {
            setOnClicked(true)
        }
    }, [isFollowing])
    return (
        <>
            {onClicked ?
                <button onClick={()=>{unFollowFunc;setOnClicked(false)}} className='text-sm custom-badge p-2 bg-gray-400 text-white rounded-xl'>
                    フォロー中
                </button>
                :
                <button onClick={()=>{doFollowFunc;setOnClicked(true)}} className="text-sm custom-badge p-2 ring-1 bg-cyan-50/50 ring-cyan-300 text-cyan-300 rounded-xl">
                    フォローする
                </button>
            }
        </>
    )
}

export default HeaderFollowButton
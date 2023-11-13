"use client"

import { useEffect, useState } from 'react';

interface Props {
    doFollowFunc: () => Promise<void>;
    unFollowFunc: () => Promise<void>;
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
                <button
                    onClick={async () => {
                        await unFollowFunc();
                        setOnClicked(false);
                    }}
                    className='text-sm custom-badge p-[6px] bg-gray-400 text-white rounded-xl'>
                    フォロー中
                </button>
                :
                <button
                    onClick={async () => {
                        await doFollowFunc();
                        setOnClicked(true);
                    }}
                    className="text-sm custom-badge p-[6px] ring-1 bg-cyan-50/50 ring-cyan-300 text-cyan-300 rounded-xl">
                    フォローする
                </button>
            }
        </>
    )
}

export default HeaderFollowButton
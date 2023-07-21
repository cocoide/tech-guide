"use client"

import { API_URL } from '@/libs/constant'
import { Article } from '@/types/model'
import { XMarkIcon } from '@heroicons/react/24/outline'
import { useQuery } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'

interface Props extends ArticleParams {
}

export default function Page({ params }: Props) {
    const router = useRouter()
    function handleClose() {
        router.back()
    }
    const getArticleDetail = async (articleId: string) => {
        const url = API_URL + `/article/${articleId}`
        console.log()
        const res = await fetch(url);
        return await res.json();
    };
    const { data } = useQuery<Article>({
        queryKey: ['articleDetail'],
        queryFn: () => getArticleDetail(params.article_id)
    })
    console.log(data)
    return (
        <>
            <button onClick={handleClose}
                className="z-30 bg-black/40  fixed inset-0 backdrop-blur-[3px] animate-appear"></button>
            <button className="lg:hidden fixed left-0 top-3 py-1 pr-1 pl-5 rounded-r-full bg-black/50 z-50" onClick={handleClose}
            ><XMarkIcon className='h-6 w-6 text-white' /></button>

            <button className="hidden lg:flex fixed left-5 top-5 p-2 rounded-full bg-white z-50" onClick={handleClose}
            ><XMarkIcon className='h-5 w-5 text-gray-800' /></button>

            <div className="z-40 fixed bg-white inset-0 sm:rounded-sm
            sm:top-[100px] lg:left-[100px] lg:right-[100px]">
            </div>
        </>
    )
}
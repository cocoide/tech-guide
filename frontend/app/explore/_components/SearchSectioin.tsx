"use client"
import { MagnifyingGlassIcon } from '@heroicons/react/24/outline';
import { useRouter } from 'next/navigation';
import { ChangeEvent, FormEvent, useRef, useState } from 'react';

export default function SearchSection() {
    const router = useRouter()
    const inputRef = useRef<HTMLInputElement>(null);
    function handleChange(e: ChangeEvent<HTMLInputElement>) {
        setQuery(e.target?.value)
    }
    function handleSubmit(e: FormEvent<HTMLFormElement>) {
        e.preventDefault();
        if (query) {
            router.push(`/search?q=${encodeURIComponent(query)}`)
        }
    }
    const [query, setQuery] = useState("")
    return (
        <div className="flex flex-row items-center dark:text-white">
            <form onSubmit={handleSubmit}
                className="flex items-center w-[100%] pl-3 pr-5 rounded-full custom-border dark:bg-gray-900">
                <MagnifyingGlassIcon className="h-5 w-5  text-gray-500" />
                <input ref={inputRef} onChange={handleChange}
                    type="search" placeholder="タイトル、#トピック"
                    className="w-[100%] p-1 border-transparent focus:ring-0 border-none h-10 dark:bg-gray-900" />
            </form>
        </div>
    )
}
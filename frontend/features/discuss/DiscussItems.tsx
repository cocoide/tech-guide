"use client";


import { Comment } from '@/types/model';
import * as React from "react";
import InfiniteScroll from "react-infinite-scroller";
import DiscussCard from './DiscussCard';
import DiscussCardLoader from './DiscussCardLoader';

interface Props {
    Discusss: Comment[];
    fetchFunc: (page?: number, filterID?: number) => Promise<Comment[]>;
    filterID?: number;
};

const paginateLimit = 12

export default function DiscussItems({ Discusss, fetchFunc, filterID }: Props) {
    const fetching = React.useRef(false);
    const [pages, setPages] = React.useState([Discusss]);
    const [hasMoreItems, setHasMoreItems] = React.useState(true);
    const items = pages.flatMap((page) => page);

    const loadMore = async (page: number) => {
        if (!fetching.current) {
            try {
                fetching.current = true;

                const data = await fetchFunc(page, filterID);

                if (data.length < paginateLimit) {
                    setHasMoreItems(false);
                } else {
                    setPages((prev) => [...prev, data]);
                }
            } finally {
                fetching.current = false;
            }
        }
    };

    return (
        <InfiniteScroll
            hasMore={hasMoreItems}
            pageStart={1}
            loadMore={loadMore}
            loader={
                <>
                    {Array(1).fill(null).map((_, index) => (
                        <DiscussCardLoader key={index + "loader"} />
                    ))}
                </>
            }
            element="main"
            className='w-full grid lg:grid-cols-2 gap-7 p-5 divide-y-[0.5px] lg:divide-x-[0.5px]'
        >
            {items && items.map((item, index) => (
                <DiscussCard key={index} comment={item} />
            ))}
        </InfiniteScroll>
    );
}

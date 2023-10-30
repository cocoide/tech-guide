"use client";

import ArticleCard from '@/app/(home)/_components/ArticleCard';
import LoaderArticleCard from '@/app/(home)/_components/LoaderArticleCard';
import { Article } from '@/types/model';
import * as React from "react";
import InfiniteScroll from "react-infinite-scroller";

type ItemsProps = {
    initialItems: Article[];
    fetchItems: (page?: number) => Promise<Article[]>;
};

const paginateLimit = 6

export default function FeedItems({ initialItems, fetchItems }: ItemsProps) {
    const fetching = React.useRef(false);
    const [pages, setPages] = React.useState([initialItems]);
    const [hasMoreItems, setHasMoreItems] = React.useState(true);
    const items = pages.flatMap((page) => page);

    const loadMore = async (page: number) => {
        if (!fetching.current) {
            try {
                fetching.current = true;

                const data = await fetchItems(page);

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
                    {Array(10).fill(null).map((_, index) => (
                        <LoaderArticleCard key={index + "loader"} />
                    ))}
                </>
            }
            element="main"
            className='w-full grid sm:grid-cols-2 xl:grid-cols-3 gap-2 p-[20px]'
        >
            {items && items.map((item) => (
                <ArticleCard key={item.title + item.id + item.created_at} article={item} />
            ))}
        </InfiniteScroll>
    );
}

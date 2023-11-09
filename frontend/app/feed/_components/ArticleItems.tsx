"use client";

import ArticleCard from '@/app/(home)/_components/ArticleCard';
import LoaderArticleCard from '@/app/(home)/_components/LoaderArticleCard';
import { Article } from '@/types/model';
import * as React from "react";
import InfiniteScroll from "react-infinite-scroller";

interface Props {
    articles: Article[];
    fetchFunc: (page?: number) => Promise<Article[]>;
};

const paginateLimit = 12

export default function ArticleItems({ articles, fetchFunc }: Props) {
    const fetching = React.useRef(false);
    const [pages, setPages] = React.useState([articles]);
    const [hasMoreItems, setHasMoreItems] = React.useState(true);
    const items = pages.flatMap((page) => page);

    const loadMore = async (page: number) => {
        if (!fetching.current) {
            try {
                fetching.current = true;

                const data = await fetchFunc(page);

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

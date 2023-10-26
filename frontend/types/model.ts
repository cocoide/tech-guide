export type Collection ={
    id: number
    name: string
    descriptoin : string
    visibility: number
    account_id: string
    created_at: string
    articles: Article[]
}
export type Article={
    id: number
    title: string
    original_url: string
    thumbnail_url: string
    summary: string
    created_at: string
    topics: Topic[]
    source: Source
    rating: ArticleRating
}

export type ArticleRating={
    owned_stocks: number
    origin_stocks: number
    pocket_stocks: number
    hatena_stocks: number
    updated_at: string
}
export type Topic={
    id: number
    name: string
    icon_url: string
}
export type Category = {
    id: number
    name: string
    icon_url: string
    topics: Topic[]
}

export type Account={
    id: number
    display_name: string
    avatar_url: string
}

export type Source={
    id : number
    name : string
    icon_url: string
    domain: string
}

export type Contribution={
    date: Date
    points: number
}

export type Comment={
    id: number
    account_id: number
    article_id: number
    account: Account
    content: string
}

export type AccountSession = {
    account_id: number
    display_name: string
    avatar_url: string
    features?: number[]
}
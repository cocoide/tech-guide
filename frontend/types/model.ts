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
}

export type Topic={
    id: number
    name: string
    icon_url: string
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
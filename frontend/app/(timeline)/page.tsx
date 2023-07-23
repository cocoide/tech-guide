import { Article } from '@/types/model';
import { authAPI } from '../_functions/auth';
import ArticleCard from './_components/ArticleCard';
import { articleAPI } from './_functions/article';

export default async function ArticlePage() {
    const { token } = await authAPI.GetAuthSession()
    const articles = await articleAPI.getLatestArticles()
    let recommends: Article[] | undefined = []
    if (token) {
        const response = await articleAPI.GetRecommendArticles(token);
        recommends = response.data;
    }
    const latest_articles = articles?.filter(article => !recommends?.some(rec => rec.id === article.id));
    return (
        <div className="flex flex-col w-full">
            {/* {recommends && recommends.length > 0 &&
                <div className="bg-cyan-50/70  lg:px-[10%] border-y-[1px] border-cyan-300/50">
                    <div className="p-3 text-cyan-300 flex flex-row items-center">
                        <ChatBubbleOvalLeftEllipsisIcon className="h-5 w-5" />
                        <div>おすすめ</div>
                    </div>
                    <div className="w-full gap-3 grid lg:grid-cols-2">
                        {recommends.map(recommend => (
                            <ArticleCard key={recommend.title} article={recommend} />
                        ))}
                    </div>
                </div>
            } */}
            <div className="min-h-screen w-full divide-y-[0.5px]">
                {latest_articles?.map(article => (
                    <ArticleCard key={article.title} article={article} />
            )
            )}
        </div>
        </div>
    )
}
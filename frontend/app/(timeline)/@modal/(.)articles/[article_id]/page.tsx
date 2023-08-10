import { articleAPI } from '@/app/(timeline)/_functions/article';
import ModalContent from './components/ModalContent';
import Overlay from './components/Overlay';

interface Props extends ArticleParams {
}

export default async function Page({ params }: Props) {
    const { data: article } = await articleAPI.GetArticleDetail(params.article_id)
    return (
        <>
            <Overlay />
            <ModalContent article={article} />
        </>
    )
}
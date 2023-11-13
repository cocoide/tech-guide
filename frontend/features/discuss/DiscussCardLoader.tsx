
import LoaderArticleCard from '@/app/(home)/_components/LoaderArticleCard';
import HStack from '@/app/_components/elements/ui/HStack';
import WStack from '@/app/_components/elements/ui/WStack';

const DiscussCardLoader = () => {
    return (
        <WStack className="custom-text">
            <div className="custom-loader h-10 w-10 rounded-full"></div>
            <HStack className='space-y-3'>
                <div className="h-5 w-15 custom-loader" />
                <div className="h-5 w-20 custom-loader" />
                <LoaderArticleCard />
            </HStack>
        </WStack>
    )
}
export default DiscussCardLoader
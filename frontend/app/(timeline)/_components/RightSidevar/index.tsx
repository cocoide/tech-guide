import SourceSection from './SourceSection';
import TopicSection from './TopicSection';

export default async function RightSidevar() {
    return (
        <div className="md:border-l-[0.5px] min-h-screen w-[270px]
        p-3 flex flex-col space-y-3">
            <TopicSection />
            <SourceSection />
        </div>
    )
}
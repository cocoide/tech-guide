import Image from 'next/image';
import HStack from '../_components/elements/ui/HStack';

export default async function AboutPage() {
    return (
        <HStack className="justify-center items-center w-full">
            <HStack className="justify-center items-center w-full space-y-3 p-10 bg-white">
                <h1
                    className="font-extrabold text-transparent text-5xl bg-clip-text bg-gradient-to-r from-cyan-300 to-indigo-300"
                >TechGuide</h1>
                <div className="text-2xl text-gray-600 font-bold">技術のキャチアップをこれ一つで</div>
                <Image src={"/about/view.png"} alt='view' width={400} height={200} className="rounded-md h-[250px] w-[450px]  custom-border"></Image>
                <button className="text-white bg-cyan-300 p-[7px] rounded-xl text-md shadow-sm">はじめる</button>
            </HStack>

            <HStack className="justify-center items-center w-full space-y-3 bg-gray-50 border-t border-gray-300 p-10">
                <div className="text-gray-600 font-bold text-2xl">Feature</div>
                <div className="w-full grid md:grid-cols-2 gap-10">
                    <FeatureSection thumbnail_url={'/about/interface.png'}
                        text={{
                            title: '記事コンテンツだけでなく,動画、スライドの情報も直感的に',
                            description: 'Youtube, Github, SpeakerDeckなどの情報もトピックで包括的にキャッチアップ可能'
                        }} />
                    <FeatureSection thumbnail_url={'/about/custom.svg'}
                        text={{
                            title: 'カスタマイズ可能な高度なフィード',
                            description: 'トピック、ドメイン、キーワード除外などを設定してあなた好みのフィードへ',
                        }} />
                    <FeatureSection thumbnail_url={'/about/point.svg'}
                        text={{
                            title: '客観的な評価ポイントでトレンドを掴む',
                            description: '期間ごとの絞り込み、急上昇中の記事にも対応',
                        }} />
                    <FeatureSection thumbnail_url={'/about/summary.svg'}
                        text={{
                            title: '記事の要約・概略機能',
                            description: 'キャチアップにかける労力を省略',
                        }} />
                </div>
            </HStack>
        </HStack>
    )
}

interface FeatureSectionProps {
    thumbnail_url: string;
    text: {
        title: string
        description?: string
    }
}
const FeatureSection = ({ text, thumbnail_url }: FeatureSectionProps) => {
    return (
        <div className="rounded-xl custom-border p-5 flex flex-col space-y-5 bg-white">
            <div className="w-full h-min-[200px] rounded-xl custom-border flex flex-row items-end justify-center overflow-hidden bg-gray-100">
                <Image src={thumbnail_url} alt='thumbnail' width={200} height={200} className="rounded-md h-[200px] w-[200px]"></Image>
            </div>
            <div className="font-bold text-gray-700">{text.title}</div>
            <div className="text-gray-500">{text.description}</div>
        </div>
    )
}
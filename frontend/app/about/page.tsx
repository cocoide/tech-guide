import Image from 'next/image';
import HStack from '../_components/elements/ui/HStack';

export default async function AboutPage() {
    return (
        <HStack className="justify-center items-center w-full space-y-3">
            <HStack className="justify-center items-center w-full space-y-3 border-t p-10 bg-black">
                <h1
                    className="font-extrabold text-transparent text-8xl bg-clip-text bg-gradient-to-r from-cyan-400 to-indigo-600"
                >TechGuide</h1>
                <div className="text-2xl text-white font-bold">技術のキャチアップをこれ一つで</div>
                <Image src={"/about/view.png"} alt='view' width={400} height={200} className="rounded-md h-[200px] w-[450px] lg:h-[300px] lg:w-[600pc] custom-border"></Image>
                <button className="text-white bg-cyan-300 p-[7px] rounded-xl text-md shadow-sm">はじめる</button>
            </HStack>

            <HStack className="justify-center items-center w-full space-y-3 bg-gray-50 border-t border-gray-300 p-10">
                <div className="text-gray-600 font-bold text-2xl">Feature</div>
                <div className="w-full grid md:grid-cols-3 gap-10">
                    <div className="rounded-xl aspect-square h-30 custom-border">A</div>
                    <div className="rounded-xl aspect-square h-30 custom-border">B</div>
                    <div className="rounded-xl aspect-square h-30 custom-border">C</div>
                </div>
            </HStack>
        </HStack>
    )
}
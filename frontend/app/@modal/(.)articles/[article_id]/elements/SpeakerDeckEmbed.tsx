import { api } from '@/app/_functions/API'

export default async function SpeakerDeckEmbed({ url }: { url: string }) {
    const { data: id } = await api.get<string>(`/speakerdeck?url=${url}`, "force-cache")
    return (
        <>
                <iframe
                className='w-[350px] md:w-[400px]  h-[250px] md:h-[300px] lg:h-[400px] lg:w-[500px] rounded-xl bg-gray-100'
                    width="400"
                    height="300"
                    src={`//speakerdeck.com/player/${id}`}
                    title="YouTube video player"
                    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                    allowFullScreen>
                </iframe>
        </>
    )
}
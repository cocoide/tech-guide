import { clsx } from '@/utils/clsx';

function YouTubeEmbed({ youtube_id, style }: { youtube_id: string, style?: string }) {
        return (
            <iframe
                className={clsx(style!, 'w-full h-[250px] md:h-[300px] lg:h-[400px] lg:w-[500px] rounded-xl animate-appear')}
                width="400"
                height="300"
                src={`https://www.youtube.com/embed/${youtube_id}`}
                title="YouTube video player"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                allowFullScreen>
            </iframe>
        );
}

export default YouTubeEmbed;
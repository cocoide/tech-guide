import { clsx } from '@/utils/clsx';

function YouTubeEmbed({ youtube_id, width_style }: { youtube_id: string, width_style?: string }) {
    let style: string = "lg:w-[500px]"
    if (width_style) {
        style = width_style
    }
        return (
            <iframe
                className={clsx(width_style!, 'w-full h-[250px] md:h-[300px] lg:h-[400px]  rounded-xl animate-appear')}
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
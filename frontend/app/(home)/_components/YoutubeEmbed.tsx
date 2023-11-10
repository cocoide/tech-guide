import { clsx } from '@/utils/clsx';

function YouTubeEmbed({ youtube_id }: { youtube_id: string }) {
        return (
            <iframe
                className={clsx('w-full h-[250px] md:h-[300px] lg:w-[500px] rounded-xl animate-appear')}
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
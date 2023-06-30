
function YouTubeEmbed({ youtube_id }: { youtube_id: string }) {
        return (
            <iframe
                className='w-full md:h-[300px] md:w-[500px] rounded-xl'
                width="500"
                height="300"
                src={`https://www.youtube.com/embed/${youtube_id}`}
                title="YouTube video player"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                allowFullScreen>
            </iframe>
        );
}

export default YouTubeEmbed;

function YouTubeEmbed({ youtube_id }: { youtube_id: string }) {
        return (
            <iframe
                className='w-full md:h-[300px] lg:h-[400px] lg:w-[600px] rounded-xl animate-appear'
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
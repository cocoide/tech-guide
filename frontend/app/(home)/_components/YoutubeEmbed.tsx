
function YouTubeEmbed({ url }: { url: string }) {
    const getYouTubeId = (url: string) => {
        const match = url.match(/^(?:https?:\/\/)?(?:www\.)?(?:m\.)?(?:youtu\.be\/|youtube\.com\/(?:watch\?(?:\S*?&)?v=|(?:embed|v|vi|user)\/))([\w-]+)/);
        return match ? match[1] : null;
    };

    const id = getYouTubeId(url);

    if (id) {
        return (
            <iframe
                width="200"
                height="120"
                src={`https://www.youtube.com/embed/${id}`}
                title="YouTube video player"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                allowFullScreen>
            </iframe>
        );
    }

    return null;
}

export default YouTubeEmbed;
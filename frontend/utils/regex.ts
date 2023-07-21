export function extractYoutubeID(url?: string): string | null {
    if (url === undefined){
        return null;
    }
    const match = url.match(/^(?:https?:\/\/)?(?:www\.)?(?:m\.)?(?:youtu\.be\/|youtube\.com\/(?:watch\?(?:\S*?&)?v=|(?:embed|v|vi|user)\/))([\w-]+)/);
    return match ? match[1] : null;
}
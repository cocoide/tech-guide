export function extractYoutubeID(url?: string): string | undefined {
    if (url === undefined){
        return undefined;
    }
    const match = url.match(/^(?:https?:\/\/)?(?:www\.)?(?:m\.)?(?:youtu\.be\/|youtube\.com\/(?:watch\?(?:\S*?&)?v=|(?:embed|v|vi|user)\/))([\w-]+)/);
    return match ? match[1] : undefined;
}
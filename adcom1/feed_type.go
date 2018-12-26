package adcom

// FeedType represents types of feeds, typically for audio.
type FeedType int

// Types of feeds, typically for audio.
const (
	FeedMusicService   FeedType = 1 // Music Service
	FeedRadioBroadcast FeedType = 2 // FM/AM Broadcast
	FeedPodcast        FeedType = 3 // Podcast
)

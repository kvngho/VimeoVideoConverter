package persistence

type ProductReview struct {
	PlayableVideo string
	ReviewVideo   string
}

type DeepingTalk struct {
	TalkVideo     string
	PlayableVideo string
}

type ProductVideo struct {
	URL           string
	PlayableVideo string
}

type UserVideo struct {
	VideoID       uint32
	VideoURL      string
	PlayableVideo string
}

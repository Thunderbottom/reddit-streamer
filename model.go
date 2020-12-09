package main

type options struct {
	ConfigFile string `short:"c" long:"config" description:"Path to the configuration file" default:"config.toml"`
}

type config struct {
	General  general  `koanf:"general"`
	Telegram telegram `koanf:"telegram"`
	Filters  filters  `koanf:"filters"`
}

type general struct {
	Debug    bool `koanf:"debug"`
	Interval int  `koanf:"interval"`
}

type filters struct {
	Subreddit      string   `koanf:"subreddit"`
	Keywords       []string `koanf:"keywords"`
	MediaWhitelist []string `koanf:"media_whitelist"`
}

type telegram struct {
	APIKey         string  `koanf:"api_key"`
	ChannelIDs     []int64 `koanf:"channel_ids"`
	PostRedditLink bool    `koanf:"post_reddit_link"`
}

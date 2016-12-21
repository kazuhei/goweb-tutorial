package settings

// ApplicationSettings db情報など
type ApplicationSettings struct {
	DB      string
	Cache   string
	IsDebug bool
}

// NewApplicationSettings 設定項目を含めたstructの作成
func NewApplicationSettings() *ApplicationSettings {
	// flagかファイルなどから取ってくると良いかもね
	return &ApplicationSettings{
		DB:      "mysql",
		Cache:   "memcache",
		IsDebug: true,
	}
}

package constant

import "time"

const (
	AppName          = "moss"
	AppVersion       = "0.1.4"
	DefaultAdminPath = "/admin"
	ThemesDir        = "./themes"
	PublicDir        = "./public"
	LogDir           = "./runtime/log"
	CacheDir         = "./runtime/cache"
	UploadDir        = "./public/upload"
	UploadDomain     = "/upload/"
	LogoFilePath     = "/logo.png"
)

var (
	AppStartTime = time.Now()
)

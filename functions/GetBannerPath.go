package ascii

func GetBannerPath(banner string) (string, bool) {
	var bannerPath string
	if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
		return bannerPath, false
	}
	bannerPath = "./banners/" + banner + ".txt"
	return bannerPath, true
}

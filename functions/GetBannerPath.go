package ascii

func GetBannerPath(banner string) string {
	var bannerPath string
	if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
		return "Error: banner type not correct!\nUse: shadow, standard, or thinkertoy."
	}
	bannerPath = "./banners/" + banner + ".txt"
	return bannerPath
}

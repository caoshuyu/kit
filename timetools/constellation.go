package timetools

const (
	CONSTELLATION_ARIES       = 1  //白羊
	CONSTELLATION_TAURUS      = 2  //金牛
	CONSTELLATION_GEMINI      = 3  //双子
	CONSTELLATION_CANCER      = 4  //巨蟹
	CONSTELLATION_LEO         = 5  //狮子
	CONSTELLATION_VIRGO       = 6  //处女
	CONSTELLATION_LIBRA       = 7  //天秤
	CONSTELLATION_SCORPIO     = 8  //天蝎
	CONSTELLATION_SAGITTARIUS = 9  //射手
	CONSTELLATION_CAPRICORN   = 10 //摩羯
	CONSTELLATION_AQUARIUS    = 11 //水瓶
	CONSTELLATION_PISCES      = 12 //双鱼
)

func GetConstellation(month, day int64) (constellation int64) {
	if 0 == month || 0 == day {
		return
	}
	switch month {
	case 1:
		if day >= 20 {
			constellation = CONSTELLATION_AQUARIUS
		} else {
			constellation = CONSTELLATION_CAPRICORN
		}
	case 2:
		if day >= 19 {
			constellation = CONSTELLATION_PISCES
		} else {
			constellation = CONSTELLATION_AQUARIUS
		}
	case 3:
		if day >= 21 {
			constellation = CONSTELLATION_ARIES
		} else {
			constellation = CONSTELLATION_PISCES
		}
	case 4:
		if day >= 20 {
			constellation = CONSTELLATION_TAURUS
		} else {
			constellation = CONSTELLATION_ARIES
		}
	case 5:
		if day >= 21 {
			constellation = CONSTELLATION_GEMINI
		} else {
			constellation = CONSTELLATION_TAURUS
		}
	case 6:
		if day >= 22 {
			constellation = CONSTELLATION_CANCER
		} else {
			constellation = CONSTELLATION_GEMINI
		}
	case 7:
		if day >= 23 {
			constellation = CONSTELLATION_LEO
		} else {
			constellation = CONSTELLATION_CANCER
		}
	case 8:
		if day >= 23 {
			constellation = CONSTELLATION_VIRGO
		} else {
			constellation = CONSTELLATION_LEO
		}
	case 9:
		if day >= 23 {
			constellation = CONSTELLATION_LIBRA
		} else {
			constellation = CONSTELLATION_VIRGO
		}
	case 10:
		if day >= 24 {
			constellation = CONSTELLATION_SCORPIO
		} else {
			constellation = CONSTELLATION_LIBRA
		}
	case 11:
		if day >= 23 {
			constellation = CONSTELLATION_SAGITTARIUS
		} else {
			constellation = CONSTELLATION_SCORPIO
		}
	case 12:
		if day >= 22 {
			constellation = CONSTELLATION_CAPRICORN
		} else {
			constellation = CONSTELLATION_SAGITTARIUS
		}
	}
	return
}

package wx

var (
	sunClear = []string{
		"\033[38;5;226m    \\   /    \033[0m",
		"\033[38;5;226m     .-.     \033[0m",
		"\033[38;5;226m  ― (   ) ―  \033[0m",
		"\033[38;5;226m     `-'     \033[0m",
		"\033[38;5;226m    /   \\    \033[0m",
	}
	moonClear = []string{
		"\033[38;5;117m    _..._    \033[0m",
		"\033[38;5;117m  .::'   `.  \033[0m",
		"\033[38;5;117m ::  :     : \033[0m",
		"\033[38;5;117m `::.     .' \033[0m",
		"\033[38;5;117m   `'-..-'   \033[0m",
	}
	sunCloud = []string{
		"\033[38;5;226m   \\  /\033[0m      ",
		"\033[38;5;226m _ /\"\"\033[38;5;250m.-.    \033[0m",
		"\033[38;5;226m   \\_\033[38;5;250m(   ).  \033[0m",
		"\033[38;5;226m   /\033[38;5;250m(___(__) \033[0m",
		"             ",
	}
	overcast = []string{
		"             ",
		"\033[38;5;240;1m     .--.    \033[0m",
		"\033[38;5;240;1m  .-(    ).  \033[0m",
		"\033[38;5;240;1m (___.__)__) \033[0m",
		"             ",
	}
	fog = []string{
		"             ",
		"\033[38;5;251m _ - _ - _ - \033[0m",
		"\033[38;5;251m  _ - _ - _  \033[0m",
		"\033[38;5;251m _ - _ - _ - \033[0m",
		"             ",
	}
	sunCloudDrizzle = []string{
		"\033[38;5;226m _`/\"\"\033[38;5;250m.-.    \033[0m",
		"\033[38;5;226m  ,\\_\033[38;5;250m(   ).  \033[0m",
		"\033[38;5;226m   /\033[38;5;250m(___(__) \033[0m",
		"\033[38;5;111m     ' ' ' ' \033[0m",
		"\033[38;5;111m    ' ' ' '  \033[0m",
	}
	sunCloudHeavyDrizzle = []string{
		"\033[38;5;226m _`/\"\"\033[38;5;240;1m.-.    \033[0m",
		"\033[38;5;226m  ,\\_\033[38;5;240;1m(   ).  \033[0m",
		"\033[38;5;226m   /\033[38;5;240;1m(___(__) \033[0m",
		"\033[38;5;21;1m   ‚'‚'‚'‚'  \033[0m",
		"\033[38;5;21;1m   ‚'‚'‚'‚'  \033[0m",
	}
	sunCloudFreezingDrizzle = []string{
		"\033[38;5;226m _`/\"\"\033[38;5;250m.-.    \033[0m",
		"\033[38;5;226m  ,\\_\033[38;5;250m(   ).  \033[0m",
		"\033[38;5;226m   /\033[38;5;250m(___(__) \033[0m",
		"\033[38;5;111m     ' \033[38;5;255m*\033[38;5;111m ' \033[38;5;255m* \033[0m",
		"\033[38;5;255m    *\033[38;5;111m ' \033[38;5;255m*\033[38;5;111m '  \033[0m",
	}
	cloudRain = []string{
		"\033[38;5;250m     .-.     \033[0m",
		"\033[38;5;250m    (   ).   \033[0m",
		"\033[38;5;250m   (___(__)  \033[0m",
		"\033[38;5;111m    ' ' ' '  \033[0m",
		"\033[38;5;111m   ' ' ' '   \033[0m",
	}
	cloudHeavyRain = []string{
		"\033[38;5;240;1m     .-.     \033[0m",
		"\033[38;5;240;1m    (   ).   \033[0m",
		"\033[38;5;240;1m   (___(__)  \033[0m",
		"\033[38;5;21;1m  ‚'‚'‚'‚'   \033[0m",
		"\033[38;5;21;1m  ‚'‚'‚'‚'   \033[0m",
	}
	cloudFreezingRain = []string{
		"\033[38;5;250m     .-.     \033[0m",
		"\033[38;5;250m    (   ).   \033[0m",
		"\033[38;5;250m   (___(__)  \033[0m",
		"\033[38;5;111m    ' \033[38;5;255m*\033[38;5;111m ' \033[38;5;255m*  \033[0m",
		"\033[38;5;255m   *\033[38;5;111m ' \033[38;5;255m*\033[38;5;111m '   \033[0m",
	}
	cloudSnow = []string{
		"\033[38;5;250m     .-.     \033[0m",
		"\033[38;5;250m    (   ).   \033[0m",
		"\033[38;5;250m   (___(__)  \033[0m",
		"\033[38;5;255m    *  *  *  \033[0m",
		"\033[38;5;255m   *  *  *   \033[0m",
	}
	cloudHeavySnow = []string{
		"\033[38;5;240;1m     .-.     \033[0m",
		"\033[38;5;240;1m    (   ).   \033[0m",
		"\033[38;5;240;1m   (___(__)  \033[0m",
		"\033[38;5;255;1m   * * * *   \033[0m",
		"\033[38;5;255;1m  * * * *    \033[0m",
	}
	sunCloudSnowShowers = []string{
		"\033[38;5;226m _`/\"\"\033[38;5;250m.-.    \033[0m",
		"\033[38;5;226m  ,\\_\033[38;5;250m(   ).  \033[0m",
		"\033[38;5;226m   /\033[38;5;250m(___(__) \033[0m",
		"\033[38;5;255m     *  *  * \033[0m",
		"\033[38;5;255m    *  *  *  \033[0m",
	}
	sunCloudHeavySnowShowers = []string{
		"\033[38;5;226m _`/\"\"\033[38;5;240;1m.-.    \033[0m",
		"\033[38;5;226m  ,\\_\033[38;5;240;1m(   ).  \033[0m",
		"\033[38;5;226m   /\033[38;5;240;1m(___(__) \033[0m",
		"\033[38;5;255;1m    * * * *  \033[0m",
		"\033[38;5;255;1m   * * * *   \033[0m",
	}
	thunderstorm = []string{
		"\033[38;5;226m _`/\"\"\033[38;5;250m.-.    \033[0m",
		"\033[38;5;226m  ,\\_\033[38;5;250m(   ).  \033[0m",
		"\033[38;5;226m   /\033[38;5;250m(___(__) \033[0m",
		"\033[38;5;228;5m    ⚡\033[38;5;111;25m' '\033[38;5;228;5m⚡\033[38;5;111;25m' ' \033[0m",
		"\033[38;5;111m    ' ' ' '  \033[0m",
	}
	thunderstormHail = []string{
		"\033[38;5;226m _`/\"\"\033[38;5;250m.-.    \033[0m",
		"\033[38;5;226m  ,\\_\033[38;5;250m(   ).  \033[0m",
		"\033[38;5;226m   /\033[38;5;250m(___(__) \033[0m",
		"\033[38;5;255m     *\033[38;5;228;5m⚡\033[38;5;255;25m*\033[38;5;228;5m⚡\033[38;5;255;25m* \033[0m",
		"\033[38;5;255m    *  *  *  \033[0m",
	}
	thunderstormHeavyHail = []string{
		"\033[38;5;240;1m     .-.     \033[0m",
		"\033[38;5;240;1m    (   ).   \033[0m",
		"\033[38;5;240;1m   (___(__)  \033[0m",
		"\033[38;5;21;1m  ‚'\033[38;5;228;5m⚡\033[38;5;21;25m'‚\033[38;5;228;5m⚡\033[38;5;21;25m‚' \033[0m",
		"\033[38;5;21;1m  ‚'‚'\033[38;5;228;5m⚡\033[38;5;21;25m'‚'  \033[0m",
	}
)

var unknownIcon = WeatherIcon{
	Description: "Unknown conditions",
	ASCII: []string{
		"             ",
		"    .--.     ",
		"  -(    )-   ",
		" (__.--.__)  ",
		"             ",
	},
}

var iconMap = map[WeatherCode]WeatherIcon{
	0:   {Description: "Clear sky", ASCII: sunClear},
	100: {Description: "Clear night", ASCII: moonClear},
	1:   {Description: "Mainly clear", ASCII: sunCloud},
	2:   {Description: "Partly cloudy", ASCII: sunCloud},
	3:   {Description: "Overcast", ASCII: overcast},
	45:  {Description: "Foggy", ASCII: fog},
	48:  {Description: "Depositing rime fog", ASCII: fog},
	51:  {Description: "Light drizzle", ASCII: sunCloudDrizzle},
	53:  {Description: "Moderate drizzle", ASCII: sunCloudDrizzle},
	55:  {Description: "Dense drizzle", ASCII: sunCloudHeavyDrizzle},
	56:  {Description: "Light freezing drizzle", ASCII: sunCloudFreezingDrizzle},
	57:  {Description: "Dense freezing drizzle", ASCII: sunCloudFreezingDrizzle},
	61:  {Description: "Slight rain", ASCII: cloudRain},
	63:  {Description: "Moderate rain", ASCII: cloudRain},
	65:  {Description: "Heavy rain", ASCII: cloudHeavyRain},
	66:  {Description: "Light freezing rain", ASCII: cloudFreezingRain},
	67:  {Description: "Heavy freezing rain", ASCII: cloudFreezingRain},
	71:  {Description: "Slight snow fall", ASCII: cloudSnow},
	73:  {Description: "Moderate snow fall", ASCII: cloudSnow},
	75:  {Description: "Heavy snow fall", ASCII: cloudHeavySnow},
	77:  {Description: "Snow grains", ASCII: cloudSnow},
	80:  {Description: "Slight rain showers", ASCII: sunCloudDrizzle},
	81:  {Description: "Moderate rain showers", ASCII: sunCloudDrizzle},
	82:  {Description: "Violent rain showers", ASCII: sunCloudHeavyDrizzle},
	85:  {Description: "Slight snow showers", ASCII: sunCloudSnowShowers},
	86:  {Description: "Heavy snow showers", ASCII: sunCloudHeavySnowShowers},
	95:  {Description: "Thunderstorm", ASCII: thunderstorm},
	96:  {Description: "Thunderstorm with slight hail", ASCII: thunderstormHail},
	99:  {Description: "Thunderstorm with heavy hail", ASCII: thunderstormHeavyHail},
}

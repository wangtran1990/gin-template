package constant

const (
	// Cache time to live -  by second
	CACHE_TTL_NO_LIMIT = -1 // strongly NOT recommended
	CACHE_TTL_MINUTE_1 = 60
	CACHE_TTL_MINUTE_5 = 5 * 60
	CACHE_TTL_HOUR_1   = 60 * 60
	CACHE_TTL_DAY_1    = 24 * 60 * 60

	// Datetime format
	DATETIME_FORMAT_DEFAULT = "20060102150405"
)

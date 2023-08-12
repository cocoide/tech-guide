package key

const (
	PopularArticleIDs      = "popularArticleIDs"
	PersonalizedArticleIDs = "%d.personalizedArticleIDs" // %d = account_id
	RecentReads            = "%d.recentReads"            // %d = account_id
	Contributions          = "contributions.%s"          // $s = today
	UserSession            = "%d.session"                // %d = account_id
)

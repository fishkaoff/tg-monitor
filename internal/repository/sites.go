package sites


type SiteRepository interface {
	Save(chatID int64, site string) error 
	Get(chatID int64) ([]string, error)
	Delete(chatID int64, site string) error
	GetSite(chatID int64, site string) (string, error)
}
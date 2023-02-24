package consts

const (

	// commands
	GETMETRICCOMMAND  = "status"
	ADDSITECOMMAND    = "addsite"
	DELETESITECOMMAND = "deletesite"
	UNKMOWNCOMMAND    = "Use /help to see available commands"

	// Errors
	SITENOTADDED     = "Error while adding site"
	SITENOTDELETED   = "Error while deleting site"
	CANNOTGETSITES   = "Error while getting sites"
	SITESNOTFOUND    = "You dont have sites, use /addsite for add site"
	NOTURL           = "Incorrect URL"
	SITEALREADYADDED = "This site already added"

	// success
	SITEADDED   = "Site successfully added"
	SITEDELETED = "Site successfully deleted"

	// messages
	SITEAWAILABLE   = "Available✅"
	SITEUNAWAILABLE = "Unavailable❌"
	SENDDATA        = "Send me absolute url (example: https://google.com/)"
)

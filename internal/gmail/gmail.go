package gmail

import (
	"google.golang.org/api/gmail/v1"
)


type Gmailer interface {

}


type Gmail struct {
	Srv *gmail.Service
}

func NewGmail(srv *gmail.Service) *Gmail {
	return &Gmail{srv}
}


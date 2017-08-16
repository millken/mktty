package dt

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Action interface {
	Response() (data gin.H, err error)
}

type Param struct {
	Get url.Values
	//Content   *gin.Context
	Db *sqlx.DB
}

var Actions = map[string]func(Param) (Action, error){}

func DtRegister(name string, actionFactory func(Param) (Action, error)) {
	if actionFactory == nil {
		panic(" actionFactory is nil")
	}
	if _, dup := Actions[name]; dup {
		panic(" Register called twice for " + name)
	}
	Actions[name] = actionFactory
}

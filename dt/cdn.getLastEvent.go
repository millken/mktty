package dt

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	DtRegister("cdn.getLastEvent", NewCdnGetLastEvent)
}

type CdnGetLastEvent struct {
	param Param
}

type CdnEvent struct {
	Servername string
	Act        string
	Utime      string
	Setting    interface{}
}

func NewCdnGetLastEvent(param Param) (Action, error) {
	return &CdnGetLastEvent{
		param: param,
	}, nil
}

func (d *CdnGetLastEvent) Response() (data gin.H, err error) {
	utime := d.param.Get.Get("utime")

	sqlstr := fmt.Sprintf("select event.servername, utime, act, setting from cdn.event left outer join cdn.server on event.servername = server.servername where utime>'%s' order by utime asc", utime)
	records := []CdnEvent{}
	err = d.param.Db.Select(&records, sqlstr)
	if err != nil {
		log.Printf("[ERROR] query last event: %s", err)
		return
	}
	data = gin.H{"records": records}
	return
}

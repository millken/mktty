package dt

import (
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	DtRegister("cdn.getMax", NewCdnGetMax)
}

type CdnGetMax struct {
	param Param
}

func NewCdnGetMax(param Param) (Action, error) {
	return &CdnGetMax{
		param: param,
	}, nil
}

func (d *CdnGetMax) Response() (data gin.H, err error) {
	utime, err := d.getMaxTime()
	if err != nil {
		log.Printf("[ERROR] getMaxTime() : %s", err)
		return
	}
	maxID, countID, err := d.getMaxCount()
	if err != nil {
		log.Printf("[ERROR] getMaxTime() : %s", err)
		return
	}
	log.Printf("[DEBUG] postgres utime=%s, maxID=%d, countID=%d", utime, maxID, countID)
	data = gin.H{
		"utime": utime,
		"maxid": maxID,
		"total": countID,
	}
	return
}

func (d *CdnGetMax) getMaxCount() (maxID, countID int, err error) {
	err = d.param.Db.QueryRow("select max(id) max_id,count(id) count_id from cdn.server").Scan(&maxID, &countID)
	return
}

func (d *CdnGetMax) getMaxTime() (utime string, err error) {
	err = d.param.Db.QueryRow("select max(utime) utime from cdn.event").Scan(&utime)
	return
}

package dt

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

func init() {
	DtRegister("dns.addDomain", NewDnsAddDomain)
}

type DnsAddDomain struct {
	param Param
}

func NewDnsAddDomain(param Param) (Action, error) {
	return &DnsAddDomain{
		param: param,
	}, nil
}

func (d *DnsAddDomain) Response() (data gin.H, err error) {
	var deval []byte
	domain := d.param.Get.Get("domain")
	value := d.param.Get.Get("value")
	deval, err = base64.StdEncoding.DecodeString(value)
	if err != nil {
		return
	}

	sqlstr := "insert into dns.record (domain, value) VALUES (:domain, :value)"
	_, err = d.param.Db.NamedExec(sqlstr,
		map[string]interface{}{
			"domain": domain,
			"value":  string(deval),
		})
	return
}

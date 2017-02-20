package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("cdn.delServer", NewDnsDelDomain)
}

type DnsDelDomain struct {
	param Param
}

func NewDnsDelDomain(param Param) (Action, error) {
	return &DnsDelDomain{
		param: param,
	}, nil
}

func (d *DnsDelDomain) Response() (data gin.H, err error) {
	domain := d.param.Get.Get("domain")

	sqlstr := "delete from dns.record where domain=:domain"
	_, err = d.param.Db.NamedExec(sqlstr,
		map[string]interface{}{
			"domain": domain,
		})
	return
}

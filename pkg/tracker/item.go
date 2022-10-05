package tracker

import (
	"errors"
	"net/url"
	"strings"

	db "github.com/gregoriusongo/price-tracker/pkg/tracker/repo/postgres"
)

var ecommerce db.Ecommerce

// insert product url, if product already exist, return it's id
func InsertUrl(url *url.URL) (id int64, err error){
	// get all ecommerce urls
	ecommerces, err := ecommerce.GetAllEcommerce()
	if err != nil{
		return 
	}

	if len(ecommerces) == 0{
		// no data
		return 0, errors.New("no ecommerce data found")
	}

	var i db.Item
	for _, ec := range ecommerces{
		if strings.Contains(url.Host, ec.SiteUrl){
			// ecommerce found
			i.EcommerceID = int8(ec.ID)
			i.Url = "https://" + url.Host + url.Path
			break
		}
	}
	
	if i.EcommerceID == 0{
		return 0, errors.New("not supported")
	}

	id, err = i.InsertBlankItem()
	if err != nil{
		return 0, err
	}

	return id, err
}
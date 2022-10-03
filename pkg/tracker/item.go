package tracker

import (
	"errors"
	"strings"

	db "github.com/gregoriusongo/price-tracker/pkg/tracker/repo/postgres"
)

var ecommerce db.Ecommerce

// insert product url
func InsertUrl(url string) error{
	// get all ecommerce urls
	ecommerces, err := ecommerce.GetAllEcommerce()
	if err != nil{
		return err
	}

	if len(ecommerces) == 0{
		// no data
		return errors.New("no ecommerce data found")
	}

	var i db.Item
	for _, ec := range ecommerces{
		if strings.Contains(url, ec.SiteUrl){
			// ecommerce found
			i.EcommerceID = int8(ec.ID)
			i.Url = url
			break
		}
	}
	
	if i.EcommerceID == 0{
		return errors.New("not supported")
	}

	if err := i.InsertBlankItem(); err != nil{
		return err
	}

	return nil
}
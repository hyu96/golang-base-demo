package calculate

import (
	"fmt"
	"time"

	pbProduct "github.com/huydq/proto/gen-go/product-service"
)

type Discount struct {
	Type       string
	Value      int
	ValidFrom  null.Time
	ValidUntil null.Time
	// MinimumOrderValue     int
	// MaximumDiscountAmount int
}

func ProductLastPrice(salePrice, quantity int, discount *Discount) (int, error) {
	switch discount.Type {
	case pbProduct.ProductDiscountType_PRODUCT_DISCOUNT_TYPE_PERCENT.String():
		if discount.Value < 0 ||
			discount.Value > 100 { //||
			// discount.Value > discount.MaximumDiscountAmount ||
			// discount.Value < discount.MinimumOrderValue {

			return 0, fmt.Errorf("discount Value out of range")
		}

		validFrom := !discount.ValidFrom.Valid || discount.ValidFrom.Time.Before(time.Now())
		validUntil := !discount.ValidUntil.Valid || discount.ValidUntil.Time.After(time.Now())

		if validFrom && validUntil {
			return quantity * salePrice * (100 - discount.Value) / 100, nil
		}
	case pbProduct.ProductDiscountType_PRODUCT_DISCOUNT_TYPE_VALUE.String():
		if discount.Value < 0 ||
			discount.Value > salePrice { //||
			// discount.Value > discount.MaximumDiscountAmount ||
			// discount.Value < discount.MinimumOrderValue {

			return 0, fmt.Errorf("discount Value out of range")
		}

		return (salePrice - discount.Value) * quantity, nil
	}

	return salePrice * quantity, nil
}

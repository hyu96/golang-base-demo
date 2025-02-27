package service

import (
	"context"
	pbProductMgmt "github.com/huydq/proto/gen-go/product"
	"syreclabs.com/go/faker"
)

// CreateOrder handles order creation
func (uc *ProductService) GetProducts(ctx context.Context, productIdList []int) ([]*pbProductMgmt.Product, error) {
	productList := make([]*pbProductMgmt.Product, len(productIdList))
	for index, id := range productIdList {
		productList[index] = &pbProductMgmt.Product{
			Id:    int32(id),
			Name:  faker.Commerce().ProductName(),
			Price: faker.Commerce().Price(),
		}
	}
	return productList, nil
}

package api

import (
	"encoding/json"
	"net/http"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)
<<<<<<< HEAD
=======
	// _, err := api.AuthMiddleWare(w, req)
	// if err != nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	encoder.Encode(ProductListErrorResponse{Error: err.Error()})
	// 	return
	// }

>>>>>>> c0397392214e368e84db7e7b9a1534ca43781bfb
	response := ProductListSuccessResponse{}
	response.Products = make([]Product, 0)

	products, err := api.productsRepo.SelectAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(DashboardErrorResponse{Error: err.Error()})
			return
		}
	}()

	if err != nil{
		return
	}
<<<<<<< HEAD
	fmt.Println(products)
=======

	// fmt.Println(products)
>>>>>>> c0397392214e368e84db7e7b9a1534ca43781bfb

	for _, product := range products{
		response.Products = append(response.Products, Product{
			Category: product.Category,
			Name: product.ProductName,
			Price: product.Price,
		})
	}
	encoder.Encode(ProductListSuccessResponse{Products: response.Products}) // TODO: replace this
}

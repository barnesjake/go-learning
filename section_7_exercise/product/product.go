package product

type Product struct {
	title string
	id    int
	price float64
}

func New(title string, id int, price float64) Product {
	//if title == "" {
	//	return Product{}, errors.New("title must not be empty.")
	//}
	//if id < 0 {
	//	return Product{}, errors.New("id must not be negative.")
	//}
	//if price < 0 {
	//	return Product{}, errors.New("price must not be negative.")
	//}

	return Product{title: title, id: id, price: price}
}

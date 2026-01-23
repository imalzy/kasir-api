package produk

import "errors"

func List() []Product {
	return ProdukList
}

func Create(newProduk Product) Product {
	newProduk.ID = len(ProdukList) + 1
	ProdukList = append(ProdukList, newProduk)
	return newProduk
}

func Update(id int, updated Product) (*Product, error) {
	for i := range ProdukList {
		if ProdukList[i].ID == id {
			updated.ID = id
			ProdukList[i] = updated
			return &updated, nil
		}
	}

	return nil, errors.New("Produk belum ada")
}

func Delete(id int) error {
	for i, p := range ProdukList {
		if p.ID == id {
			ProdukList = append(ProdukList[:i], ProdukList[i+1:]...)
			return nil
		}
	}
	return errors.New("Produk belum ada")
}

func GetByID(id int) (*Product, error) {
	for _, p := range ProdukList {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("Produk belum ada")
}

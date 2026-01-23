package kategori

import "errors"

func List() []Kategori {
	return KategoriList
}

func Create(newKategori Kategori) Kategori {
	newKategori.ID = len(KategoriList) + 1
	KategoriList = append(KategoriList, newKategori)
	return newKategori
}

func Update(id int, updated Kategori) (*Kategori, error) {
	for i := range KategoriList {
		if KategoriList[i].ID == id {
			updated.ID = id
			KategoriList[i] = updated
			return &updated, nil
		}
	}

	return nil, errors.New("Kategori belum ada")
}

func Delete(id int) error {
	for i, p := range KategoriList {
		if p.ID == id {
			KategoriList = append(KategoriList[:i], KategoriList[i+1:]...)
			return nil
		}
	}
	return errors.New("Kategori belum ada")
}

func GetByID(id int) (*Kategori, error) {
	for _, p := range KategoriList {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("Kategori belum ada")
}

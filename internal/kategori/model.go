package kategori

type Kategori struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
}

var KategoriList = []Kategori{
	{ID: 1, Nama: "Makanan", Deskripsi: "Kategori Makanan"},
	{ID: 2, Nama: "Minuman", Deskripsi: "Kategori Minuman"},
	{ID: 3, Nama: "Cemilan", Deskripsi: "Kategori Cemilan"},
}

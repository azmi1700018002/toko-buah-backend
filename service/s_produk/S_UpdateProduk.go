package s_produk

import (
	"mime/multipart"
	"strconv"
	"toko-buah/model/m_produk"
	"toko-buah/repository/r_produk"

	"github.com/gin-gonic/gin"
)

type UpdateProdukService struct {
	updateProdukRepo r_produk.UpdateProdukRepository
}

func NewUpdateProdukService(updateProdukRepo r_produk.UpdateProdukRepository) *UpdateProdukService {
	return &UpdateProdukService{updateProdukRepo}
}

func (s *UpdateProdukService) UpdateProduk(ctx *gin.Context, produk m_produk.Produk, gambarFile *multipart.FileHeader) (*m_produk.Produk, error) {
	// Mendapatkan ID produk dari parameter route
	idProdukStr := ctx.Param("produk_id")

	// Konversi tipe data ID produk dari string ke int
	idProduk, err := strconv.Atoi(idProdukStr)
	if err != nil {
		return nil, err
	}

	// Set ID produk ke dalam struct produk
	produk.ProdukID = idProduk

	// Memanggil repository untuk update produk
	updatedProduk, err := s.updateProdukRepo.UpdateProduk(&produk, gambarFile)
	if err != nil {
		return nil, err
	}

	return updatedProduk, nil
}

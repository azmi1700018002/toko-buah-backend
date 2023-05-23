package s_produk

import (
	"toko-buah/model/m_produk"
	"toko-buah/repository/r_produk"

	"github.com/gin-gonic/gin"
)

type AddProdukService struct {
	addProdukRepo r_produk.AddProdukRepository
}

func NewAddProdukService(addProdukRepo r_produk.AddProdukRepository) *AddProdukService {
	return &AddProdukService{addProdukRepo}
}

func (s *AddProdukService) AddProduk(ctx *gin.Context, produk m_produk.Produk) (*m_produk.Produk, error) {
	gambarFile, err := ctx.FormFile("gambar")
	if err != nil {
		return nil, err
	}

	updatedProduct, err := s.addProdukRepo.AddProduk(&produk, gambarFile)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

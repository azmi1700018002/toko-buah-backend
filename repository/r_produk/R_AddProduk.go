package r_produk

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"toko-buah/config/db"
	"toko-buah/model/m_produk"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

type AddProdukRepository interface {
	AddProduk(produk *m_produk.Produk, gambarFile *multipart.FileHeader) (*m_produk.Produk, error)
}

type addProdukRepository struct{}

func NewAddProdukRepository() AddProdukRepository {
	return &addProdukRepository{}
}

func (r *addProdukRepository) AddProduk(produk *m_produk.Produk, gambarFile *multipart.FileHeader) (*m_produk.Produk, error) {
	// Generate unique filename using UUID
	fileExt := filepath.Ext(gambarFile.Filename)
	newFileName := uuid.New().String() + fileExt

	// Open the source file
	src, err := gambarFile.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Determine the destination path
	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/"
	savePath := filepath.Join(basePath, newFileName)
	// savePath := "./uploads/" + newFileName
	dst, err := os.Create(savePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(dst, src)
	if err != nil {
		return nil, err
	}

	// Open the saved image
	img, err := imaging.Open(savePath)
	if err != nil {
		return nil, err
	}

	// Resize the image to 500x500
	resizedImg := imaging.Resize(img, 500, 500, imaging.Lanczos)

	// Save the resized image
	err = imaging.Save(resizedImg, savePath)
	if err != nil {
		return nil, err
	}

	// Menyimpan di database hanya nama file saja
	produk.Gambar = newFileName

	// Simpan produk ke database
	err = db.Server().Create(&produk).Error
	if err != nil {
		return nil, err
	}

	return produk, nil
}

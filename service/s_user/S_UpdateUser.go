package s_user

import (
	"mime/multipart"
	"strings"

	// "strings"
	"toko-buah/model/m_user"
	"toko-buah/repository/r_user"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUserService struct {
	updateUserRepo r_user.UpdateUserRepository
}

func NewUpdateUserService(updateUserRepo r_user.UpdateUserRepository) *UpdateUserService {
	return &UpdateUserService{updateUserRepo}
}

func (s *UpdateUserService) UpdateUser(ctx *gin.Context, user m_user.User, file *multipart.FileHeader) (*m_user.User, error) {
	// Mendapatkan ID user dari parameter route
	idUser := ctx.Param("id_user")

	// Set ID user ke dalam struct user
	user.UserID = uuid.MustParse(idUser)

	// cek apakah file foto profil di-upload
	if file != nil {
		// upload file ke Cloudinary
		cloudinaryConfig, err := cloudinary.NewFromParams("ddee7paye", "898949133356251", "Jn3rtgch_6Api6XU5BWmvBUMsuA")
		if err != nil {
			return nil, err
		}
		// convert file ke format yang bisa diupload ke cloudinary
		fileReader, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer fileReader.Close()
		uploadParams := uploader.UploadParams{Format: "jpg"}
		uploadResult, err := cloudinaryConfig.Upload.Upload(ctx, fileReader, uploadParams)
		if err != nil {
			return nil, err
		}

		// hapus gambar lama dari Cloudinary
		if strings.HasPrefix(user.ProfilePicture, "http://res.cloudinary.com") {
			publicID := strings.TrimPrefix(user.ProfilePicture, "http://res.cloudinary.com/ddee7paye/image/upload/")
			publicID = strings.TrimSuffix(publicID, ".jpg")
			_, err = cloudinaryConfig.Admin.DeleteAssetsByPrefix(ctx, admin.DeleteAssetsByPrefixParams{
				Prefix: []string{publicID},
			})
			if err != nil {
				return nil, err
			}
		}

		user.ProfilePicture = uploadResult.URL
	} else {
		// jika file tidak diupload, set nilai default untuk gambar profil
		user.ProfilePicture = "https://www.pngall.com/wp-content/uploads/12/Avatar-No-Background.png"
	}

	// Memanggil repository untuk update user
	updatedUser, err := s.updateUserRepo.UpdateUser(&user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

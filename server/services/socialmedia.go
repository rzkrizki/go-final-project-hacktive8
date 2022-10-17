package services

import (
	"errors"
	"final-project/server/controllers/view"
	"final-project/server/repositories"
	"final-project/server/repositories/models"
	"final-project/server/request"
)

type SocmedService struct {
	socmedRepo repositories.SocialMediaRepo
}

func NewSocialMediaService(socmedRepo repositories.SocialMediaRepo) *SocmedService {
	return &SocmedService{socmedRepo: socmedRepo}
}

func (s *SocmedService) Create(req *request.CreateSocialMedia, id int) (view.ResponseCreateSocmed, error) {
	var socmed models.SocialMedia

	socmed.Name = req.Name
	socmed.SocialMediaUrl = req.SocialMediaUrl
	socmed.UserId = id

	data, err := s.socmedRepo.Create(&socmed)

	if err != nil {
		return view.ResponseCreateSocmed{}, err
	}

	return view.ResponseCreateSocmed{
		Id:             data.Id,
		Name:           data.Name,
		SocialMediaUrl: data.SocialMediaUrl,
		UserId:         data.UserId,
		CreatedAt:      data.CreatedAt,
	}, nil
}

func (s *SocmedService) Get(idUser int) (view.ReturnGetSocmed, error) {
	data, err := s.socmedRepo.GetSocmedByUserId(idUser)

	if err != nil {
		return view.ReturnGetSocmed{}, err
	}

	if len(data) == 0 {
		return view.ReturnGetSocmed{}, errors.New("Social Media Not Found")
	}

	var response []view.ResponseGetSocmed

	for _, v := range data {
		response = append(response, view.ResponseGetSocmed{
			Id:             v.Id,
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			UserId:         v.UserId,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			User: view.ResponseWithUserIdSocmed{
				Id:              v.User.Id,
				Username:        v.User.Username,
				ProfileImageUrl: "",
			},
		})
	}

	return view.ReturnGetSocmed{
		SocialMedia: response,
	}, nil
}

func (s *SocmedService) Update(req *request.UpdateSocialMedia, idUser, socmedId int) (view.ResponseUpdateSocmed, error) {
	var socmed models.SocialMedia

	socmed.Name = req.Name
	socmed.SocialMediaUrl = req.SocialMediaUrl

	checkIfExist, _ := s.socmedRepo.CheckSocmedByIdAndUserId(socmedId, idUser)

	if !checkIfExist {
		return view.ResponseUpdateSocmed{}, errors.New("Unauthorized")
	}

	data, err := s.socmedRepo.UpdateSocialMediaById(socmedId, &socmed)

	if err != nil {
		return view.ResponseUpdateSocmed{}, err
	}

	return view.ResponseUpdateSocmed{
		Id:             data.Id,
		Name:           data.Name,
		SocialMediaUrl: data.SocialMediaUrl,
		UserId:         data.UserId,
		UpdatedAt:      data.UpdatedAt,
	}, nil
}

func (s *SocmedService) Delete(idUser, idsocmed int) (view.ResponseDeleteSocmed, error) {
	checkIfExist, _ := s.socmedRepo.CheckSocmedByIdAndUserId(idsocmed, idUser)

	if !checkIfExist {
		return view.ResponseDeleteSocmed{}, errors.New("Unauthorized")
	}

	err := s.socmedRepo.DeleteSocialMediaById(idsocmed)

	if err != nil {
		return view.ResponseDeleteSocmed{}, err
	}

	return view.ResponseDeleteSocmed{
		Message: "Your social media has been successfully deleted",
	}, nil
}

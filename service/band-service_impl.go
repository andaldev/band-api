package service

import (
	"band_api/data/request"
	"band_api/data/response"
	"band_api/helper"
	"band_api/model"
	"band_api/repository"

	"github.com/go-playground/validator/v10"
)

type BandServiceImpl struct {
	BandRepository repository.BandRepository
	Validate       *validator.Validate
}

func NewBandServiceImpl(bandRepository repository.BandRepository, validate *validator.Validate) BandService {
	return &BandServiceImpl{
		BandRepository: bandRepository,
		Validate:       validate,
	}
}

// Create implements BandService.
func (b *BandServiceImpl) Create(band request.CreateBandRequest) {
	err := b.Validate.Struct(band)
	helper.ErrorPanic(err)
	bandModel := model.Band{
		Name: band.Name,
	}
	b.BandRepository.Save(bandModel)
}

// Delete implements BandService.
func (b *BandServiceImpl) Delete(bandId int) {
	b.BandRepository.Delete(bandId)
}

// FindAll implements BandService.
func (b *BandServiceImpl) FindAll() []response.BandResponse {
	result := b.BandRepository.FindAll()

	var bands []response.BandResponse
	for _, value := range result {
		band := response.BandResponse{
			Name: value.Name,
		}
		bands = append(bands, band)
	}

	return bands
}

// FindById implements BandService.
func (b *BandServiceImpl) FindById(bandId int) response.BandResponse {
	bandData, err := b.BandRepository.FindById(bandId)
	helper.ErrorPanic(err)

	bandResponse := response.BandResponse{
		Name: bandData.Name,
	}

	return bandResponse
}

// Update implements BandService.
func (b *BandServiceImpl) Update(band request.UpdateBandRequest) {
	bandData, err := b.BandRepository.FindById(band.Id)
	helper.ErrorPanic(err)
	bandData.Name = band.Name
	b.BandRepository.Update(bandData)
}

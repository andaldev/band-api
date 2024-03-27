package service

import (
	"band_api/data/request"
	"band_api/model"
)

type BandService interface {
	Create(band request.CreateBandRequest)
	Update(band request.UpdateBandRequest)
	Delete(bandId int)
	FindById(bandId int) model.Band
	FindAll() []model.Band
}

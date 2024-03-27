package service

import (
	"band_api/data/request"
	"band_api/data/response"
)

type BandService interface {
	Create(band request.CreateBandRequest)
	Update(band request.UpdateBandRequest)
	Delete(bandId int)
	FindById(bandId int) response.BandResponse
	FindAll() []response.BandResponse
}

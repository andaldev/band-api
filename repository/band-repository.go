package repository

import (
	"band_api/data/request"
	"band_api/helper"
	"band_api/model"
	"errors"

	"gorm.io/gorm"
)

type BandRepositoryImpl struct {
	Db *gorm.DB
}

func NewBandRepositoryImpl(Db *gorm.DB) BandRepository {
	return &BandRepositoryImpl{Db: Db}
}

// FindAll implements BandRepository.
func (b *BandRepositoryImpl) FindAll() []model.Band {
	var band []model.Band
	result := b.Db.Find(&band)
	helper.ErrorPanic(result.Error)
	return band
}

// FindById implements BandRepository.
func (b *BandRepositoryImpl) FindById(bandId int) (band model.Band, err error) {
	var uniqueBand model.Band
	result := b.Db.Find(&uniqueBand, bandId)
	if result != nil {
		return uniqueBand, nil
	} else {
		return uniqueBand, errors.New("band not found")
	}
}

// Save implements BandRepository.
func (b *BandRepositoryImpl) Save(band model.Band) {
	result := b.Db.Create(&band)
	helper.ErrorPanic(result.Error)
}

// Update implements BandRepository.
func (b *BandRepositoryImpl) Update(band model.Band) {
	var updatedBand = request.UpdateBandRequest{
		Name: band.Name,
	}
	result := b.Db.Model(&band).Updates(updatedBand)
	helper.ErrorPanic(result.Error)
}

// delete implements BandRepository.
func (b *BandRepositoryImpl) Delete(bandId int) {
	var band []model.Band
	result := b.Db.Where("id = ?", bandId).Delete(&band)
	helper.ErrorPanic(result.Error)
}

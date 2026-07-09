package api

import (
	"test/3-struct/bins"
	"test/3-struct/storage"
)

type API struct {
	storage storage.Storage
}

func NewApi(s storage.Storage) *API {
	return &API{
		storage: s,
	}
}

func (a *API) CreateBin(bin bins.Bin) error {
	return a.storage.SaveBin(bin)
}

func (a *API) GetBin(id string) (bins.Bin, error) {
	bin, err := a.storage.GetBin(id)
	if err != nil {
		return bins.Bin{}, err
	}
	return bin, nil
}

func (a *API) DeleteBin(id string) error {
	return a.storage.DeleteBin(id)
}

func (a *API) GetAll() (bins.BinList, error) {
	allBins, err := a.storage.GetAllBins()
	if err != nil {
		return bins.BinList{}, err
	}
	return allBins, nil
}

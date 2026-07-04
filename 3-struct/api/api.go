package api

import (
	"encoding/json"
	"test/3-struct/bins"
	"test/3-struct/config"
	"test/3-struct/storage"
)

type API struct {
	storage storage.Storage
	Key     config.Config
}

func NewApi(s storage.Storage, config config.Config) *API {
	return &API{
		storage: s,
	}
}

func (a *API) SaveBin(bin *bins.Bin) error {
	data, _ := json.Marshal(bin)
	_, err := a.storage.Write(data)
	return err
}

func (a *API) LoadBins() (*bins.BinList, error) {
	data, err := a.storage.Read()
	if err != nil {
		return nil, err
	}
	var binList bins.BinList
	json.Unmarshal(data, &binList)
	return &binList, nil
}

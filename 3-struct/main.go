package main

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bin []Bin
}

func NewBin(id, name string, private bool) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
func NewBinList(bins ...Bin) *BinList {

	return &BinList{
		Bin: bins,
	}

}

func main() {

}

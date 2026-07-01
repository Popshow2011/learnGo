package main

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	Bin []Bin
}

func NewBin(id, name string, private bool) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}
func NewBinList(bins ...Bin) *BinList {

	return &BinList{
		Bin: bins,
	}

}

func main() {

}

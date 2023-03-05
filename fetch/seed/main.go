package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-rel/rel"
	"github.com/go-rel/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type EfisheryData struct {
	Uuid         string `json:"uuid" db:"uuid,primary"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         int    `json:"size"`
	Price        int    `json:"price"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

func main() {
	adapter, err := sqlite3.Open("fetch.db")
	if err != nil {
		panic(err)
	}
	repo := rel.New(adapter)

	ctx := context.Background()

	json_file, err := os.Open("efishery-data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()
	byteValue, _ := io.ReadAll(json_file)
	var efisheryData []EfisheryData

	json.Unmarshal(byteValue, &efisheryData)

	for i := 0; i < len(efisheryData); i++ {
		repo.Insert(ctx, &efisheryData[i])
	}
}

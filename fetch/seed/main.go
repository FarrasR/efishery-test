package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Uuid         string `json:"uuid" db:"uuid,primary"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

func main() {
	adapter, err := sqlite3.Open("product.db")
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
	var products []Product

	json.Unmarshal(byteValue, &products)

	for i := 0; i < len(products); i++ {
		dateString := products[i].TglParsed
		// fmt.Println(dateString)
		rawTime, _ := time.Parse(time.RFC3339, dateString)
		// fmt.Println(rawTime)
		products[i].TglParsed = rawTime.Format("2006-01-02 15:04:05.000")
		// fmt.Println(products[i].TglParsed)
		repo.Insert(ctx, &products[i])
	}
}

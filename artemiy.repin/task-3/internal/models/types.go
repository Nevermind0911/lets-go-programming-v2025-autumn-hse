package models

import "encoding/xml"

type InputValute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type InputData struct {
	XMLName xml.Name      `xml:"ValCurs"`
	Items   []InputValute `xml:"Valute"`
}

type FinalValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type Currencies []FinalValute

func (c Currencies) Len() int           { return len(c) }
func (c Currencies) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Currencies) Less(i, j int) bool { return c[i].Value > c[j].Value }

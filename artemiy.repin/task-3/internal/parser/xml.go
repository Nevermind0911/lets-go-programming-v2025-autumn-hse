package parser

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Nevermind0911/task-3/internal/models"
	"golang.org/x/net/html/charset"
)

func ReadAndConvert(path string) (models.Currencies, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening XML file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("err while closing file: %v", err))
		}
	}()

	var rawData models.InputData

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&rawData); err != nil {
		return nil, fmt.Errorf("XML structure err: %w", err)
	}

	result := make(models.Currencies, 0, len(rawData.Items))

	for _, item := range rawData.Items {
		valStr := strings.Replace(item.Value, ",", ".", 1)

		valFloat, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return nil, fmt.Errorf("val number err %s: %w", item.CharCode, err)
		}

		result = append(result, models.FinalValute{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    valFloat,
		})
	}

	return result, nil
}

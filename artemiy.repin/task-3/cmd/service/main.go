package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/Nevermind0911/task-3/internal/config"
	"github.com/Nevermind0911/task-3/internal/parser"
	"github.com/Nevermind0911/task-3/internal/writer"
)

func main() {
	cfgPath := flag.String("config", "", "Путь к YAML конфигурации")
	flag.Parse()

	if *cfgPath == "" {
		panic("Флаг -config обязателен")
	}

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		panic(fmt.Sprintf("Ошибка загрузки конфига: %v", err))
	}

	data, err := parser.ReadAndConvert(cfg.SourceFile)
	if err != nil {
		panic(fmt.Sprintf("Ошибка обработки XML: %v", err))
	}

	sort.Sort(data)

	if err := writer.SaveJSON(cfg.TargetFile, data); err != nil {
		panic(fmt.Sprintf("Ошибка сохранения: %v", err))
	}

	fmt.Println("Готово! Результат в", cfg.TargetFile)
}

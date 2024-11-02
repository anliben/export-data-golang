package utils

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func CleanOldFiles(directory string) error {
	expiration := 24 * time.Hour

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && time.Since(info.ModTime()) > expiration {
			err := os.Remove(path)
			if err != nil {
				log.Printf("Erro ao apagar o arquivo %s: %v", path, err)
				return err
			}
			log.Printf("Arquivo %s apagado com sucesso", path)
		}
		return nil
	})

	return err
}

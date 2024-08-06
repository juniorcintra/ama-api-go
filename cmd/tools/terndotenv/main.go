package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Definir o comando a ser executado
	cmd := exec.Command("tern", "migrate", "--migrations", "./internal/store/pgstore/migrations", "--config", "./internal/store/pgstore/migrations/tern.conf")

	// Executar o comando e capturar a saída
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao executar o comando: %v\nSaída: %s", err, string(output))
	}

	// Imprimir a saída do comando
	fmt.Printf("Saída do comando: %s\n", string(output))
}
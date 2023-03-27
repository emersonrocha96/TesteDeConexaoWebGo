package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Monitoramento básico de uma site web

func main() {

	agora := time.Now() // Coleta o momento da consulta ao site
	url := os.Args[1]  // parametro de url do site a ser monitorado, exemplo: https://www.google.com

	get, err := http.Get(url) // Pega a informação da URL e retorna um valor X ou algum erro relacionado a consulta.

	if err != nil { // Se a mensagem de erro for diferente de nullo, então o bloco abaixo é executado. 
		
		fmt.Println("Ocorreu um erro ao executar o get(url)")
	    panic(err) // Utilizado a modo de parar imediatamente a aplicação já que um erro ocorreu.
		
	}

	decorrido := time.Since(agora).Seconds() // Será feito um calculo de quantos segundos foram decorridos desde o momento da coleta
	status := get.StatusCode

	fmt.Printf("Status: [%d] | Tempo de Carga: [%f]\n", status, decorrido) // Formatador %d é usado para exibir dados inteiros
								              // Formatador %f é usado para um valor flutuante. 

}

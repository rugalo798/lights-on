package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lights-on/service"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	URL := os.Getenv("URL")

	var powerOn service.Request = service.Request{
		Header: service.Header{
			Name:           "turnOnOff",
			Namespace:      "control",
			PayloadVersion: 1,
		},
		Payload: service.Payload{
			AccessToken: os.Getenv("ACCESSTOKEN"),
			DevID:       os.Getenv("DEVID"),
			Value:       1,
		},
	}

	var powerOff service.Request = service.Request{
		Header: service.Header{
			Name:           "turnOnOff",
			Namespace:      "control",
			PayloadVersion: 1,
		},
		Payload: service.Payload{
			AccessToken: os.Getenv("ACCESSTOKEN"),
			DevID:       os.Getenv("DEVID"),
			Value:       0,
		},
	}

	powerOnReq, err := json.Marshal(powerOn)

	if err != nil {
		log.Fatalf("Erro: %v", err.Error())
	}
	powerOffReq, err := json.Marshal(powerOff)

	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	rootCmd := &cobra.Command{
		Use:   "lights",
		Short: "Um app para ligar e delisgar minha luz",
		Long:  "Um app desenvolvido com golang e cobra para desligar e ligar minha luz smart",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(os.Getenv("ACCESSTOKEN"))
		},
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "on",
		Short: "Comando para ligar as luzes do quarto",
		Run: func(cmd *cobra.Command, args []string) {
			req, err := http.NewRequest("POST", URL, bytes.NewBuffer(powerOnReq))
			if err != nil {
				log.Fatalf("Erro: %v", err)
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()

		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "off",
		Short: "Comando para desligar as luzes do quarto",
		Run: func(cmd *cobra.Command, args []string) {
			req, err := http.NewRequest("POST", "https://px1.tuyaus.com/homeassistant/skill", bytes.NewBuffer(powerOffReq))

			if err != nil {
				log.Fatalf("Algum erro ocorreu: %v\n", err.Error())
			}

			client := &http.Client{}

			resp, err := client.Do(req)

			if err != nil {
				log.Fatalf("Erro: %v", err.Error())
			}

			defer resp.Body.Close()
		},
	})

	err = rootCmd.Execute()

	if err != nil {
		log.Fatalf("Algo deu errado: %v", err)
	}
}

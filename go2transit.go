package go2transit

import (
	"net/http"
	"time"
)

type GoTransitClient struct {
	Debug    bool
	Language string
	Client   *http.Client
}

// const VERSION = "0.0.1"

func New() (g *GoTransitClient, err error) {
	g = &GoTransitClient{
		Debug:    true,
		Language: "en",
		Client:   &http.Client{Timeout: (20 * time.Second)},
	}
	return
}

func Nouveau() (g *GoTransitClient, err error) {
	g = &GoTransitClient{
		Debug:    true,
		Language: "fr",
		Client:   &http.Client{Timeout: (20 * time.Second)},
	}
	return
}

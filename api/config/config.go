package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	DataSourceURL string
	BaseURL       string
	Port          int
)

func init() {
	DataSourceURL = os.Getenv("DATA_SOURCE_URL")
	BaseURL = os.Getenv("BASE_URL")

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("unable to convert PORT [%s]", portStr))
	}
	Port = port
}

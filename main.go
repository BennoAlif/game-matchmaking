package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func main() {
	fmt.Println("Gorm is Installed")
}
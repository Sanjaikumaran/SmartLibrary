package main

import (
	"fmt"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver"
)

func main() {
    // Call functions from the dbdriver package
    // Assuming you have functions like Connect and Disconnect in the dbdriver package
    err := dbdriver.Connect()
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return
    }

    defer func() {
        err := dbdriver.Disconnect()
        if err != nil {
            fmt.Println("Error disconnecting from the database:", err)
        }
    }()

    fmt.Println("Connected to the database")
    // Add your remaining code here
}

package main

import (
    "encoding/base64"
    "fmt"
    "github.com/btcsuite/btcutil/bech32"
)

func main() {
    // Example Bech32 address
    bech32Address := "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllst77y4l"

    // Step 1: Decode the Bech32 address
    _, decodedBytes, err := bech32.Decode(bech32Address)
    if err != nil {
        fmt.Println("Error decoding Bech32 address:", err)
        return
    }

    // Step 2: Encode the decoded bytes as base64
    base64String := base64.StdEncoding.EncodeToString(decodedBytes)

    // Step 3: Print the base64 representation
    fmt.Println("Base64:", base64String)
}
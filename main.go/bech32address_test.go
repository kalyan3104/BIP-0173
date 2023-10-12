package main

import (
 "encoding/base64"
 "fmt"
 "github.com/btcsuite/btcutil/bech32"
)

func main() {
 // Base64-encoded string
 base64Encoded := "AAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABHx8fEA=="
 hrp := "moa" // Specify the desired HRP

 // Step 1: Decode the base64-encoded string
 decodedData, err := base64.StdEncoding.DecodeString(base64Encoded)
 if err != nil {
  fmt.Println("Error decoding base64:", err)
  return
 }

 // Step 2: Encode the decoded binary data as Bech32 with the specified HRP
 bech32Address, err := bech32.Encode(hrp, decodedData)
 if err != nil {
  fmt.Println("Error encoding Bech32:", err)
  return
 }

 // Print the Bech32 address
 fmt.Println("Bech32 Address:", bech32Address)
}
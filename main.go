package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/hyperledger/fabric-sdk-go/pkg/gateway"
    "log"
    "net/http"
)

// Asset represents the asset structure used in the API
type Asset struct {
    DealerID    string `json:"dealerId"`
    MSISDN      string `json:"msisdn"`
    MPIN        string `json:"mpin"`
    Balance     int    `json:"balance"`
    Status      string `json:"status"`
    TransAmount int    `json:"transAmount"`
    TransType   string `json:"transType"`
    Remarks     string `json:"remarks"`
}

// CreateAsset handles the POST request to create a new asset
func CreateAsset(w http.ResponseWriter, r *http.Request) {
    // Initialize the Gateway
    gw, err := gateway.Connect(gateway.WithConfig("connection.yaml"), gateway.WithIdentity("Org1", "User1"))
    if err != nil {
        http.Error(w, "Failed to connect to gateway: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer gw.Close()

    // Get the contract from the gateway
    contract := gw.GetNetwork("mychannel").GetContract("asset-management")

    // Decode the JSON request body
    var asset Asset
    err = json.NewDecoder(r.Body).Decode(&asset)
    if err != nil {
        http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Create the asset on the blockchain
    _, err = contract.SubmitTransaction("CreateAsset", asset.DealerID, asset.MSISDN, asset.MPIN, fmt.Sprintf("%d", asset.Balance), asset.Status, fmt.Sprintf("%d", asset.TransAmount), asset.TransType, asset.Remarks)
    if err != nil {
        http.Error(w, "Failed to create asset: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(asset)
}

// QueryAsset handles the GET request to retrieve an asset by dealer ID
func QueryAsset(w http.ResponseWriter, r *http.Request) {
    // Initialize the Gateway
    gw, err := gateway.Connect(gateway.WithConfig("connection.yaml"), gateway.WithIdentity("Org1", "User1"))
    if err != nil {
        http.Error(w, "Failed to connect to gateway: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer gw.Close()

    // Get the contract from the gateway
    contract := gw.GetNetwork("mychannel").GetContract("asset-management")

    // Get the dealer ID from the URL
    dealerID := mux.Vars(r)["dealerId"]

    // Query the asset from the blockchain
    result, err := contract.EvaluateTransaction("QueryAsset", dealerID)
    if err != nil {
        http.Error(w, "Failed to query asset: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the asset details
    var asset Asset
    err = json.Unmarshal(result, &asset)
    if err != nil {
        http.Error(w, "Failed to unmarshal asset: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(asset)
}

// UpdateAsset handles the PUT request to update an existing asset
func UpdateAsset(w http.ResponseWriter, r *http.Request) {
    // Initialize the Gateway
    gw, err := gateway.Connect(gateway.WithConfig("connection.yaml"), gateway.WithIdentity("Org1", "User1"))
    if err != nil {
        http.Error(w, "Failed to connect to gateway: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer gw.Close()

    // Get the contract from the gateway
    contract := gw.GetNetwork("mychannel").GetContract("asset-management")

    // Get the dealer ID from the URL
    dealerID := mux.Vars(r)["dealerId"]

    // Decode the JSON request body
    var asset Asset
    err = json.NewDecoder(r.Body).Decode(&asset)
    if err != nil {
        http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Update the asset on the blockchain
    _, err = contract.SubmitTransaction("UpdateAsset", dealerID, fmt.Sprintf("%d", asset.Balance), asset.Status)
    if err != nil {
        http.Error(w, "Failed to update asset: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// main function sets up the HTTP server and routes
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/assets", CreateAsset).Methods("POST")
    router.HandleFunc("/assets/{dealerId}", QueryAsset).Methods("GET")
    router.HandleFunc("/assets/{dealerId}", UpdateAsset).Methods("PUT")

    log.Println("Starting REST API server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}

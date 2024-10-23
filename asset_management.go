package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Asset defines the structure for an asset
type Asset struct {
    DealerID    string `json:"dealerId"`    // Unique identifier for the dealer
    MSISDN      string `json:"msisdn"`      // Mobile number associated with the asset
    MPIN        string `json:"mpin"`        // MPIN for the asset
    Balance     int    `json:"balance"`     // Current balance of the asset
    Status      string `json:"status"`      // Status of the asset (e.g., active, inactive)
    TransAmount int    `json:"transAmount"` // Transaction amount for the asset
    TransType   string `json:"transType"`   // Type of transaction (e.g., credit, debit)
    Remarks     string `json:"remarks"`     // Additional remarks related to the asset
}

// SmartContract defines the structure of the smart contract
type SmartContract struct {
    contractapi.Contract
}

// CreateAsset creates a new asset and stores it in the ledger
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealerID string, msisdn string, mpin string, balance int, status string, transAmount int, transType string, remarks string) error {
    asset := Asset{
        DealerID:    dealerID,
        MSISDN:      msisdn,
        MPIN:        mpin,
        Balance:     balance,
        Status:      status,
        TransAmount: transAmount,
        TransType:   transType,
        Remarks:     remarks,
    }

    // Convert asset struct to JSON format
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return fmt.Errorf("failed to marshal asset: %v", err)
    }

    // Store the asset in the ledger using the dealer ID as the key
    return ctx.GetStub().PutState(dealerID, assetJSON)
}

// QueryAsset retrieves an asset by its dealer ID
func (s *SmartContract) QueryAsset(ctx contractapi.TransactionContextInterface, dealerID string) (*Asset, error) {
    // Retrieve the asset JSON from the ledger
    assetJSON, err := ctx.GetStub().GetState(dealerID)
    if err != nil {
        return nil, fmt.Errorf("failed to read asset: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("asset %s does not exist", dealerID)
    }

    // Unmarshal JSON back to Asset struct
    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal asset: %v", err)
    }

    return &asset, nil
}

// UpdateAsset updates the balance and status of an existing asset
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, dealerID string, balance int, status string) error {
    // Retrieve the existing asset
    asset, err := s.QueryAsset(ctx, dealerID)
    if err != nil {
        return err
    }

    // Update the asset's balance and status
    asset.Balance = balance
    asset.Status = status

    // Convert updated asset to JSON and store it back in the ledger
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return fmt.Errorf("failed to marshal asset: %v", err)
    }

    return ctx.GetStub().PutState(dealerID, assetJSON)
}

// Main function to start the smart contract
func main() {
    chaincode, err := contractapi.NewChaincode(new(SmartContract))
    if err != nil {
        fmt.Printf("Error creating asset-management chaincode: %v", err)
        return
    }

    // Start the chaincode
    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting asset-management chaincode: %v", err)
    }
}

#!/bin/bash

# This script sets up a Hyperledger Fabric test network.

# Step 1: Clone the fabric-samples repository from GitHub
# This repository contains sample code and scripts for working with Hyperledger Fabric.
git clone https://github.com/hyperledger/fabric-samples.git

# Step 2: Change directory to the test-network folder
cd fabric-samples/test-network

# Step 3: Start the test network
# This command will create and start the Docker containers needed for the Fabric network.
./network.sh up

# Step 4: Create a channel
# A channel is a private subnet of communication between the network members. 
# This command creates a channel called 'mychannel' where assets will be managed.
./network.sh createChannel

# Step 5: Deploy the chaincode
# Chaincode is the smart contract that contains the business logic for managing assets.
# This command deploys the sample asset transfer chaincode to the channel created above.
./network.sh deployCC

# At this point, the test network is set up and the chaincode is deployed.
echo "Hyperledger Fabric test network is up and running with the asset transfer chaincode deployed."

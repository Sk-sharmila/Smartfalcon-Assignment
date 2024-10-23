Hyperledger Fabric Asset Management System
This project is about creating a system to manage assets using Hyperledger Fabric, which is a type of blockchain technology. It helps a financial institution keep track of assets securely and transparently.

What’s Inside the Project?
The project has a few important files:

===>setup_hyperledger_fabric_network.sh: This is a script that helps you set up a test network for Hyperledger Fabric. It grabs the necessary files, starts everything up, creates a channel for communication, and deploys the smart contracts.

===>asset_management.go: This file contains the smart contract code. It describes what an asset is and includes functions to create new assets, look up existing ones, and update their details. Each asset has information like DealerID, MSISDN, MPIN, Balance, and more.

===>main.go: This is where the REST API lives. It allows users to interact with the smart contracts through web requests. There are routes to create new assets, get information about an asset using its dealer ID, and update an asset's balance and status.

===>Dockerfile: This file tells Docker how to build the application into a container. It specifies the base image, sets up the working directory, installs any necessary packages, compiles the application, and explains how to run it.

===>connection.yaml: This is a configuration file that helps the REST API connect to the Hyperledger Fabric network. It includes details about the organization, peers, channels, and other connection settings.

How Does It Work?
Set Up the Network: First, run the setup_hyperledger_fabric_network.sh script to start everything. This will set up the Hyperledger Fabric network, create a channel called mychannel, and deploy the asset management smart contract.

Create Assets: To add a new asset, you can send a POST request to the /assets endpoint of the REST API with the asset details like DealerID and Balance.

Get Assets: If you want to find out about a specific asset, just send a GET request to the /assets/{dealerId} endpoint. It will return the details for that asset.

Update Assets: To change an existing asset’s balance or status, send a PUT request to /assets/{dealerId} with the new information.

Getting Started
To get started, clone the project to your computer. Go to the project folder and run the setup script to initialize the network. Then, build the Docker image and run the REST API service. You can use tools like Postman or Curl to test the API.

In Summary
This project shows how to use blockchain technology to manage assets effectively. By using Hyperledger Fabric, we make sure that all asset transactions are recorded safely and transparently, which builds trust among everyone involved.

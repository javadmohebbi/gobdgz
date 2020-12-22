# Golang-Bitdefender GravityZone API
**gobdgz** go-module helps GravityZone API developers who are using Google GoLang integrate their apps easier.

### Current Integrations
Here is the List of integrated features
- **Accounts**: You can read [this guide](http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf#page=11&zoom=100,33,112) for more information
    - **getAccountsList**: Get list of all Control Center accounts
    - **createAccount**: Create a new Control Center account
    - **deleteAccount**: Delete Control Center accounts using accountId
    - **updateAccount**: Update Control Center accounts using accountId
    - **getNotificationsSettings**: Get information about notifications settings for an account
    - **configureNotificationsSettings**: Configure/Update notifications settings for an account

### Future integrations
This list will be available in the future releases:
- **Network**
    - **getContainers**: Returns the network containers.
    - **getNetworkInventoryItems**: Returns network inventory items.
    - **createScanTask**: Returns true if the task was successfully created.
    - **createReconfigureClientTask**: Creates a new Reconfigure Client task.
    - **getScanTasksList**: Returns the list of scan tasks.
    - **getEndpointsList**: Returns the list of endpoints.
    - **getManagedEndpointDetails**: Returns the details about a managed endpoints
    - **createCustomGroup**: Creates a new group under an existing one or under Computers and Groups.
    - **deleteCustomGroup**: Deletes a custom group.
    - **moveCustomGroup**: Moves a custom group under another custom group.
    - **moveEndpoints**: Moves the specified list of endpoints to a custom group.
    - **deleteEndpoint**: Deletes a specified endpoint.
    - **setEndpointLabel**: Sets a label to an endpoint.
    - **createScanTaskByMac**: Generates scan tasks for managed endpoints identified by MAC address.
    - **assignPolicy**: This method is used to assign a policy template on the specified endpoints or containers.


# Get the module
Run this command to install **gobdgz** module on your system
```
go get github.com/javadmohebbi/gobdgz
```


# Documentation
If you are interested in reading this document, you might want to integrate **Bitdefender GravityZone** with your infrastructure. This module have been developing based on **Bitdefender GravityZone On-Premises API Guide** which is available in [this link](http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf)

### Quick Start Guide
1. Get API Keys
- An API Key will help you to communicate with GravityZone API Server. To generate an API key you can read [this document](http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf#page=7&zoom=100,33,85)
- We RECOMMEND you to create a dedicated account for using APIs
2. Get **gobdgz** module
    - Run this command to get the module: ```go get github.com/javadmohebbi/gobdgz```

3. Use the this example directory to get more infromation about implementation
    - [Accounts examples](example/accounts)

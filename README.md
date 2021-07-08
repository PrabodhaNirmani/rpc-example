**GoLand Example for Remote Procedure Calls**

A simple distributed application for vegetable store system using Go language and Remote Procedure Calls (RPC).

Repository has both server and client code to communicate using RPC. 

Server will maintain a file which keeps records of different available vegetables including price per kg and available amount of kg for each vegetable

The server maintains logic to do below operations
1. Query the file and output names of all available vegetables.
2. Output the price per kg of a given vegetable.
3. Output the available amount of kg for a given vegetable.
4. Add new vegetable to the file with price per kg and among of kg.
5. Update the price or available amount of a given vegetable.

Clients can use server functions to do the following tasks.
1. Receive a list of all available vegetables and display.
2. Get the price per kg of a given vegetable and display.
3. Get the available amount of kg of a given vegetable and display.
4. Send a new vegetable name to the server to be added to the server file.
5. Send new price or available amount for a given vegetable to be updated in the server file.


**How to run this project**
1. Clone the repository to your local machine
2. Open up a terminal and change to the directory "vegetable-store"
2. Execute the command `go run server.go` (Now server is up)
3. Open up a new terminal and change the directory to "client".
4. Execute `go run client.go` (Now client side program is running. It takes inputs from the command line to execute functions in the server side)

(data.txt file contains the vegetable data in the in store)

**Tested version** | Go - 1.16

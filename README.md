Tic-Tac-Toe API
============== 
This is an API for the traditional `3x3` Tic-Tac-Toe where two users (`X` and `O`) try to capture `3` consecutive cells of the board to win. As a rule, `X` always starts the game.

Any of the following values are considered a valid **Cell Value**:

- **0**: Denotes an empty cell
- **1**: Denotes X
- **2**: Denotes O

There are two endpoints:
- `/ping`: A `GET` request sent to this endpoint will let the user know if the API is up and running
- `/play`: A `POST` request made to this endpoint along with a valid body causes a complete response to be sent to the enduser.

A valid body for the play endpoint must be in JSON format and include the following fields:
- `row`: The row number of the cell that we want to set
- `col`: The column number of the cell that we want to set
- `content`: A 3 x 3 array of Cell Values

The Response contains the following fields:
- `status`: This field gives more information about the requested action through the following subfields:
    - `turn`: Determines whose turn is **next**; a value of 0 (empty) in this field indicates an error 
    - `state`: More information about the response is encoded in this field which contains the following sub-fields:
        - `code`: A numeric code
        - `message`: A textual description for the code
 - `board`: A 3 x 3 array indicating the resulting board after making the requested move. If any error happens, this is usually the same board as the input. 
# Dockerization
In order to isolate this project from everything else on the host machine, I've included a docker directory in the source code, where you can use the provided Makefile to have your local environment up and running in seconds. The only requirement is to have Docker installed on the host machine.

## Makefile targets
Here's what the Makefile offer right off-the-bat:
``` bash
Usage:
  make <target>

Targets:
  up                   build, compose up, and daemonize
  down                 compose down
  down-v               compose down and remove all volumes
  bash                 execute bash inside the container
  attach               attach to the container
  ip                   show IP address of the container
  dep                  install dependencies
  all-tests            run all tests from container
  all-tests-v          run all tests from container in verbose mode
  all-tests-c          run all tests from container in coverage mode
  restart              restart the whole stack
  help                 Show help
```
So, to have a local running environment, just `cd docker` and run `make up`. This will bring a containerized server listening to port `8080` as indicated in the docker-compose file.
# Sample API Consumer
``` go
package main  
  
import (  
   "fmt"  
   "io/ioutil" "net/http" "strings")  
  
func main() {  
   url := "http://localhost:8080/play"  
   payload := strings.NewReader("{\"row\":1,\"col\":2,\"content\":[\n[0,1,0],\n[0,0,0],\n[0,0,0]\n\n]}")  
  
   req, _ := http.NewRequest("POST", url, payload)  
   res, _ := http.DefaultClient.Do(req)  
   defer res.Body.Close()  
   body, _ := ioutil.ReadAll(res.Body)  
   fmt.Println(string(body))  
}
```

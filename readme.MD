# Auction Bid Tracker

You have been asked with building part of a simple online auction system which will allow users to concurrently bid on items for sale. The system needs to be built in Go and/or Python.

Please, provide a bid-tracker interface and concrete implementation with the following functionality:

- record a user’s bid on an item;
- get the current winning bid for an item;
- get all the bids for an item;
- get all the items on which a user has bid;
- build simple REST API to manage bids.

You are not required to implement a GUI (or CLI) or persistent store (events are for reporting only). You may use any appropriate libraries to help.

Test the performance of your solution.

## About solution
This project is developed using goLang and it is containerzid application, It create all the basic setup inside container like Redis, Http server and dependences.

## Installation

Please go through the given below command to run auction bid tracker application.

```bash
- Clone or download given project.
- Make sure you have docker in you computer.
- Go to project base directory.
- Then type docker command and make build.
- docker build -t berlin . 
- Once build complete docker build, Then RUN docker
- docker run -p 8080:8000 -it berlin
- It start serving your request on 8080 port.
- Healthcheck -> http://localhost:8080/healthcheck

```

## Unit testing 

### Run inside docker image
```
- If you want to run unit test inside docker, Then open docker image.
- docker build -t berlin .
- docker run -p 8080:8000 -it berlin
- For getting container_id, run docker command (docker ps)
- docker exec -it CONTAINER_ID bash
- Then run test case : go test ./... -v  
```
### Run locally
```
- Install redis on you computer
- brew install redis
- redis-server
- Run redis in background (redis-server --daemonize yes)
- redis-cli ping
- Then run you application or unit test
- go test ./... -v  || go run main.go 

## Usage

```
- Step 1 ::
   - Add bid 
      	[ POST ] http://localhost:8080/api/berlin/internal/add-biding?username=daniel&amount=10&item=bitcoin
	    params {
		  username: daniel,
		  item: bitcoin,
		  amount: 10
	    }
- Step 2 ::
   - Get Bid by User
         [ POST ] http://localhost:8000/api/berlin/internal/all-bids-by-user?username=daniel
	     params {
		   username: daniel
	     }
- Step 3 ::
    - Get Bid by Item
          [ POST ] http://localhost:8080/api/berlin/internal/all-item-by-bid?item=bitcoin
	      params {
		    item: bitcoin
	      }
- Step 4 ::
     - Get Winner
	      [ POST ] http://localhost:8080/api/berlin/internal/winner-by-item?item=bitcoin
	      params {
		    item: bitcoin
	      }
```

## Author
[CHANDAN KUMAR]
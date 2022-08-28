# Faraway Server

What could be done to improve this code even more?
1. Adding contexts to the server for goroutines synchronization 
2. Probably some config file to store the port and maybe the difficulty of the PoW algorithm
3. Integration tests
4. Adjusting the difficulty of the PoW algorithm in real-time according to the load of our server

I could've done all these things but I just don't think that these things are so essential to demonstrate my skills.

# Commands

## Docker-compose

Run the application
```
make up
```

Down the application
```
make down
```

---
## Tests

Run the unit tests
```
make test
```

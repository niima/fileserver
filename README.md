# fileserver
A lightweight file server to host your files over local network
This repository contains a simple Go server that serves files from a specified directory. The server tries to bind to a port in the range 8080 to 8100 and starts serving files once it finds an available port.


## Features

- Serves files from a specified directory.
- Automatically creates the directory if it does not exist.
- Tries to bind to a port in the range 8080 to 8100.
- Logs the IP addresses and port where the server/your computer is running.

## Usage

### Downloading the binary
Download the binary for your operating system from the [releases](https://github.com/niima/fileserver/releases/tag/v1.0.0) page.

To run the server, use the following command:

### with Go not installed on Windows
```bash
fileserver.exe [directory]
```

### with Go not installed on macOS or Linux
```bash
./fileserver [directory]
```

### with Go installed running from source
```bash
go run main.go [directory]
```

You can also compile this code to binary file with this command to have the binary file for later use without compiling and go installed:
```bash
go build . -o fileserver.exe
```
For that, you need to install golang on your system. from [here](https://golang.org/doc/install)




Replace [directory] with the path to the directory you want to serve files from. If you don't specify a directory, the server will serve files from the default directory.
Default value is {current_directory}/files
## Requirements
- Go 1.16 or later

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
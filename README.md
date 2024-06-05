
# File Download Server

A simple REST based file download server.


## Requirements
Go >= 1.20

OR

Or use Dockerfile to build the image

## Build

Using `go build` command

```bash
  go build -o <build_directory> ./cmd/file-download-server
```

Using `Docker`
```bash
  docker build -t file-download-server .
```

**_NOTE:_** Run these commands from the root of the project

## Run
Invoke directly
```bash
<build_directory>/file-download-server -host <host_address> -port <port_number> -dir <full_path_files_directory> 
```
*Example:* ./file-download-server -host "0.0.0.0" -port "80808" -dir /app/files

**_NOTE_**: Files are picked from `<full_path_files_directory>`. So make sure files exist.

Docker
```bash
docker run -e HOST=<host_address> -e PORT=<port_number> -e CREATE_FILES="<comma_seperated_list_of_file_size:file_names>" -p 8080:<port_number> --name file-download-server file-download-server
```
*Example:* docker run -e HOST=0.0.0.0 -e PORT=8080 -e CREATE_FILES="100:100MB.bin,1024:1GB.bin,20480:10GB" -p 8080:8080 --name file-download-server file-download-server

server will be reachable via `localhost:8080` on host mahcine.

**_NOTE_**: For docker, use host as 0.0.0.0

`CREATE_FILES` format - comma seperated string containing`file_size:file_name`.

`file_size` creates a file of size 1MB*`file_size`

`file_name` API endpoints should match file names (case sensitive) check API reference section.
## API Reference

#### Get a file

```http
  GET /download-<file_name>
```

`<file_name>` should be exact match of the file name (case sensitive)

*Example:* http://127.0.0.1:8080/download-1GB.bin
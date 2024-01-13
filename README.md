# Command Execution Server

This Go program sets up a simple HTTP server that allows remote execution of commands on the host machine. It provides a basic interface to run shell commands via HTTP requests.

## Usage

1.  Build the executable:

    bashCopy code

    `go build`

2.  Run the executable:

    bashCopy code

    `./<executable_name>`

3.  The server will start listening on port `8080`.

## API Endpoint

### Execute Command

- **URL**: `/machine`
- **Method**: `POST`
- **Parameters**:
  - `cmd`: The shell command to be executed.

#### Example:

`curl -X POST http://localhost:8080/machine -d "cmd=ls -l"`

This will execute the `ls -l` command on the host machine and return the command's output as the response.

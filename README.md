# GoCliDemo

## Setup
- Install [Go](https://go.dev/).
- Install the Go extension for VSCode.

## Demo
- Create an executable with `go build`.
- Open PowerShell and run `./GoCliDemo.exe -h` to see the help message.
- Run `.\GoCliDemo.exe` from PowerShell to see the greeting message.
- Run `./GoCliDemo.exe` from bash to see the greeting message.
- Other PowerShell commands
  - `.\GoCliDemo.exe -name Chris -command random`
  - `.\GoCliDemo.exe -name Steve -command list`
  - `.\GoCliDemo.exe -command=random`
  - `.\GoCliDemo.exe -command=random -name=Chris time`
  - `.\GoCliDemo.exe time`
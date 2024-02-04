# MyApp Scenario API Server

## Repo Links

|                                 Repo                                  | Desc                       |
| :-------------------------------------------------------------------: | -------------------------- |
|        [myapp_admin](https://github.com/BlaxBerry/myapp_admin)        | Admin Dashboard Pages      |
|     [myapp_protobuf](https://github.com/BlaxBerry/myapp_protobuf)     | Protocol Buffers Generator |
| [myapp_scenario_api](https://github.com/BlaxBerry/myapp_scenario_api) | Backend for Frontend       |
| [myapp_scenario_api](https://github.com/BlaxBerry/myapp_scenario_api) | Scenario App's API Server  |
|    [myapp_notes_api](https://github.com/BlaxBerry/myapp_notes_api)    | Notes App's API Server     |

## Commands

```shell
make dev        # start dev server
make build      # run build
make preview    # run build then preview
make format     # run gofmt
make test       # run go test
```

## Tech Stacks

- [go]() v1.20
- [gin]() v1.9.1
- [makefile]() v3

## Project Structure

```txt
/
├── cmd/
│   ├── myapp_scenario_api/
│   │  └── main.go
│   └── app
├── internal/
│   ├── app/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── router/
│   │   ├── routes/
│   │   └── ...
│   └── pkg/
│       ├── constants/
│       ├── tools/
│       └── ...
├── go.mod
├── go.sum
├── Makefile
└── ...
```

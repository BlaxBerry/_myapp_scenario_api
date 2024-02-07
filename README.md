# MyApp Scenario API Server

## Repo Links

|                                 Repo                                  | Desc                            |                 Domain                  | Port |
| :-------------------------------------------------------------------: | ------------------------------- | :-------------------------------------: | :--: |
|        [myapp_admin](https://github.com/BlaxBerry/myapp_admin)        | Admin Dashboard Pages           |     https://myapp-bff.onrender.com      | 3000 |
|     [myapp_protobuf](https://github.com/BlaxBerry/myapp_protobuf)     | Protocol Buffers Code Generator |                    -                    |  -   |
|          [myapp_bff](https://github.com/BlaxBerry/myapp_bff)          | Backend for Frontend            |     https://myapp-bff.onrender.com      | 8080 |
| [myapp_scenario_api](https://github.com/BlaxBerry/myapp_scenario_api) | Scenario App's API Server       | https://myapp-scenario-api.onrender.com | 8000 |
|    [myapp_notes_api](https://github.com/BlaxBerry/myapp_notes_api)    | Notes App's API Server          |  https://myapp_notes_api.onrender.com   | 8400 |

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

.PHONY: generate migration

generate-pre:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

generate-service:
	oapi-codegen --config pkg/api/config/server.cfg.yaml pkg/api/spec/api.yaml

generate-types:
	oapi-codegen --config pkg/api/config/types.cfg.yaml pkg/api/spec/api.yaml

generate: generate-pre generate-service generate-types

abi-types:
	abigen --abi pkg/blockchain/abi/IQuoterV2.json --pkg contracts --type IQuoterV2 --out pkg/blockchain/contracts/IQuoterV2.go
	abigen --abi pkg/blockchain/abi/ISwapRouter02.json --pkg contracts --type ISwapRouter02 --out pkg/blockchain/contracts/ISwapRouter02.go
	abigen --abi pkg/blockchain/abi/IERC20.json --pkg contracts --type IERC20 --out pkg/blockchain/contracts/IERC20.go
	abigen --abi pkg/blockchain/abi/ProgrammableTokenTransfer.json --pkg contracts --type ProgrammableTokenTransfer --out pkg/blockchain/contracts/ProgrammableTokenTransfer.go

migration:
	atlas migrate diff --env gorm $(name)

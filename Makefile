build:
	@go build -o target/pokedexcli cmd/pokedexcli/main.go

run: build
	@./target/pokedexcli


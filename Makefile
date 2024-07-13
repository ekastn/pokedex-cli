build:
	@go build -o target/pokedexcli cmd/pokedexcli/*

run: build
	@./target/pokedexcli


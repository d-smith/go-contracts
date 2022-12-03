# go-contracts - sample call to smart contract via go

This uses the Token.sol contract in the hardhat-getting-started project. We use solc to generate the abi, and abigen to generate a golang binding.

## Solc set up

```
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

## Abigen

Abigen is available via the [geth install](https://geth.ethereum.org/docs/install-and-build/installing-geth)... maybe solc too?

## Code set up

Compile the smart contract

solc --abi ../../../hardhat-getting-started/contracts/Token.sol -o build

Generate the go bindings via abigen

solc --abi ../hardhat-getting-started/contracts/Token.sol -o build
cd token
abigen --abi=../build/Token.abi --pkg=token --out=Token.go

## Running the Sample

cd caller
go run caller.go

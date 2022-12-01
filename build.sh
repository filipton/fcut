#!/bin/bash
currentDir=$(pwd)

echo "Starting RUST build..."
cd ./Rust
bash ./build.sh
cd "$currentDir"
echo "RUST build done! (Docker image: filipton/fcut-rust:latest)"


echo "Starting DOTNET build..."
cd ./Dotnet/fcut-dotnet
bash ./build.sh
cd "$currentDir"
echo "DOTNET build done! (Docker image: filipton/fcut-dotnet:latest)"


echo "Starting NODE build..."
cd ./Node
bash ./build.sh
cd "$currentDir"
echo "NODE build done! (Docker image: filipton/fcut-node:latest)"


echo "Starting GO build..."
cd ./Golang
bash ./build.sh
cd "$currentDir"
echo "GO build done! (Docker image: filipton/fcut-go:latest)"


echo "Deploying images to docker registry..."
docker push filipton/fcut-rust:latest
docker push filipton/fcut-dotnet:latest
docker push filipton/fcut-node:latest
docker push filipton/fcut-go:latest

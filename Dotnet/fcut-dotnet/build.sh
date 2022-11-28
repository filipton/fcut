#!/bin/bash

dotnet publish -c Release
docker build -t filipton/fcut-dotnet:latest .

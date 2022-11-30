#!/bin/bash

cargo build --release --target x86_64-unknown-linux-musl
docker build --no-cache -t filipton/fcut-rust:latest .

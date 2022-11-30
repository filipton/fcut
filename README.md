
# fcut
Basic url shortener written in 3 languages (rust,php,C#). 
It's using redis as datastore.

Rust version should be rewritten to use connection pooling 
(so tests might not be conclusive!)
## Environment Variables
### PHP
`REDIS_HOST` (localhost)

`REDIS_PORT` (6379)

`REDIS_PASSWORD` (password)

### RUST
`REDIS_STRING` (redis://user:pass@ip:port/db)

### C#
`REDIS_ENDPOINT` (localhost:6379)

`REDIS_PASSWORD` (password)

## Performance testing
I used loader.io for performance testing, 
so the tests should be close to real-world use.

And yes, as you can se I've got a lot of invalid redirects, but it's because of redirecting to localhost. 
Test URL: https://fcut.filipton.space/cmujl06s

### PHP
![PHP](https://github.com/filipton/fcut/blob/main/tests/php.png?raw=true)

### RUST
![RUST](https://github.com/filipton/fcut/blob/main/tests/rust.png?raw=true)

### C#
![DOTNET](https://github.com/filipton/fcut/blob/main/tests/dotnet.png?raw=true)
## Contributing

Contributions are always welcome!

If you see any way to improve my code, just do it.

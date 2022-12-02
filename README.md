
# fcut
Basic url shortener written in 5 languages (rust,php,C#,node,go). 
It's using redis as datastore.

Rust version should be rewritten to use connection pooling 
(so tests might not be conclusive!)

## Currently the winner is.... GO
When i saw huge performance difference between go and dotnet i started searching for cause.
After some testing i found out that docker reduce total throughput of app. 
(In next commit i'm just going to test performance without docker)

C# (WITH DOCKER)
![DOTNET](https://github.com/filipton/fcut/blob/main/tests/dotnet.png?raw=true)

C# (WITHOUT DOCKER)
![DOTNET](https://github.com/filipton/fcut/blob/main/tests/dotnet-nodocker.png?raw=true)


## Environment Variables
### PHP
`REDIS_HOST` (localhost)

`REDIS_PORT` (6379)

`REDIS_PASSWORD` (password)

### RUST (ROCKET.RS)
`REDIS_STRING` (redis://user:pass@ip:port/db)

### C#
`REDIS_ENDPOINT` (localhost:6379)

`REDIS_PASSWORD` (password)

### NODE (FASTIFY)
`REDIS_STRING` (redis://user:pass@ip:port/db)

### GO (ECHO)
`REDIS_ENDPOINT` (localhost:6379)

`REDIS_PASSWORD` (password)

## Performance testing

I used loader.io for performance testing, 
so the tests should be close to real-world use.
(RUST AND C# ARE TESTED WHILE RUNNING IN DOCKER)

And yes, as you can se I've got a lot of invalid redirects, but it's because of redirecting to localhost. 
Test URL: https://fcut.filipton.space/cmujl06s

### PHP
![PHP](https://github.com/filipton/fcut/blob/main/tests/php.png?raw=true)

### RUST (ROCKET.RS)
![RUST](https://github.com/filipton/fcut/blob/main/tests/rust.png?raw=true)

### C#
![DOTNET](https://github.com/filipton/fcut/blob/main/tests/dotnet.png?raw=true)

### NODE (FASTIFY)
![NODE](https://github.com/filipton/fcut/blob/main/tests/nodefastify.png?raw=true)

### GO (ECHO)
![GO](https://github.com/filipton/fcut/blob/main/tests/go.png?raw=true)

## Contributing

Contributions are always welcome!

If you see any way to improve my code, just do it.

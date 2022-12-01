const redis = require('redis');
const fastify = require('fastify')({ logger: false })

fastify.register(require('@fastify/formbody'))

const redisClient = redis.createClient({
    url: process.env.REDIS_STRING
});;
redisClient.connect();

// Declare a route
fastify.get('/', async (request, reply) => {
    reply.header("Content-Type", "text/html");
    return "<!DOCTYPE html><html lang=\"en\"><head> <meta charset=\"UTF-8\"> <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"> <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"> <title>FCUT<\/title><\/head><body><div style=\"text-align: center;\"> <form method=\"post\" action=\"\/shorten\"> <label for=\"url\">URL:<\/label> <br\/> <input id=\"url\" name=\"url\" type=\"url\" required\/> <br\/> <br\/> <br\/> <input type=\"submit\" value=\"SHORTEN\"\/> <\/form><\/div><\/body><\/html>";
})

fastify.get('/:short', async (request, reply) => {
    const { short } = request.params;
    const url = await redisClient.get(short);

    if (url !== null) {
        reply.redirect(url);
    }

    return "404: Not found";
})

fastify.post('/shorten', async (request, reply) => {
    var url = await request.body.url;
    const id = makeid(8);

    if (url != null) {
        await redisClient.set(id, url);

        // re-define url to generated url 
        url = request.protocol + "://" + request.headers['host'] + "/" + id; 

        reply.header("Content-Type", "text/html");
        return "<head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>FCUT</title></head><div style=\"text-align: center\"><h1>GENERATED!</h1> <a href=\"" + url + "\">" + url + "</a></div>";
    }
    else {
        return "500";
    }
})

// Run the server!
const start = async () => {
    try {
        await fastify.listen(8002, '0.0.0.0')
    } catch (err) {
        fastify.log.error(err)
        process.exit(1)
    }
}
start()


// https://stackoverflow.com/questions/1349404/generate-random-string-characters-in-javascript
function makeid(length) {
    var result = '';
    var characters = '0123456789abcdefghijklmnopqrstuvwxyz';
    var charactersLength = characters.length;
    for (var i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}
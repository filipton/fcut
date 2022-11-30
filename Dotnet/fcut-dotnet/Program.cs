using System.Text;
using Microsoft.AspNetCore.Mvc;
using StackExchange.Redis;

ConnectionMultiplexer multiplexer = await ConnectionMultiplexer.ConnectAsync(
    new ConfigurationOptions()
    {
        EndPoints = {Environment.GetEnvironmentVariable("REDIS_ENDPOINT") ?? "localhost:6379"},
        Password = Environment.GetEnvironmentVariable("REDIS_PASSWORD") 
    }
);

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();
    
app.UseDefaultFiles();
app.UseStaticFiles();

app.MapPost("/shorten", async (HttpResponse response, HttpRequest req) =>
{
    string generated = RandomString(8);
    string dataUrl = req.Form["url"].ToString();
    
    var db = multiplexer.GetDatabase();
    var newKey = await db.StringSetAsync(generated, dataUrl);

    if (newKey)
    {
        string url = (Environment.GetEnvironmentVariable("URL") ?? "http://localhost:5002") + $"/{generated}";
        response.ContentType = "text/html";
        await response.Body.WriteAsync(
            Encoding.UTF8.GetBytes($"<head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>FCUT</title></head><div style=\"text-align: center\"><h1>GENERATED!</h1> <a href=\"{url}\">{url}</a></div>")
            );
        return;
    }

    response.StatusCode = 500;
    await response.Body.WriteAsync(
        Encoding.UTF8.GetBytes($"500: Internal server error!")
    );
});

app.MapGet("/{id}", async (HttpResponse response, string id) =>
{
    var db = multiplexer.GetDatabase();
    var key = await db.StringGetAsync(id);

    if (!key.HasValue)
    {
        await response.Body.WriteAsync(Encoding.UTF8.GetBytes("404: Not found"));
        return;
    }
    response.Redirect(key.ToString(), true);
});

app.Run();


string RandomString(int length)
{
    const string chars = "1234567890abcdefghijklmnopqrstuvwxyz";
    return new string(Enumerable.Repeat(chars, length)
        .Select(s => s[Random.Shared.Next(s.Length)]).ToArray());
}

struct ShortUrl
{
    public string Url { get; set; }
}
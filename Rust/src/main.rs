use rand::Rng;
use rocket::response::content::{self, RawHtml};
use rocket::{form::Form, fs::FileServer, response::Redirect};

#[macro_use]
extern crate rocket;

const CHARSET: &[u8] = b"1234567890\
                        abcdefghijklmnopqrstuvwxyz";
const PASSWORD_LEN: usize = 8;
const REDIS_STRING: &str = "redis://:ReallyFcut123+@34.116.201.191/";

#[derive(FromForm)]
struct ShortUrl {
    url: String,
}

#[get("/<id>")]
async fn shorten_redirect(id: &str) -> Redirect {
    let mut con = redis::Client::open(REDIS_STRING)
        .expect("Invalid connection URL")
        .get_connection()
        .expect("Failed to connect");

    let url: String = redis::cmd("GET")
        .arg(id)
        .query(&mut con)
        .expect("failed to execute get key");

    return Redirect::to(url);
}

#[post("/shorten", data = "<form>")]
async fn shorten_post(form: Form<ShortUrl>) -> content::RawHtml<String> {
    let generated: String = (0..PASSWORD_LEN)
        .map(|_| {
            let idx = rand::thread_rng().gen_range(0..CHARSET.len());
            CHARSET[idx] as char
        })
        .collect();

    let mut con = redis::Client::open(REDIS_STRING)
        .expect("Invalid connection URL")
        .get_connection()
        .expect("Failed to connect");

    let _: () = redis::cmd("SET")
        .arg(generated.clone())
        .arg(form.url.to_string())
        .query(&mut con)
        .expect("Failed to set new key");

    let url: String = std::env::var("URL").unwrap_or("http://localhost:8000".to_string()) + "/" + &generated;
    return RawHtml(format!("<head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>FCUT</title></head><div style=\"text-align: center\"><h1>GENERATED!</h1> <a href=\"{}\">{}</a></div>", url, url));
}

#[launch]
async fn rocket() -> _ {
    let static_path: String = std::env::var("STATIC").unwrap_or("static".to_string());

    rocket::build()
        .mount("/", FileServer::from(static_path))
        .mount("/", routes![shorten_post, shorten_redirect])
}

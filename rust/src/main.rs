#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;

#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

#[get("/puzzle/<word>")]
fn puzzle(word: String) -> String {
    format!("Hello, {}!", word)
}

fn main() {
    rocket::ignite().mount("/", routes![index, puzzle]).launch();
}

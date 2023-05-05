use diesel::prelude::*;
use dotenv::dotenv;
use std::env;

pub fn ConnectToDB()->PgConnection{
    dotenv.ok();
    let Database_url = env::var("DATABASE_URL")
    .expect("DATABASE_URL must be set");

PgConnection::establish(&Database_url)
.expect(&format!("Error connecting to {}", Database_url))


}
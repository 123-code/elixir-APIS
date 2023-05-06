
#[macro_use]
extern crate diesel;



use actix_web ::{HttpServer, App, web::Data, middleware::Logger};
mod routes; 
mod models;
use routes::*; 

#[actix_web::main]

async fn main() -> std::io::Result<()> {
    std::env::set_var("RUST_LOG","debug");
    std::env::set_var("RUST_BACKTRACE","1");
    env_logger::init();

    // HTTP server 
    HttpServer::new(||{
        let logger = Logger::default();
        App::new()
        .wrap(logger)
        .service(getItemByID)


    })
    .bind(("127.0.0.1",8080))?
    .run()
    .await

}

use actix_web::{
    get,
    post,
    put,
    error::ResponseError,
    web::Path,
    web::Data,
    web::Json,
    HttpResponse,
    http::{header::ContentType, StatusCode},
};

use serde::{Serialize, Deserialize};
use derive_more::{Display};

#[derive(Deserialize,Serialize)]

pub struct ItemIdentifier{
    id: String,
}

#[get("/item/{id}")]

pub async fn getItemByID(item_identifier:Path<ItemIdentifier>)-> Json <String> {
    return Json("Hello World".to_string())
    // return Json(item_identifier.into_inner().id)
}
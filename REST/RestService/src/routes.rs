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

#[post("/item")]
pub async fn createItem(item:Json<models::NewItem>)->Json<models::Item>{
    return Json(models::Item{
        id: 1,
        name: item.name.clone(),
        description: item.description.clone(),
        date_created: chrono::Local::now().naive_local(),
        value: item.value,
    })
}
use crate::args::{

}

use crate::models::{
    Item,
    NewItem,
}

use crate::db::{
    ConnectToDB,
}

use diesel::prelude::*;



#[post("/items", data = "<new_item>")]
fn create_item(connection: db::PgConnection, new_item: Json<NewItem<'_>>) -> Result<JsonValue, String> {
    let insert_result = diesel::insert_into(db::Item::table)
        .values(&*new_item)
        .get_result::<db::Item>(&connection);

    match insert_result {
        Ok(item) => Ok(json!({"status": "ok", "item": item})),
        Err(error) => Err(format!("Failed to create item: {:?}", error)),
    }
}
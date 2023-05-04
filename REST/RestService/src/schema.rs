// @generated automatically by Diesel CLI.

diesel::table! {
    item (id) {
        id -> Int4,
        name -> Varchar,
        description -> Varchar,
        date_created -> Timestamp,
        value -> Int4,
    }
}

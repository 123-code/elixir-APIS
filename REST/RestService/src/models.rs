use schema::items;

#[derive(Insertable)]
#[table_name="Item"]

// lifetime annotation
pub struct NewItem<'a>{
    pub name: &'a str,
    pub description: &'a str,
    pub value: i32,
}


#[derive(Debug,Queryable,AsChangeset)]

pub struct Item(
    pub id::i32,
    pub name::String,
    pub description::String,
    pub date_created::chrono::NaiveDateTime,
    pub value::i32,
);
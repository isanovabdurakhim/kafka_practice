create table if not exists car (
    "id" uuid PRIMARY KEY,
    "category_id" uuid,
    "investor_id" uuid,
    "state_number" varchar,
    "updated_at" timestamp default current_timestamp not null,
    "created_at" timestamp default current_timestamp not null
);

create table if not exists brand  (
    "id" uuid primary key,
    "name" varchar not null,
    "created_at" timestamp default current_timestamp not null,
    "updated_at" timestamp default current_timestamp not null
);

create table if not exists model  (
    "id" uuid primary key,
    "name" varchar not null,
    "created_at" timestamp default current_timestamp not null,
    "updated_at" timestamp default current_timestamp not null
);

Chevrolet -> Gentra -> 10 A 110 AA

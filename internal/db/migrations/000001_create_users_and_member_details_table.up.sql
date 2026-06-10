create table users (
    id serial primary key,
    username varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create table member_details (
    id serial primary key,
    user_id integer not null references users(id) on delete cascade,
    first_name varchar(255),
    last_name varchar(255),
    saldo bigint default 0,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);
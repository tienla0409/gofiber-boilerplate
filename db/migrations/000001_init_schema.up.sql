create table users (
   username   varchar(255) not null primary key,
   password   varchar(255) not null,
   email      varchar(255) not null,
   method     varchar(255) not null,
   created_at timestamptz not null default now(),
   updated_at timestamptz not null default now()
);
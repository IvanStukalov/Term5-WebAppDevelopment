drop table if exists "events" CASCADE;

drop table if exists "stars" CASCADE;

drop table if exists "star_events" CASCADE;

drop table if exists "users" CASCADE;

create table stars (
    star_id serial not null constraint star_pk primary key,
    name varchar(30) not null UNIQUE,
    description varchar(200),
    distance real,
    age real,
    magnitude real,
    image varchar(100),
    is_active boolean
);

alter table
    stars owner to postgres;

create table "users" (
    user_id serial not null constraint user_pk primary key,
    login varchar(50),
    password varchar(200),
    is_moderator boolean,
    registration_date timestamp
);

alter table
    "users" owner to postgres;

create table events (
    event_id serial not null constraint event_pk primary key,
    name varchar(50),
    status varchar(20),
    creation_date timestamp,
    formation_date timestamp,
    completion_date timestamp,
    moderator_id integer,
    creator_id integer constraint creator_id_fk references "users"(user_id)
);

alter table
    events owner to postgres;

create table star_events (
    star_id integer constraint star_event_star_star_id_fk references stars (star_id) on delete cascade,
    event_id integer constraint star_event_event_event_id_fk references events (event_id) on delete cascade,
    primary key (star_id, event_id)
);

alter table
    star_events owner to postgres;
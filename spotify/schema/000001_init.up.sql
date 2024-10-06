CREATE TABLE music_library (
    id bigserial primary key not null,
    group_name varchar(64) not null,
    song_name varchar(64) not null,
    release_date varchar(32) not null,
    text text not null,
    link text not null
)
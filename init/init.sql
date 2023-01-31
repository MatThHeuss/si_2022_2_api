create table IF NOT EXISTS users (
    id varchar(255) primary key not null,
    name varchar(80) not null,
    profile_image_url varchar(255) not null,
    email varchar(255) unique  not null,
    password varchar(255) not null,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


create table IF NOT EXISTS announcement_images (
    announcement_id varchar(255)  not null,
    image_url varchar(255) primary key not null
);


create table if not EXISTS announcement (
    id varchar(255) primary key not null,
    name varchar(255) not null,
    category varchar(255) not null,
    description varchar(255) not null,
    address varchar(255) not null,
    postal_code varchar(255) not null,
    user_id varchar(255) not null
);


create table if not Exists chat (
    id varchar(255) primary key not null,
    sender varchar(255) not null,
    receiver varchar(255) not null,
    content varchar(255) not null,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE announcement_images ADD FOREIGN KEY (announcement_id) REFERENCES announcement (id);
ALTER TABLE announcement ADD FOREIGN KEY (user_id) REFERENCES users(id);
CREaTE table sellers (
id bigserial not null primary key,
first_name varchar(20) not null,
second_name varchar(20) not null,
phone_number varchar(13) not null unique,
encrypted_password varchar not null
);


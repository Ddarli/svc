-----------------------*CREATE USERS*---------------------------
create table users
(
    id         uuid,
    pkey       text unique        not null,
    username   varchar(50) unique not null,
    password   text               not null,
    email      varchar(50) unique not null,
    phone      varchar(50) unique not null,
    role       int,
    created_at timestamp default now()
);

-----------------------*ADD TEST USERS*---------------------------
insert into users (id, pkey, username, password, email, phone, role)
values ('550e8400-e29b-41d4-a716-446655440000
', 'fa1c134b11761e9e02cbad596e27fb4a97fa28f7960e658473c3f538de95394d', 'admin', 'admin', 'test@mail.com', '89776453424', 1),
       ('230e8411-e29b-41d4-a716-446655440000
', 'edd8c5c0936002eb0a8e05122ab49255c7a9fba99ef01c6d0db2cbbf5702d5f4', 'user', 'user', 'test2@mail.com', '89776453421', 2);
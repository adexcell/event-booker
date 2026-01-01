create table if not exists events (
    event_id uuid primary key,
    title text not null,
    description text,
    total_slots int not null check(total_slots >= 0),
    available_slots int not null check(available_slots >= 0),
    book_timeout bigint not null,
    event_at timestamp with time zone not null,
    created_at timestamp with time zone default NOW(),
    deleted_at timestamp with time zone
);

create table if not exists bookings (
    booking_id uuid primary key,
    event_id uuid not null references events(event_id) on delete cascade,
    user_id uuid not null references users(user_id),
    status small int not null default 0,
    expires_at timestamp with time zone not null,
    created_at timestamp with time zone default NOW(),
);

create table if not exists users (
    user_id uuid primary key,
    name text not null,
    surname text not null,
    email text not null,
    password_hash text not null,
    telegram_id bigint,
    created_at timestamp with time zone default NOW(),
    deleted_at timestamp with time zone
);

create index idx_bookings_expires_status on bookings(expires_at) where status = 0;
create index idx_users_not_deleted on users(id) where deleted_at is null;
create index idx_events_not_deleted on events(id) where deleted_at is null;


create table if not exists users (
  id bigint primary key generated always as identity,
  username text not null unique,
  email text not null unique,
  password text not null,
  created_at timestamptz default now()
);

create table if not exists posts (
  id bigint primary key generated always as identity,
  user_id bigint references users (id),
  title text not null,
  content text not null,
  created_at timestamptz default now()
);

create table if not exists categories (
  id bigint primary key generated always as identity,
  name text not null unique
);

create table if not exists post_categories (
  post_id bigint references posts (id),
  category_id bigint references categories (id),
  primary key (post_id, category_id)
);

create table if not exists comments (
  id bigint primary key generated always as identity,
  post_id bigint references posts (id),
  user_id bigint references users (id),
  content text not null,
  created_at timestamptz default now()
);

create table if not exists likes (
  id bigint primary key generated always as identity,
  user_id bigint references users (id),
  post_id bigint references posts (id),
  comment_id bigint references comments (id),
  is_like boolean not null,
  created_at timestamptz default now()
);

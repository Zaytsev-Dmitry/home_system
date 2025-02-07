create database auth;
CREATE USER auth_user WITH LOGIN PASSWORD '8c6b18ac-84ff-4436-8916-6e11aaa41e92';
GRANT ALL PRIVILEGES ON DATABASE auth TO auth_user;

create database notes;
CREATE USER notes_user WITH LOGIN PASSWORD '5fea3e86-29b6-4f85-a5ed-3117d7ff7c7f';
GRANT ALL PRIVILEGES ON DATABASE notes TO notes_user;

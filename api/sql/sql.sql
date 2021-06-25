CREATE DATABASE IF NOT EXISTS twitter;
USE twitter;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE  usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(120) not null,
    criadoEm timestamp default current_timestamp()

)ENGINE=INNODB;

DROP TABLE IF EXISTS seguidores;
CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    References usuarios(id)
    on Delete cascade,

    seguidor_id int not null, 
    FOREIGN key (seguidor_id)
    References usuarios(id)
    on delete cascade,

    primary key(usuario_id, seguidor_id)

)ENGINE=INNODB;
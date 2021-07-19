CREATE DATABASE IF NOT EXISTS twitter;
USE twitter;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
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


CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autor_id int not null, 
    FOREIGN key (autor_id)
    References usuarios(id)
    on delete cascade,
    curtidas int default 0,
    criadaEm timestamp default current_timestamp
)ENGINE=INNODB;
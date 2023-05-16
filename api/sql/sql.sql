CREATE DATABASE IF NOT EXIST devbook;
USE devbook;

DROP TABLE IF EXIST publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100)  NOT NULL,
    criandoEm TIMESTAMP default CURRENT_TIMESTAMP()
)ENGINE=InnoDB;

CREATE TABLE seguidores (
    usuario_id INT NOT NULL, FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    seguidor_id INT NOT NULL, FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE,

    PRIMARY KEY(usuario_id, seguidor_id)
)ENGINE=InnoDB;


CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo VARCHAR(50) NOT NULL,
    conteudo VARCHAR(300) NOT NULL,
    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id) ON DELETE CASCADE,

    curtidas int DEFAULT 0,
    criadaEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;
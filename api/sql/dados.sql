insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "u1", "us1@gmail.com", "$2a$10$arVNiQ.no5HAQ4KxJISlPuqxSROtbK/iXe2h2D13r1N.cQqKPitOy"),
("Usuario 2", "u2", "us2@gmail.com", "$2a$10$arVNiQ.no5HAQ4KxJISlPuqxSROtbK/iXe2h2D13r1N.cQqKPitOy"),
("Usuario 3", "u3", "us3@gmail.com", "$2a$10$arVNiQ.no5HAQ4KxJISlPuqxSROtbK/iXe2h2D13r1N.cQqKPitOy"),
("Usuario 4", "u4", "us4@gmail.com", "$2a$10$arVNiQ.no5HAQ4KxJISlPuqxSROtbK/iXe2h2D13r1N.cQqKPitOy");


insert into seguidores(usuario_id, seguidor_id)
values
(1,2),
(3, 1),
(1,3),
(4,1);

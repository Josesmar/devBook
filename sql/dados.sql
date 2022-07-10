insert into usuarios (nome, nick, email, senha)
values 
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$10$HK/b0564d/TSyKUOBDOHe.qs3k6J2Q6RvUI2I2JFJ32nmd4HTgg7C"),
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$10$HK/b0564d/TSyKUOBDOHe.qs3k6J2Q6RvUI2I2JFJ32nmd4HTgg7C"),
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$10$HK/b0564d/TSyKUOBDOHe.qs3k6J2Q6RvUI2I2JFJ32nmd4HTgg7C"),
("Usuario 4", "usuario_4", "usuario4@gmail.com", "$10$HK/b0564d/TSyKUOBDOHe.qs3k6J2Q6RvUI2I2JFJ32nmd4HTgg7C");

insert into seguidores (usuario_id, seguidor_id)
values 
(1,2),
(3,1),
(1,3);

insert into publicacoes (titulo, conteudo, autor_id)
values
("Publicação do usuário 1", "Essa é a publicação do ususuário 1! Obaa!!", 1),
("Publicação do usuário 2", "Essa é a publicação do ususuário 1! Obaa!!", 2),
("Publicação do usuário 3", "Essa é a publicação do ususuário 1! Obaa!!", 3);
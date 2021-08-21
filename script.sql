create database aluraloja;

create table produtos
(
	id int auto_increment,
	nome varchar(256) null,
	descricao text null,
	preco decimal null,
	quantidade int null,
	constraint produtos_pk
		primary key (id)
);

insert into produtos (nome, descricao, preco, quantidade)
values
('Camiseta', 'Preta', 19, 10),
('Fone', 'Muito bom', 3, 6);

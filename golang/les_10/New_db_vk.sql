
drop database if exists vk;
create database vk;
use vk;

drop table if exists users;
create table users(
	id BIGINT unsigned not null auto_increment primary key,
	firstname VARCHAR(100),
	lastname VARCHAR(100),
	email VARCHAR(100) unique,
	password_hash VARCHAR(256),
	phone BIGINT unsigned unique,
	
	index idx_users_username(firstname, lastname)
);

drop table if exists profiles;
create table profiles(
	user_id BIGINT unsigned not null primary key,
	gender CHAR(1),
	hometown VARCHAR(200),
	created_at DATETIME default NOW()
);

#1 x 1
alter table profiles add constraint fk_profiles_user_id
foreign key (user_id) references users(id)
on update cascade
on delete restrict
;

alter table profiles add column birthday DATETIME;

#alter table profiles modify column birthday date;
#alter table profiles rename column birthday to date_of_brth;
#alter table profiles drop column date_of_brth;

# 1 X M
drop table if exists messages;
create table messages(
	id SERIAL,
	from_user_id BIGINT unsigned not null,
	to_user_id BIGINT unsigned not null,
	body TEXT,
	created_at DATETIME default NOW(),
	
	foreign key (from_user_id) references users(id),
	foreign key (to_user_id) references users(id)
);

drop table if exists friends_requests;
create table friends_requests(
	initiator_user_id BIGINT unsigned not null,
	target_user_id BIGINT unsigned not null,
	#status TINYINT(1), # 1, 2, 3, 4, -98
	status ENUM('requested', 'approved', 'declined', 'unfriended'),
	created_at DATETIME default NOW(),
	updated_at DATETIME on update CURRENT_TIMESTAMP,
	
	primary key (initiator_user_id, target_user_id),
	foreign key (initiator_user_id) references users(id),
	foreign key (target_user_id) references users(id),
	check (initiator_user_id != target_user_id)
);

alter table friends_requests
add check(initiator_user_id <> target_user_id);


drop table if exists communities;
create table communities(
	id SERIAL,
	name VARCHAR(255),
	admin_user_id BIGINT UnSIGNED not null,
	user_id BIGINT unsigned not null,
	
	index (name),
	foreign key (admin_user_id) references users(id)
);

drop table if exists users_communities;
create table users_communities(
	user_id BIGINT unsigned not null,
	community_id BIGINT unsigned not null,
	
	primary key (user_id, community_id),
	foreign key (user_id) references users(id),
	foreign key (community_id) references communities(id)
);

drop table if exists media_types;
create table media_types(
	id SERIAL,
	name VARCHAR(255), # 'text', 'video'
	created_at DATETIME default NOW(),
	updated_at DATETIME on update CURRENT_TIMESTAMP
);

drop table if exists media;
create table media(
	id SERIAL,
	user_id BIGINT unsigned not null,
	#media_type ENUM('text', 'video', 'music', 'image'),
	media_type_id BIGINT unsigned not null,
	body VARCHAR(4000),
	#file blob,
	filename VARCHAR(255),
	metadata JSON,
	created_at DATETIME default NOW(),
	updated_at DATETIME on update CURRENT_TIMESTAMP,
	
	foreign key (user_id) references users(id),
	foreign key (media_type_id) references media_types(id)
);

drop table if exists likes;
create table likes(
	id SERIAL,
	user_id BIGINT unsigned not null,
	media_id BIGINT unsigned not null,
	
	foreign key (user_id) references users(id),
	foreign key (media_id) references media(id)
);

	DROP TABLE IF EXISTS `photo_albums`;
CREATE TABLE `photo_albums` (
	`id` SERIAL,
	`name` varchar(255) DEFAULT NULL,
    `user_id` BIGINT UNSIGNED DEFAULT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id),
  	PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `photos`;
CREATE TABLE `photos` (
	id SERIAL,
	`album_id` BIGINT unsigned NULL,
	`media_id` BIGINT unsigned NOT NULL,

	FOREIGN KEY (album_id) REFERENCES photo_albums(id),
    FOREIGN KEY (media_id) REFERENCES media(id)
);	
	# M x M
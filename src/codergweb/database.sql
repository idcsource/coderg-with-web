-- CoderG的项目网站所用数据库SQL文件，此文件针对PostgreSQL


-- 有问不一定必答数据表
drop table IF EXISTS ask_not_answer CASCADE;
create table ask_not_answer (
	id bigserial NOT NULL,
	node int,
	title char(255),
	email char(255),
	content text,
	datetime bigint,
	answertime bigint,
	answer text,
	CONSTRAINT ask_id PRIMARY KEY (id)
);
CREATE INDEX ask_answer_node ON ask_not_answer USING btree (node);

-- 管理员表
drop table if exists admins cascade;
create table admins (
	id serial not null,
	name char(255),
	pass char(40),
	CONSTRAINT admins_id PRIMARY KEY (id)
);
CREATE INDEX admins_name ON admins USING btree (name);

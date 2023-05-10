-- 创建用户表
create table user(
  uuid varchar2(32) primary key, -- uuid
  name varchar(32) unique, -- 用户名,未加密
  realname varchar2(32), -- 真实姓名
  username varchar2(32) unique, -- 加密后的用户名
  password varchar2(32), -- 加密后的密码
  sex int, -- 性别
  email varchar2(32), -- 邮箱地址
  area_province varchar2(6), -- 省份
  area_city varchar2(6), -- 城市 
  status int default 0, -- 状态,0-禁用,1-启用
  update_time datetime, -- 更新日期
  avatar_img varchar2(32) -- 头像图片地址
);
-- 创建角色表
create table role(
  uuid varchar2(20) primary key, -- uuid
  role_id varchar2(20) unique,  -- 角色id
  role_name varchar2(20), -- 角色名
  update_time datetime -- 更新时间
);
-- 创建用户角色关联表
create table user_with_role(
  role_id varchar2(20), -- 角色的uuid
  user_id varchar2(32), -- 用户的uuid
  update_time datetime -- 更新时间
);
-- 创建菜单表
create table menu(
  uuid varchar2(20) primary key, -- uuid
  menu_path varchar2(20), -- 菜单路径
  menu_name varchar2(20), -- 菜单名
  menu_icon varchar2(20), -- 菜单图标
  parent_menu varchar2(20), -- 上一级菜单
  level int, --等级
  update_time datetime -- 更新时间
);
-- 创建菜单角色关联表
create table menu_with_role (
  menu_id varchar2(20), -- 菜单的uuid
  role_id varchar2(20), -- 角色的uuid
  update_time datetime -- 更新时间
);
-- 创建部门表
create table department(
  uuid varchar2(20) primary key, -- 部门的uuid
  department_name varchar2(20),  -- 部门名称
  parent_department varchar2(20), -- 上一级部门的uuid
  level int,  -- 部门等级
  update_time datetime  -- 更新时间
);
-- 创建用户和部门关联表
create table user_with_department(
  user_id varchar2(20) unique, -- 用户的uuid
  department_id varchar2(20) unique, -- 部门的uuid
  update_time datetime -- 更新时间
);
-- 创建文章表
create table article(
  uuid varchar2(20) primary key, -- 文章的uuid
  title varchar2(20), -- 文章标题
  author varchar2(20), -- 作者的uuid
  is_all_read int default 1, -- 是否所有人可读,0-不是,1-是
  is_all_write int default 0, -- 是否所有人可写,0-不是,1-是
  create_time datetime, -- 创建时间
  update_time datetime  -- 更新时间
);
-- 创建文章权限表
-- 注意，作者创建的文章具有可读写全部权限
create table article_permission(
  user_id varchar2(20), -- 用户uuid
  article_id varchar2(20), -- 文章uuid
  read int, -- 读权限,1-可读
  write int, -- 写权限,1-可写
  update_time datetime -- 更新时间
);
-- 创建主机表
create table host(
  ip varchar2(20) primary key, -- ip地址
  port varchar2(10) default '22', -- 端口
  hostname varchar2(20), -- 主机名
  username varchar2(20), -- 用户名
  password varchar2(20), -- 密码
  update_time datetime -- 更新时间
);
-- 创建告警表


-- 测试数据 --
--- 插入用户测试数据
insert into user(uuid,name,username,password,email,status) values('123456','admin','21232f297a57a5a743894a0e4a801fc3','e10adc3949ba59abbe56e057f20f883e','123456@qq.com',1);
insert into user(uuid,name,username,password,email) values('123457','test','098f6bcd4621d373cade4e832627b4f6','e10adc3949ba59abbe56e057f20f883e','123457@qq.com');
-- 插入菜单测试数据
insert into menu(uuid,menu_path,menu_name,level) values('111','/user','用户',1);
insert into menu(uuid,menu_path,menu_name,level) values('112','/alarm','告警',1);
insert into menu(uuid,menu_name,level) values('113','主机',1);
insert into menu(uuid,menu_path,menu_name,parent_menu,level) values('114','/host','113','详情',2);
-- 插入角色测试数据
insert into role(uuid,role_id,role_name) values('1','admin','管理员');
insert into role(uuid,role_id,role_name) values('2','test','普通用户');
-- 插入角色与菜单测试数据
insert into menu_with_role(menu_id,role_id) values('111','1');
insert into menu_with_role(menu_id,role_id) values('112','1');
insert into menu_with_role(menu_id,role_id) values('113','1');
insert into menu_with_role(menu_id,role_id) values('114','1');
insert into menu_with_role(menu_id,role_id) values('111','2');
-- 插入用户与角色测试数据
insert into user_with_role(role_id,user_id) values('1','123456');
-- 插入部门测试数据
insert into department(uuid,department_name,level) values('11','维护室',1);
insert into department(uuid,department_name,level) values('12','办公室',1);
insert into department(uuid,department_name,parent_department,level) values('13','经理办公室室','12',2);
-- 插入用户与部门测试数据
insert into user_with_department(user_id,department_id) values('123456','11');

-- 查询样例 --
-- 用户登陆,验证用户名、密码
select a.uuid,a.name from user a where a.status = 1 and a.username = '21232f297a57a5a743894a0e4a801fc3' and a.password = 'e10adc3949ba59abbe56e057f20f883e';
-- 获取用户角色
select a.role_id,a.role_name from role a where a.uuid = (select b.role_id from user_with_role b where b.user_id = '123456');
-- 根据用户获取角色拥有的菜单权限
select a.uuid,a.menu_path,a.menu_name,a.menu_icon,a.parent_menu,a.level from menu a where a.uuid in (select b.menu_id from menu_with_role b where b.role_id in (select c.uuid from role c where c.));
-- 
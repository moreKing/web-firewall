-- noinspection SqlNoDataSourceInspectionForFile

-- CREATE DATABASE firewalld;
--
-- 用户表格
--
CREATE TABLE "public"."users" (
                                  "id" bigserial,
                                  "loginname" VARCHAR ( 255 ) NOT NULL,
                                  "username" VARCHAR ( 255 ) NOT NULL,
                                  "state" INT2 NOT NULL,
                                  "slat" TEXT,
                                  "password" TEXT,
                                  "totp_state" BOOLEAN,
                                  "totp_token" TEXT,
                                  "email" VARCHAR ( 255 ),
                                  "mobile" VARCHAR ( 20 ),
                                  "authenticate_id" INT8,
                                  "role_id" INT8 NOT NULL,
                                  "pwd_update_at" INT8,
                                  "lastlogin_at" INT8,
                                  "created_at" INT8,
                                  "updated_at" INT8,
                                  "deleted_at" INT8,
                                  PRIMARY KEY ( "id" )
);
COMMENT ON COLUMN "public"."users"."id" IS '主键';
COMMENT ON COLUMN "public"."users"."loginname" IS '登录名称';
COMMENT ON COLUMN "public"."users"."username" IS '用户名';
COMMENT ON COLUMN "public"."users"."state" IS '1.启用 2.有效期 3.禁用';
COMMENT ON COLUMN "public"."users"."slat" IS '密码加盐';
COMMENT ON COLUMN "public"."users"."password" IS 'SM3加密密码';
COMMENT ON COLUMN "public"."users"."totp_state" IS '手机令牌绑定状态，true代表需要重新绑定';
COMMENT ON COLUMN "public"."users"."totp_token" IS 'totp自动生成的token';
COMMENT ON COLUMN "public"."users"."pwd_update_at" IS '最后改密时间';
COMMENT ON COLUMN "public"."users"."lastlogin_at" IS '最后登录时间';
INSERT INTO "public"."users"
VALUES
    (
        1,
        'admin',
        '超级管理员',
        1,
        'C4Hsv3wp6oO8GVpkkIbQp',
        'a66135dceac3c1e4a3a03b021c6da227a351513522c4ea6d8ea558082dce9334',
        'f',
        NULL,
        NULL,
        NULL,
        1,
        1,
        NULL,
        NULL,
        NULL,
        NULL,
        0
    );
SELECT
    setval( '"public"."users_id_seq"', 2, FALSE );
--
-- 系统配置
--
CREATE TABLE "public"."system_conf" ( "id" smallserial, "name" VARCHAR ( 255 ) NOT NULL, "config" JSON, PRIMARY KEY ( "id" ) );
COMMENT ON COLUMN "public"."system_conf"."id" IS '主键';
COMMENT ON COLUMN "public"."system_conf"."name" IS '配置项名称';
COMMENT ON COLUMN "public"."system_conf"."config" IS '配置项内容';
INSERT INTO "public"."system_conf"
VALUES
    ( 1, 'passwordComplex', '{"length":6,"validity":1,"expire":0,"differTimes":2,"complex":1}' );
INSERT INTO "public"."system_conf"
VALUES
    ( 2, 'authenticate', '{"totpOffset":3,"totpIssuer":"Web Firewall","messageOffset":3,"emailOffset":3}' );
INSERT INTO "public"."system_conf"
VALUES
    ( 3, 'webTimeout', '{"timeout":15}' );
INSERT INTO "public"."system_conf"
VALUES
    ( 4, 'message_server', '{"state":0,"url":"","method":"POST","encType":"json","parameters":[],"content":"Web Firewall 验证码：{code}, {validity}分钟内有效，请勿告知他人。"}' );
INSERT INTO "public"."system_conf"
VALUES
    ( 5, 'email_server', '{"enable":false,"smtp":"","port":25,"email":"","account":"Web Firewall 运维团队","protocol":1,"password":""}' );
INSERT INTO "public"."system_conf"
VALUES
    ( 8, 'account_exception_rule', '{"zombie":90,"pwdTimeout":30,"pwdWeak":{"length":6,"digit":0,"upper":1,"lower":1,"special":1}}' );

--
-- 登陆日志
--
CREATE TABLE "public"."log_logins" (
                                       "uuid" CHAR ( 21 ) NOT NULL,
                                       "loginname" VARCHAR ( 255 ) NOT NULL,
                                       "username" VARCHAR ( 255 ),
                                       "client_ip" INET NOT NULL,
                                       "user_id" INT8 NOT NULL,
                                       "totp_code" VARCHAR ( 10 ),
                                       "success" BOOLEAN NOT NULL,
                                       "online" BOOLEAN NOT NULL,
                                       "department_id" INT8 NOT NULL,
                                       "log" TEXT,
                                       "login_at" INT8 NOT NULL,
                                       "logout_at" INT8 NOT NULL,
                                       PRIMARY KEY ( "uuid" )
);
COMMENT ON COLUMN "public"."log_logins"."uuid" IS '主键';
COMMENT ON COLUMN "public"."log_logins"."loginname" IS '登陆名';
COMMENT ON COLUMN "public"."log_logins"."username" IS '用户名';
COMMENT ON COLUMN "public"."log_logins"."client_ip" IS '登陆IP';
COMMENT ON COLUMN "public"."log_logins"."user_id" IS '用户id，登陆失败为0';
COMMENT ON COLUMN "public"."log_logins"."totp_code" IS '手机令牌totp，防止短时间再次使用';
COMMENT ON COLUMN "public"."log_logins"."success" IS 'TRUE 为登陆成功';
COMMENT ON COLUMN "public"."log_logins"."online" IS 'TRUE 用户在线';
COMMENT ON COLUMN "public"."log_logins"."department_id" IS '登陆用户所属部门，审计管理员只能查看本部门的，0所有部门都可以查看';
COMMENT ON COLUMN "public"."log_logins"."log" IS '登出日志，如果登陆失败则为登陆失败日志';
--
-- 配置日志
--
CREATE TABLE "public"."log_settings" (
                                         "id" bigserial,
                                         "name" VARCHAR ( 100 ) NOT NULL,
                                         "loginname" VARCHAR ( 255 ) NOT NULL,
                                         "username" VARCHAR ( 255 ),
                                         "client_ip" INET NOT NULL,
                                         "user_id" INT8 NOT NULL,
                                         "success" BOOLEAN NOT NULL,
                                         "department_id" INT8 NOT NULL,
                                         "request_method" VARCHAR ( 10 ) NOT NULL,
                                         "request_path" VARCHAR ( 255 ) NOT NULL,
                                         "request_body" TEXT NOT NULL,
                                         "response_code" INT4 NOT NULL,
                                         "response_error" TEXT,
                                         "response_body" TEXT NOT NULL,
                                         "created_at" INT8,
                                         PRIMARY KEY ( "id" )
);
COMMENT ON COLUMN "public"."log_settings"."id" IS '主键';
COMMENT ON COLUMN "public"."log_settings"."name" IS '接口名称';
COMMENT ON COLUMN "public"."log_settings"."loginname" IS '登陆名';
COMMENT ON COLUMN "public"."log_settings"."username" IS '用户名';
COMMENT ON COLUMN "public"."log_settings"."client_ip" IS '登陆IP';
COMMENT ON COLUMN "public"."log_settings"."user_id" IS '用户id';
COMMENT ON COLUMN "public"."log_settings"."success" IS '操作成功为true，当响应码与code==0 时才为成功';
COMMENT ON COLUMN "public"."log_settings"."department_id" IS '部门id';
COMMENT ON COLUMN "public"."log_settings"."request_method" IS '请求方式';
COMMENT ON COLUMN "public"."log_settings"."request_body" IS '请求内容';
COMMENT ON COLUMN "public"."log_settings"."response_code" IS '相应码';
COMMENT ON COLUMN "public"."log_settings"."response_error" IS '错误内容';
COMMENT ON COLUMN "public"."log_settings"."response_body" IS '相应内容';
--
-- shell操作记录
--
CREATE TABLE "public"."log_shell" (
                                      "id" bigserial,
                                      "loginname" VARCHAR ( 255 ) NOT NULL,
                                      "username" VARCHAR ( 255 ) NOT NULL,
                                      "client_ip" INET NOT NULL,
                                      "user_id" INT8 NOT NULL,
                                      "success" BOOLEAN NOT NULL,
                                      "online" BOOLEAN NOT NULL,
                                      "filename" VARCHAR ( 255 ) NOT NULL,
                                      "md5" TEXT,
                                      "size" INT8,
                                      "logout_at" INT8 NOT NULL,
                                      "created_at" INT8,
                                      PRIMARY KEY ( "id" )
);
COMMENT ON COLUMN "public"."log_shell"."id" IS '主键';
COMMENT ON COLUMN "public"."log_shell"."loginname" IS '登陆名';
COMMENT ON COLUMN "public"."log_shell"."username" IS '用户名';
COMMENT ON COLUMN "public"."log_shell"."client_ip" IS '登陆IP';
COMMENT ON COLUMN "public"."log_shell"."md5" IS '记录完成后防止记录被篡改，需要记录文件md5值';
COMMENT ON COLUMN "public"."log_shell"."user_id" IS '用户id';
COMMENT ON COLUMN "public"."log_shell"."success" IS '操作成功为true';
COMMENT ON COLUMN "public"."log_shell"."online" IS 'true在线';
--
-- rule记录
--
CREATE TABLE "public"."rulesets" (
                                     "id" bigserial,
--"name" VARCHAR ( 255 ) NOT NULL,
                                     "comment" VARCHAR ( 255 ) NOT NULL,
                                     "chain" INT4 NOT NULL,
                                     "position" INT4 NOT NULL,
                                     "expr" JSON NOT NULL,
                                     "created_at" INT8,
                                     "deleted_at" INT8,
                                     PRIMARY KEY ( "id" )
);
COMMENT ON TABLE "public"."rulesets" IS '防火墙规则表';
COMMENT ON COLUMN "public"."rulesets"."id" IS '主键';
--COMMENT ON COLUMN "public"."rulesets"."name" IS '规则名，无意义给用户看的';
COMMENT ON COLUMN "public"."rulesets"."comment" IS '备注，无意义给用户看的';
COMMENT ON COLUMN "public"."rulesets"."chain" IS '属于链 1 入站策略 2 出站策略 3 目的地址转换 4 源地址转换 5 入站限流 6 出站限流 7 ip黑白名单';
COMMENT ON COLUMN "public"."rulesets"."position" IS '规则位置，重启服务时按此从小到大排序';
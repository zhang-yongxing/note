// postgresql
CREATE TABLE "sort"(
    "sort_id" char(32) primary key NOT NULL ,
    "sort_name" varchar(15) NOT NULL,
    "user_id" varchar(32) NOT NULL references "user"(user_id),
    "create_time" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "update_time" timestamp NOT NULL
);
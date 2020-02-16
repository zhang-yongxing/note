// postgresql
CREATE TABLE "note"(
    "note_id" char(32) primary key NOT NULL ,
    "note_name" varchar(50) NOT NULL,
    "note_des" varchar(50),
    "note_content" text NOT NULL,
    "is_hide" bool DEFAULT true,
    "is_deleted" bool DEFAULT false ,
    "sort_id" varchar(32) NOT NULL references "sort"(sort_id),
    "create_time" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "update_time" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

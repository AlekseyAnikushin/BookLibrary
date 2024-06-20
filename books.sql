CREATE TABLE IF NOT EXISTS public.books
(
    "ID" integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    "Title" character varying(500) COLLATE pg_catalog."default" NOT NULL,
    "AuthorID" integer NOT NULL,
    "Year" smallint,
    "ISBN" character(13) COLLATE pg_catalog."default",
    CONSTRAINT books_pkey PRIMARY KEY ("ID")
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.books
    OWNER to go_user;
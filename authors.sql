CREATE TABLE IF NOT EXISTS public.authors
(
    "ID" integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    "Name" character varying(50) COLLATE pg_catalog."default" NOT NULL,
    "Surname" character varying(50) COLLATE pg_catalog."default" NOT NULL,
    "Biography" text COLLATE pg_catalog."default",
    "Birthdate" date,
    CONSTRAINT authors_pkey PRIMARY KEY ("ID")
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.authors
    OWNER to go_user;
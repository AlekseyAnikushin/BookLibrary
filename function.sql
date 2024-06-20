CREATE OR REPLACE FUNCTION public.UpdateBookAndAuthor(
	book_id integer,
	book_title character varying,
	book_author_id integer,
	book_year smallint,
	book_isbn character,
	author_id integer,
	author_name character varying,
	author_surname character varying,
	author_biography text,
	author_birthdate date)
    RETURNS integer
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE resultCode integer;
BEGIN
	IF NOT EXISTS(SELECT "ID" FROM public.Authors WHERE "ID" = $6) THEN
		RETURN 1;
	END IF;

	IF NOT EXISTS(SELECT "ID" FROM public.Books WHERE "ID" = $1) THEN
		RETURN 2;
	END IF;
	
    BEGIN
		UPDATE public.Books SET
			"Title" =    $2,
			"AuthorID" = $3,
			"Year" =     $4,
			"ISBN" =     $5
		WHERE "ID" =     $1;

		UPDATE public.Authors SET 
			"Name" =      $7,
			"Surname" =   $8,
			"Biography" = $9,
			"Birthdate" = $10
		WHERE "ID" =      $6;

		resultCode := 0;
	EXCEPTION
		WHEN OTHERS THEN
		resultCode := -1;
	END;

    RETURN resultCode;
END;
$BODY$;

ALTER FUNCTION public.UpdateBookAndAuthor(integer, character varying, integer, smallint, character, integer, character varying, character varying, text, date)
    OWNER TO go_user;
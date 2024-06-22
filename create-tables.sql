DROP TABLE IF EXISTS public.album;
CREATE TABLE IF NOT EXISTS public.album (
  id         INT GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL DEFAULT 0
);

ALTER TABLE IF EXISTS public.album
    OWNER to postgres;

GRANT ALL ON TABLE public.album TO postgres WITH GRANT OPTION;

INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
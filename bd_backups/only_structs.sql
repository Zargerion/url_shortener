--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Debian 13.3-1.pgdg100+1)
-- Dumped by pg_dump version 13.3 (Debian 13.3-1.pgdg100+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: url_pairs; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.url_pairs (
    id integer NOT NULL,
    short_url character varying(255) NOT NULL,
    original_url character varying(255) NOT NULL
);


ALTER TABLE public.url_pairs OWNER TO "user";

--
-- Name: url_pairs_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.url_pairs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.url_pairs_id_seq OWNER TO "user";

--
-- Name: url_pairs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.url_pairs_id_seq OWNED BY public.url_pairs.id;


--
-- Name: url_pairs id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.url_pairs ALTER COLUMN id SET DEFAULT nextval('public.url_pairs_id_seq'::regclass);


--
-- Name: url_pairs url_pairs_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.url_pairs
    ADD CONSTRAINT url_pairs_pkey PRIMARY KEY (id);


--
-- Name: url_pairs url_pairs_short_url_key; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.url_pairs
    ADD CONSTRAINT url_pairs_short_url_key UNIQUE (short_url);


--
-- PostgreSQL database dump complete
--


--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

-- Started on 2025-02-22 17:48:45

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
-- TOC entry 215 (class 1259 OID 33226)
-- Name: tracks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tracks (
    "group" character varying(128) NOT NULL,
    song character varying(128) NOT NULL,
    release_date character varying(10),
    text text,
    link text
);


ALTER TABLE public.tracks OWNER TO postgres;

--
-- TOC entry 4832 (class 0 OID 33226)
-- Dependencies: 215
-- Data for Name: tracks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tracks ("group", song, release_date, text, link) FROM stdin;
\.


--
-- TOC entry 4688 (class 2606 OID 33232)
-- Name: tracks tracks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tracks
    ADD CONSTRAINT tracks_pkey PRIMARY KEY ("group", song);


-- Completed on 2025-02-22 17:48:45

--
-- PostgreSQL database dump complete
--


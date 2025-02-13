--
-- PostgreSQL database dump
--

-- Dumped from database version 15.6
-- Dumped by pg_dump version 15.6

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
-- Name: go_crm_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.go_crm_user (
    usr_id integer NOT NULL,
    usr_email text DEFAULT ''::character varying NOT NULL,
    usr_phone text DEFAULT ''::character varying NOT NULL,
    usr_username text DEFAULT ''::character varying NOT NULL,
    usr_password text DEFAULT ''::character varying NOT NULL,
    usr_created_at integer DEFAULT 0 NOT NULL,
    usr_updated_at integer DEFAULT 0 NOT NULL,
    usr_create_ip text DEFAULT ''::character varying NOT NULL,
    usr_last_login_at integer DEFAULT 0 NOT NULL,
    usr_last_login_ip text DEFAULT ''::character varying NOT NULL,
    usr_login_times integer DEFAULT 0 NOT NULL,
    usr_status smallint DEFAULT 0 NOT NULL
);


ALTER TABLE public.go_crm_user OWNER TO postgres;

--
-- Name: COLUMN go_crm_user.usr_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_id IS 'Account ID';


--
-- Name: COLUMN go_crm_user.usr_email; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_email IS 'Email';


--
-- Name: COLUMN go_crm_user.usr_phone; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_phone IS 'Phone Number';


--
-- Name: COLUMN go_crm_user.usr_username; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_username IS 'Username';


--
-- Name: COLUMN go_crm_user.usr_password; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_password IS 'Password';


--
-- Name: COLUMN go_crm_user.usr_created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_created_at IS 'Creation Time';


--
-- Name: COLUMN go_crm_user.usr_updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_updated_at IS 'Update Time';


--
-- Name: COLUMN go_crm_user.usr_create_ip; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_create_ip IS 'Creation IP';


--
-- Name: COLUMN go_crm_user.usr_last_login_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_last_login_at IS 'Last Login Time';


--
-- Name: COLUMN go_crm_user.usr_last_login_ip; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_last_login_ip IS 'Last Login IP';


--
-- Name: COLUMN go_crm_user.usr_login_times; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_login_times IS 'Login Times';


--
-- Name: COLUMN go_crm_user.usr_status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_user.usr_status IS 'Status 1:enable, 0:disable, -1:deleted';


--
-- Name: go_crm_user_usr_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.go_crm_user_usr_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.go_crm_user_usr_id_seq OWNER TO postgres;

--
-- Name: go_crm_user_usr_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.go_crm_user_usr_id_seq OWNED BY public.go_crm_user.usr_id;


--
-- Name: go_crm_userv2; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.go_crm_userv2 (
    usr_id integer NOT NULL,
    usr_email text NOT NULL,
    usr_phone text NOT NULL,
    usr_username text NOT NULL,
    usr_password text NOT NULL,
    usr_created_at integer NOT NULL,
    usr_updated_at integer NOT NULL,
    usr_create_ip text NOT NULL,
    usr_last_login_at integer NOT NULL,
    usr_last_login_ip text NOT NULL,
    usr_login_times integer NOT NULL,
    usr_status smallint NOT NULL
);


ALTER TABLE public.go_crm_userv2 OWNER TO postgres;

--
-- Name: COLUMN go_crm_userv2.usr_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_id IS 'Account ID';


--
-- Name: COLUMN go_crm_userv2.usr_email; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_email IS 'Email';


--
-- Name: COLUMN go_crm_userv2.usr_phone; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_phone IS 'Phone Number';


--
-- Name: COLUMN go_crm_userv2.usr_username; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_username IS 'Username';


--
-- Name: COLUMN go_crm_userv2.usr_password; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_password IS 'Password';


--
-- Name: COLUMN go_crm_userv2.usr_created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_created_at IS 'Creation Time';


--
-- Name: COLUMN go_crm_userv2.usr_updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_updated_at IS 'Update Time';


--
-- Name: COLUMN go_crm_userv2.usr_create_ip; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_create_ip IS 'Creation IP';


--
-- Name: COLUMN go_crm_userv2.usr_last_login_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_last_login_at IS 'Last Login Time';


--
-- Name: COLUMN go_crm_userv2.usr_last_login_ip; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_last_login_ip IS 'Last Login IP';


--
-- Name: COLUMN go_crm_userv2.usr_login_times; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_login_times IS 'Login Times';


--
-- Name: COLUMN go_crm_userv2.usr_status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_crm_userv2.usr_status IS 'Status 1:enable, 0:disable, -1:deleted';


--
-- Name: go_crm_userv2_usr_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.go_crm_userv2_usr_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.go_crm_userv2_usr_id_seq OWNER TO postgres;

--
-- Name: go_crm_userv2_usr_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.go_crm_userv2_usr_id_seq OWNED BY public.go_crm_userv2.usr_id;


--
-- Name: go_db_role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.go_db_role (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    role_name character varying(50),
    role_note text
);


ALTER TABLE public.go_db_role OWNER TO postgres;

--
-- Name: COLUMN go_db_role.id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_db_role.id IS '''Primary key''';


--
-- Name: go_db_role_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.go_db_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.go_db_role_id_seq OWNER TO postgres;

--
-- Name: go_db_role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.go_db_role_id_seq OWNED BY public.go_db_role.id;


--
-- Name: go_db_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.go_db_users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    uuid uuid NOT NULL,
    user_name character varying(50) NOT NULL,
    is_actived boolean
);


ALTER TABLE public.go_db_users OWNER TO postgres;

--
-- Name: go_db_users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.go_db_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.go_db_users_id_seq OWNER TO postgres;

--
-- Name: go_db_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.go_db_users_id_seq OWNED BY public.go_db_users.id;


--
-- Name: go_user_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.go_user_roles (
    user_id bigint NOT NULL,
    user_uuid uuid NOT NULL,
    role_id bigint NOT NULL
);


ALTER TABLE public.go_user_roles OWNER TO postgres;

--
-- Name: COLUMN go_user_roles.role_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.go_user_roles.role_id IS '''Primary key''';


--
-- Name: go_crm_user usr_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_crm_user ALTER COLUMN usr_id SET DEFAULT nextval('public.go_crm_user_usr_id_seq'::regclass);


--
-- Name: go_crm_userv2 usr_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_crm_userv2 ALTER COLUMN usr_id SET DEFAULT nextval('public.go_crm_userv2_usr_id_seq'::regclass);


--
-- Name: go_db_role id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_db_role ALTER COLUMN id SET DEFAULT nextval('public.go_db_role_id_seq'::regclass);


--
-- Name: go_db_users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_db_users ALTER COLUMN id SET DEFAULT nextval('public.go_db_users_id_seq'::regclass);


--
-- Data for Name: go_crm_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.go_crm_user (usr_id, usr_email, usr_phone, usr_username, usr_password, usr_created_at, usr_updated_at, usr_create_ip, usr_last_login_at, usr_last_login_ip, usr_login_times, usr_status) FROM stdin;
\.


--
-- Data for Name: go_crm_userv2; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.go_crm_userv2 (usr_id, usr_email, usr_phone, usr_username, usr_password, usr_created_at, usr_updated_at, usr_create_ip, usr_last_login_at, usr_last_login_ip, usr_login_times, usr_status) FROM stdin;
\.


--
-- Data for Name: go_db_role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.go_db_role (id, created_at, updated_at, deleted_at, role_name, role_note) FROM stdin;
\.


--
-- Data for Name: go_db_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.go_db_users (id, created_at, updated_at, deleted_at, uuid, user_name, is_actived) FROM stdin;
\.


--
-- Data for Name: go_user_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.go_user_roles (user_id, user_uuid, role_id) FROM stdin;
\.


--
-- Name: go_crm_user_usr_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.go_crm_user_usr_id_seq', 1, false);


--
-- Name: go_crm_userv2_usr_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.go_crm_userv2_usr_id_seq', 1, false);


--
-- Name: go_db_role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.go_db_role_id_seq', 1, false);


--
-- Name: go_db_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.go_db_users_id_seq', 1, false);


--
-- Name: go_crm_user go_crm_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_crm_user
    ADD CONSTRAINT go_crm_user_pkey PRIMARY KEY (usr_id);


--
-- Name: go_crm_userv2 go_crm_userv2_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_crm_userv2
    ADD CONSTRAINT go_crm_userv2_pkey PRIMARY KEY (usr_id);


--
-- Name: go_db_role go_db_role_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_db_role
    ADD CONSTRAINT go_db_role_pkey PRIMARY KEY (id);


--
-- Name: go_db_users go_db_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_db_users
    ADD CONSTRAINT go_db_users_pkey PRIMARY KEY (id, uuid);


--
-- Name: go_user_roles go_user_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_user_roles
    ADD CONSTRAINT go_user_roles_pkey PRIMARY KEY (user_id, user_uuid, role_id);


--
-- Name: go_db_users uni_go_db_users_uuid; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_db_users
    ADD CONSTRAINT uni_go_db_users_uuid UNIQUE (uuid);


--
-- Name: idx_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_email ON public.go_crm_user USING btree (usr_email);


--
-- Name: idx_go_db_role_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_go_db_role_deleted_at ON public.go_db_role USING btree (deleted_at);


--
-- Name: idx_go_db_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_go_db_users_deleted_at ON public.go_db_users USING btree (deleted_at);


--
-- Name: idx_phone; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_phone ON public.go_crm_user USING btree (usr_phone);


--
-- Name: idx_username; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_username ON public.go_crm_user USING btree (usr_username);


--
-- Name: idx_uuid; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_uuid ON public.go_db_users USING btree (uuid);


--
-- Name: go_user_roles fk_go_user_roles_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_user_roles
    ADD CONSTRAINT fk_go_user_roles_role FOREIGN KEY (role_id) REFERENCES public.go_db_role(id);


--
-- Name: go_user_roles fk_go_user_roles_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.go_user_roles
    ADD CONSTRAINT fk_go_user_roles_user FOREIGN KEY (user_id, user_uuid) REFERENCES public.go_db_users(id, uuid);


--
-- PostgreSQL database dump complete
--


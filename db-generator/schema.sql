--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    category_id smallint NOT NULL,
    category_name character varying(15),
    description text
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: transaction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction (
    transact_id character varying(6) NOT NULL,
    type_transact character varying(15),
    description text,
    nominal money,
    time_of_transaction date,
    category_id smallint
);


ALTER TABLE public.transaction OWNER TO postgres;

--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.categories (category_id, category_name, description) VALUES (1, 'Pemasukan', 'Pemasukan yang diterima user baik berupa Gaji, Tunjangan maupun Bonus');
INSERT INTO public.categories (category_id, category_name, description) VALUES (2, 'Pengeluaran', 'Pengeluaran yang dilakukan oleh user baik berupa Sewa kost, Makan, Pakaian, Nonton Bioskop');


--
-- Data for Name: transaction; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transaction (transact_id, type_transact, description, nominal, time_of_transaction, category_id) VALUES ('t-0001', 'Gaji', 'Gaji Bulan Mei terhitung dari tanggal 28 april - 28 mei', 'Rp7.500.000,00', '2023-05-28', 1);
INSERT INTO public.transaction (transact_id, type_transact, description, nominal, time_of_transaction, category_id) VALUES ('t-0002', 'Bonus', 'Bonus Lembur pada periode 10 mei - 15 mei', 'Rp1.000.000,00', '2023-05-28', 1);
INSERT INTO public.transaction (transact_id, type_transact, description, nominal, time_of_transaction, category_id) VALUES ('t-0003', 'Sewa Kost', 'Pembayaran Sewa Kost Bulan 6', 'Rp1.500.000,00', '2023-05-30', 2);
INSERT INTO public.transaction (transact_id, type_transact, description, nominal, time_of_transaction, category_id) VALUES ('t-0004', 'Makan', 'Pembayaran Makan Warteg Bahari Putra', 'Rp20.000,00', '2023-06-02', 2);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (category_id);


--
-- Name: transaction transaction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (transact_id);


--
-- Name: transaction category_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT category_id_fk FOREIGN KEY (category_id) REFERENCES public.categories(category_id);

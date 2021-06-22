INSERT INTO advert (id, title, description, price, created_at) VALUES
(DEFAULT, 'Самолет', 'Самолет-описание', 10000000.00, DEFAULT);

INSERT INTO photo (id, id_advert, url) VALUES
(DEFAULT, currval('advert_id_seq'), 'https://st.depositphotos.com/1008438/3928/i/600/depositphotos_39288517-stock-photo-clouds-and-sky-as-seen.jpg' ),
(DEFAULT, currval('advert_id_seq'), 'https://st.depositphotos.com/1008438/3928/i/600/depositphotos_39288517-stock-photo-clouds-and-sky-as-seen.jpg' ),
(DEFAULT, currval('advert_id_seq'), 'https://st.depositphotos.com/1008438/3928/i/600/depositphotos_39288517-stock-photo-clouds-and-sky-as-seen.jpg' );


INSERT INTO advert (id, title, description, price, created_at) VALUES
(DEFAULT, 'Вертолет', 'Вертолет-описание', 12000000.00, DEFAULT);

INSERT INTO photo (id, id_advert, url) VALUES
(DEFAULT, currval('advert_id_seq'), 'https://upload.wikimedia.org/wikipedia/commons/2/25/Robinson-R44_2.jpg' ),
(DEFAULT, currval('advert_id_seq'), 'https://upload.wikimedia.org/wikipedia/commons/2/25/Robinson-R44_2.jpg' ),
(DEFAULT, currval('advert_id_seq'), 'https://upload.wikimedia.org/wikipedia/commons/2/25/Robinson-R44_2.jpg' );


INSERT INTO advert (id, title, description, price, created_at) VALUES
(DEFAULT, 'Подлодка', 'Подлодка-описание', 300000000.00, DEFAULT);

INSERT INTO photo (id, id_advert, url) VALUES
(DEFAULT, currval('advert_id_seq'), 'http://santehprospekt.ru/wp-content/uploads/2017/04/galun5.gif' ),
(DEFAULT, currval('advert_id_seq'), 'http://santehprospekt.ru/wp-content/uploads/2017/04/galun5.gif' ),
(DEFAULT, currval('advert_id_seq'), 'http://santehprospekt.ru/wp-content/uploads/2017/04/galun5.gif' );


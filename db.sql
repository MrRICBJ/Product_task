CREATE TABLE orders (
                        order_id BIGSERIAL PRIMARY KEY,
                        cour_id BIGINT REFERENCES couriers (courier_id) DEFAULT NULL,
                        weight FLOAT NOT NULL,
                        regions INT NOT NULL,
                        delivery_hours TEXT[] NOT NULL,
                        cost INT NOT NULL,
                        completed_time TIMESTAMPTZ DEFAULT NULL);

INSERT INTO orders (weight, regions, delivery_hours, cost, completed_time) VALUES(10.5, 3, ARRAY['10:00-12:00', '14:00-16:00'], 200, '2023-04-29 15:30:00'),
                                                                                 (8.2, 2, ARRAY['09:00-11:00', '13:00-15:00'], 150, '2023-04-30 09:45:00'),
                                                                                 (5.7, 1, ARRAY['16:00-18:00', '19:00-21:00'], 120, '2023-05-01 17:20:00'),
                                                                                 (12.1, 4, ARRAY['12:00-14:00', '15:00-17:00'], 250, '2023-05-02 13:15:00'),
                                                                                 (12.1, 4, ARRAY['12:00-14:00', '15:00-17:00'], 250, '2023-05-02 13:15:00');
INSERT INTO orders (cour_id, weight, regions, delivery_hours, cost) VALUES(1, 10, 3, ARRAY['10:00-12:00', '14:00-16:00'], 300),
                                                                 (2, 8, 2, ARRAY['09:00-11:00', '13:00-15:00'], 50),
                                                                 (2, 5, 1, ARRAY['16:00-18:00', '19:00-21:00'], 10),
                                                                 (3, 12, 4, ARRAY['12:00-14:00', '15:00-17:00'], 20),
                                                                 (1, 15, 4, ARRAY['12:00-14:00', '15:00-17:00'], 500);

CREATE TABLE couriers (
                          courier_id BIGSERIAL PRIMARY KEY,
                          courier_type VARCHAR(4) NOT NULL,
                          regions INT[] NOT NULL,
                          working_hours TIME[] NOT NULL);

INSERT INTO couriers (courier_type, regions, working_hours)
VALUES
    ('BIKE', ARRAY[1, 2, 3], ARRAY[
        TIME '08:00:00', TIME '12:00:00',
     TIME '13:00:00', TIME '17:00:00',
     TIME '18:00:00', TIME '22:00:00'
         ]),
    ('CAR', ARRAY[4, 5], ARRAY[
        TIME '09:00:00', TIME '18:00:00',
     TIME '19:00:00', TIME '22:00:00'
         ]),
    ('FOOT', ARRAY[1, 3, 5], ARRAY[
        TIME '10:00:00', TIME '14:00:00',
     TIME '15:00:00', TIME '19:00:00'
         ]);
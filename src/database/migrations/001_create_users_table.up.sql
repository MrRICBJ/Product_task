CREATE TABLE couriers (
                          courier_id BIGSERIAL PRIMARY KEY,
                          courier_type VARCHAR(4) NOT NULL,
                          regions INT[] NOT NULL,
                          working_hours TIME[] NOT NULL);

CREATE TABLE orders (
                        order_id BIGSERIAL PRIMARY KEY,
                        cour_id BIGINT REFERENCES couriers (courier_id) DEFAULT NULL,
                        weight FLOAT NOT NULL,
                        regions INT NOT NULL,
                        delivery_hours VARCHAR[] NOT NULL,
                        cost INT NOT NULL,
                        completed_time TIMESTAMPTZ DEFAULT NULL);
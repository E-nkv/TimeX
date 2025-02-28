INSERT INTO categories (id, name, parent_category_id) VALUES 
	(1, 'study', NULL),
	(2, 'web dev', 1),
	(3, 'backend', 2),
	(4, 'frontend', 2),
	(5, 'golang', 3),
	(6, 'entrepreneurship', 1),
	(7, 'work', NULL),
	(8, 'my-job', 7)
;

-- seed sessions between last year and 15 - 02 - 2025 at 19:44
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2025-02-15 10:00:00.0000-05', '2025-02-15 10:45:00.0000-05', 2, 3);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2025-02-15 11:00:00.0000-05', '2025-02-15 12:30:00.0000-05', 3, 1);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2025-02-15 14:00:00.0000-05', '2025-02-15 14:20:00.0000-05', 3, 5);

-- Yesterday
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2025-02-14 15:00:00.0000-05', '2025-02-14 16:15:00.0000-05', 3, 2);

-- The day before yesterday
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2025-02-13 09:00:00.0000-05', '2025-02-13 10:00:00.0000-05', 4, 4);

-- Last Month (January)
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2025-01-10 12:00:00.0000-05', '2025-01-10 13:00:00.0000-05', 3, 1);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2025-01-25 16:00:00.0000-05', '2025-01-25 17:30:00.0000-05', 3, 3);

-- Last Last Month (December)
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-12-05 18:00:00.0000-05', '2024-12-05 19:45:00.0000-05', 3, 5);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-12-15 20:00:00.0000-05', '2024-12-15 21:00:00.0000-05', 3, 2);

-- Year 2024
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-02-10 08:00:00.0000-05', '2024-02-10 08:30:00.0000-05', 3, 4);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-03-15 11:00:00.0000-05', '2024-03-15 12:00:00.0000-05', 2, 1);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-04-20 14:00:00.0000-05', '2024-04-20 15:30:00.0000-05', 2, 3);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-05-01 16:00:00.0000-05', '2024-05-01 17:00:00.0000-05', 2, 5);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-06-10 18:00:00.0000-05', '2024-06-10 18:45:00.0000-05', 2, 2);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-07-15 09:00:00.0000-05', '2024-07-15 10:00:00.0000-05', 4, 4);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-08-20 12:00:00.0000-05', '2024-08-20 13:30:00.0000-05', 4, 1);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-09-01 15:00:00.0000-05', '2024-09-01 16:00:00.0000-05', 4, 3);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-10-10 17:00:00.0000-05', '2024-10-10 17:30:00.0000-05', 4, 5);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-11-15 19:00:00.0000-05', '2024-11-15 20:00:00.0000-05', 4, 2);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-02-01 10:00:00.0000-05', '2024-02-01 11:30:00.0000-05', 3, 4);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-03-10 13:00:00.0000-05', '2024-03-10 14:00:00.0000-05', 3, 1);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-04-15 16:00:00.0000-05', '2024-04-15 17:30:00.0000-05', 3, 3);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-05-20 18:00:00.0000-05', '2024-05-20 19:00:00.0000-05', 1, 5);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-06-01 09:00:00.0000-05', '2024-06-01 09:45:00.0000-05', 1, 2);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-07-10 11:00:00.0000-05', '2024-07-10 12:00:00.0000-05', 1, 4);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-08-15 14:00:00.0000-05', '2024-08-15 15:30:00.0000-05', 1, 1);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-09-20 16:00:00.0000-05', '2024-09-20 17:00:00.0000-05', 1, 3);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-10-01 18:00:00.0000-05', '2024-10-01 18:30:00.0000-05', 1, 5);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-11-10 20:00:00.0000-05', '2024-11-10 21:00:00.0000-05', 4, 2);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-02-05 08:00:00.0000-05', '2024-02-05 09:30:00.0000-05', 4, 4);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-03-01 10:00:00.0000-05', '2024-03-01 11:00:00.0000-05', 4, 1);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-04-10 13:00:00.0000-05', '2024-04-10 14:30:00.0000-05', 4, 3);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-05-15 15:00:00.0000-05', '2024-05-15 16:00:00.0000-05', 4, 5);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-06-20 17:00:00.0000-05', '2024-06-20 17:45:00.0000-05', 4, 2);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-07-01 19:00:00.0000-05', '2024-07-01 20:00:00.0000-05', 3, 4);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-08-10 09:00:00.0000-05', '2024-08-10 10:30:00.0000-05', 3, 1);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-09-15 11:00:00.0000-05', '2024-09-15 12:00:00.0000-05', 3, 3);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-10-20 14:00:00.0000-05', '2024-10-20 14:30:00.0000-05', 3, 5);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-11-01 16:00:00.0000-05', '2024-11-01 17:00:00.0000-05', 3, 2);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-12-10 18:00:00.0000-05', '2024-12-10 19:30:00.0000-05', 4, 4);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-02-15 20:00:00.0000-05', '2024-02-15 21:00:00.0000-05', 4, 1);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-03-05 09:00:00.0000-05', '2024-03-05 10:30:00.0000-05', 4, 3);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-04-01 11:00:00.0000-05', '2024-04-01 12:00:00.0000-05', 4, 5);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-05-10 13:00:00.0000-05', '2024-05-10 13:30:00.0000-05', 4, 2);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-06-15 15:00:00.0000-05', '2024-06-15 16:00:00.0000-05', 4, 4);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-07-20 17:00:00.0000-05', '2024-07-20 18:30:00.0000-05', 4, 1);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-08-01 19:00:00.0000-05', '2024-08-01 20:00:00.0000-05', 3, 3);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-09-10 08:00:00.0000-05', '2024-09-10 08:30:00.0000-05', 3, 5);
INSERT INTO sessions (start, "end", category_id, focus) VALUES
('2024-10-15 10:00:00.0000-05', '2024-10-15 11:00:00.0000-05', 3, 2);

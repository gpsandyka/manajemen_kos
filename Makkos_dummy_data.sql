INSERT INTO residents (id, stay_type, bills, full_name, "nik_number", username, "password", created_at, updated_at, deleted_at)
    VALUES 
    (3, 'PERMANENT', 800000, 'Bale', '327314010119980001', 'baledog', '$2a$10$07LNZC1ycOH5rBb8BFKuzeB3RtvZ5q5SiaR2JUEeAjG.zEiwj9yBa', null, null, null), /*Password hashed. "ngapung"*/
    (1, 'TEMPORARY', 400000, 'Syarif', '327314020219990001', NULL, NULL, null, null, null),
    (2, 'PERMANENT', 0, 'Riri', '327314030319980001', 'gambari', '$2a$10$s6a4GelHiyfTPLKbkhCxTefZkPXJyu5qh8bQi2PdnfTi0K2UIevay', null, null, null); /*Password hashed. "gambarisu"*/

INSERT INTO rooms (id, resident_id, price, start_occupied_date, end_occupied_date, status, qr_api_code, created_at, updated_at, deleted_at)
    VALUES 
    (100, 1, 80000, '2022-04-04', NULL, 'OCCUPIED', '2xce23r2', null, null, null),
    (201, 2, 80000, '2022-01-01', '2022-02-02', 'OCCUPIED', '98ufdhuw', null, null, null),
    (302, 3, 80000, '2022-05-05', NULL, 'OCCUPIED', '68fsdw34', null, null, null),
    (105, NULL, 80000, '2022-02-02', NULL, 'UNOCCUPIED', '55hgt5u3', null, null, null);

INSERT INTO invoices (id, resident_id, room_id, total_price, pay_date, payment_type, pay_type, created_at, updated_at, deleted_at)
    VALUES 
    (1, 1, 100, 400000, '2022-03-03', 'CASH', 'ROOMS', null, null, null),
    (2, 2, 201, 400000, '2022-02-02', 'OVO', 'ROOMS', null, null, null);

INSERT INTO admins (id, admin_name, username, "password", created_at, updated_at, deleted_at)
    VALUES 
    (3, 'Miyuki', 'Shirogane', '$2a$10$A7vS0z5wjwGwfULnkTeQSOn6m6Sd.B64ZLNSSKl3UgD8HwDRdRHm2', null, null, null), /*Password hashed. "kaguyadaisuki"*/
    (4, 'Kaguya', 'Shinomiya', '$2a$10$aKKHLAvBfXyn6lLfzkPRy..eUd8C5kA7M2zqs4qarTOKkkSK8cBu.', null, null, null); /*Password hashed. "miyukidaisuki"*/
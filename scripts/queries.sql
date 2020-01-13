/* 
insert data into the hospital table 
*/
INSERT INTO hospital (name, max_patient_amount)
VALUES ('Artemis Hospital', 1000),
       ('Medanta Hospital', 1000);

/* 
insert data into the location table 
*/
INSERT INTO location (name, hospital_id)
VALUES ('Inpatient Ward', 1),
       ('Laboratory', 2),
       ('Registration Desk', 1),
       ('Isolation Ward', 2);


/* 
insert data into the patient table 
*/
INSERT INTO patient (name, illness, birth_date, location_id)
VALUES ('Ramesh', 'Tuberculosis', '1988-05-19', 1),
       ('Mukesh', 'Cold', '2000-12-08', 1),
       ('Rachit', 'Lung Cancer', '1987-11-01', 2),
       ('Gopal', 'Asthma', '1999-08-17', 2),
       ('Ishaan', 'Jaundice', '2003-03-24', 2),
       ('Laxya', 'Dyslexia', '1988-12-25', 3),
       ('Devjyoti', 'Dengue', '1979-11-14', 3),
       ('Aravind', 'Flu', '2002-06-21', 4),
       ('Aaryak', 'Typhoid', '2005-01-20', 4);
       
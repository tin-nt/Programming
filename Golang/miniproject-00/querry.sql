DROP TABLE IF EXISTS cities;
CREATE TABLE cities(id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255), population INT);
INSERT INTO cities(name, population) VALUES('Bratislava', 432000);
INSERT INTO cities(name, population) VALUES('Budapest', 1759000);
INSERT INTO cities(name, population) VALUES('Prague', 1280000);
INSERT INTO cities(name, population) VALUES('Warsaw', 1748000);
INSERT INTO cities(name, population) VALUES('Los Angeles', 3971000);
INSERT INTO cities(name, population) VALUES('New York', 8550000);
INSERT INTO cities(name, population) VALUES('Edinburgh', 464000);
INSERT INTO cities(name, population) VALUES('Zimbabwe', 3213567);
INSERT INTO cities(name, population) VALUES('Ha Noi', 97584563);
INSERT INTO cities(name, population) VALUES('HCM', 76592123);
INSERT INTO cities(name, population) VALUES('caddd', 54567897);
INSERT INTO cities(name, population) VALUES('fgnc', 77654213);
INSERT INTO cities(name, population) VALUES('.,mkjh', 4567630);
INSERT INTO cities(name, population) VALUES('xfgl', 56935621);
INSERT INTO cities(name, population) VALUES('yuuer', 3671043474600);
INSERT INTO cities(name, population) VALUES('ghuotyuyt', 76568576);

select * from cities where id = 2
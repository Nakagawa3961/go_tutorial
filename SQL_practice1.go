//[課題６]
mysql> use db_practice;
// Reading table information for completion of table and column names
// You can turn off this feature to get a quicker startup with -A
// Database changed

mysql> CREATE TABLE IF NOT EXISTS `db_practice`.`practice_table` (
	     ->   `id` VARCHAR(255) NOT NULL,
	     ->   `name` VARCHAR(255) NOT NULL,
	     ->   `sex` VARCHAR(255) NOT NULL,
	     ->   `age` INT NOT NULL,
	     ->   `post` VARCHAR(255) NOT NULL,
         ->   PRIMARY KEY (`id`)
	     -> );

mysql> show tables;
// +-----------------------+
// | Tables_in_db_practice |
// +-----------------------+
// | practice_table        |
// +-----------------------+
// 1 row in set (0.01 sec)

mysql> insert into practice_table(id,name,sex,age,post) values(00001,'Satou','man','35','soumubu');
mysql> insert into practice_table(id,name,sex,age,post) values(00002,'Suzuki','man','40','zinnzibu');
mysql> insert into practice_table(id,name,sex,age,post) values(00003,'Takahashi','man','32','eigyoubu');
mysql> insert into practice_table(id,name,sex,age,post) values(00004,'Tanaka','woman','25','kaihatubu');
mysql> insert into practice_table(id,name,sex,age,post) values(00005,'Watanabe','man','20','eigyoubu');

mysql> select * from practice_table;
// +----+-----------+-------+-----+-----------+
// | id | name      | sex   | age | post      |
// +----+-----------+-------+-----+-----------+
// | 1  | Satou     | man   |  35 | soumubu   |
// | 2  | Suzuki    | man   |  40 | zinnzibu  |
// | 3  | Takahashi | man   |  32 | eigyoubu  |
// | 4  | Tanaka    | woman |  25 | kaihatubu |
// | 5  | Watanabe  | man   |  20 | eigyoubu  |
// +----+-----------+-------+-----+-----------+
// 5 rows in set (0.00 sec)



-----------------------------------------------------------------------------------------------


//[課題７]
mysql> SELECT id,name,sex,age,post from practice_table ORDER BY age ASC;
// +----+-----------+-------+-----+-----------+
// | id | name      | sex   | age | post      |
// +----+-----------+-------+-----+-----------+
// | 5  | Watanabe  | man   |  20 | eigyoubu  |
// | 4  | Tanaka    | woman |  25 | kaihatubu |
// | 3  | Takahashi | man   |  32 | eigyoubu  |
// | 1  | Satou     | man   |  35 | soumubu   |
// | 2  | Suzuki    | man   |  40 | zinnzibu  |
// +----+-----------+-------+-----+-----------+
// 5 rows in set (0.00 sec)

mysql> SELECT id,name,sex,age,post from practice_table ORDER BY age DESC;
// +----+-----------+-------+-----+-----------+
// | id | name      | sex   | age | post      |
// +----+-----------+-------+-----+-----------+
// | 2  | Suzuki    | man   |  40 | zinnzibu  |
// | 1  | Satou     | man   |  35 | soumubu   |
// | 3  | Takahashi | man   |  32 | eigyoubu  |
// | 4  | Tanaka    | woman |  25 | kaihatubu |
// | 5  | Watanabe  | man   |  20 | eigyoubu  |
// +----+-----------+-------+-----+-----------+
// 5 rows in set (0.00 sec)



---------------------------------------------------------------------------------------



//[課題８]
mysql> select * from practice_table order by id desc limit 3;
// +----+-----------+-------+-----+-----------+
// | id | name      | sex   | age | post      |
// +----+-----------+-------+-----+-----------+
// | 5  | Watanabe  | man   |  20 | eigyoubu  |
// | 4  | Tanaka    | woman |  25 | kaihatubu |
// | 3  | Takahashi | man   |  32 | eigyoubu  |
// +----+-----------+-------+-----+-----------+
// 3 rows in set (0.00 sec)
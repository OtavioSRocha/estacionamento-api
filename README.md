```sql
CREATE TABLE estacionamento.cars (
	id INT auto_increment NOT NULL,
	model varchar(100) NOT NULL,
	licenseplate varchar(50) NOT NULL,
	CONSTRAINT cars_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

```

```sql
CREATE TABLE estacionamento.spots (
	id INT auto_increment NOT NULL,
	vehicle varchar(100) NOT NULL,
	isempty BOOL DEFAULT true NOT NULL,
	car INT NULL,
	CONSTRAINT spots_PK PRIMARY KEY (id),
	CONSTRAINT spots_FK FOREIGN KEY (car) REFERENCES estacionamento.cars(id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
CREATE INDEX spots_car_IDX USING BTREE ON estacionamento.spots (car);
```


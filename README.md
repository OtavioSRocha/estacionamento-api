## Introdução
Api para um sistema de estacionamento desenvolvido em linguagem Go, com finalidade de aprender e testar a tecnologia  


## Creating database
```sql
CREATE DATABASE estacionamento;

CREATE TABLE estacionamento.cars (
  id int NOT NULL AUTO_INCREMENT,
  model varchar(100) NOT NULL,
  licenseplate varchar(50) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY cars_UN (licenseplate)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE estacionamento.spots (
	id INT auto_increment NOT NULL,
	vehicle varchar(100) NOT NULL,
	isempty BOOL DEFAULT true NOT NULL,
	car varchar(50) NULL,
	CONSTRAINT spots_PK PRIMARY KEY (id),
	CONSTRAINT spots_FK FOREIGN KEY (car) REFERENCES estacionamento.cars(licenseplate)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
CREATE INDEX spots_car_IDX USING BTREE ON estacionamento.spots (car);
```

## Rotas

### Listar todas as vagas
```
[GET] /vagas
```

### Listar uma vaga por ID
```
[GET] /vagas/{ID}
```

### Criar uma nova vaga
```
[POST] /vagas
```
Body
```json
{
	"vehicle": "bike",
	"isempty": true
}
```
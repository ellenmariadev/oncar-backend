# Oncar Backend
This repository corresponds to the backend made in **Go** for the [frontend application](https://github.com/ellenmariadev/oncar-frontend).\
To learn more about the requirements for this project, go to: https://github.com/oncarapp/job_challenge 

If you like, use **[insomnia](https://github.com/ellenmariadev/oncar-backend/blob/main/data/insomnia.json)** - <kbd>data/insomnia.json</kbd>. 

### ○ [ Vehicles ] Routes 

| Method | Route | Description |
| ------ | ----- | ----------- |
| **`GET`** | <kbd>api/vehicles</kbd> | Returns all vehicles. |
|  **`GET`** | <kbd>api/vehicle/id</kbd> | Returns a specific vehicle |
|  **`POST`** | <kbd>api/vehicle</kbd> | Create a new vehicle.  |
|  **`PATCH`** | <kbd>api/vehicle/id</kbd> | Update vehicle.
|  **`DELETE`** | <kbd>api/vehicle/id</kbd> | Delete vehicle.

Body JSON Data Request (**POST** Vehicle)
```json
{
    "brand": "Audi",
    "model": "A4",
    "price": 40000,
    "year": 2016
}

```
   

### ○ [ Leads ] Routes 

| Method | Route | Description |
| ------ | ----- | ----------- |
| **`GET`** | <kbd>api/leads</kbd> | Returns all leads. |
|  **`POST`** | <kbd>api/lead</kbd> | Create a new lead.  |

### ○ Create and connect your database 
Go to the file ./config/database.go and change the credentials.

```go
host     = "localhost"
port     = 5432
user     = "<your db user>"
password = "<your db password>"
dbName   = "<your db name>"

```



### ○ Create the tables in your database (using PostgresSQL)

```sql
  CREATE TABLE public.vehicle(
    id serial4 NOT NULL,
    "brand" varchar NULL,
    "model" varchar NULL,
    "price" decimal NULL,
    "year" integer NULL,
    CONSTRAINT vehicle_pk PRIMARY KEY (id)
);

CREATE TABLE public.lead(
    id serial4 NOT NULL,
    "name" varchar NULL,
    "email" varchar NULL,
    "phone" varchar NULL,
    "vehicleid" int4 NULL,
    CONSTRAINT lead_pk PRIMARY KEY (id),
    CONSTRAINT vehicle_fk FOREIGN KEY (vehicleid) REFERENCES public.vehicle(id)
);

```


## ⚙️ Running this project locally

<samp>
  
> **Warning** 
> Requirements: Go.

</samp>

```bash
# Clone repository
$ git clone <https://github.com/ellenmariadev/oncar-backend.git>

# Navigate to the root directory of the project
$ cd oncar-backend

# Run application 
$ go run main.go

# Server running at PORT 5000
$ <http://localhost:5000>
```

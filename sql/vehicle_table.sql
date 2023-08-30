CREATE TABLE public.vehicle(
    id serial4 NOT NULL,
    "brand" varchar NULL,
    "model" varchar NULL,
    "price" decimal NULL,
    "year" integer NULL,
    CONSTRAINT vehicle_pk PRIMARY KEY (id)
)

CREATE TABLE public.lead(
    id serial4 NOT NULL,
    "name" varchar NULL,
    "email" varchar NULL,
    "phone" varchar NULL,
    "vehicleid" int4 NULL,
    CONSTRAINT lead_pk PRIMARY KEY (id),
    CONSTRAINT vehicle_fk FOREIGN KEY (vehicleid) REFERENCES public.vehicle(id)
)
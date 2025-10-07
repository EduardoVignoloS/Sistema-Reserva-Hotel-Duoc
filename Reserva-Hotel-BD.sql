-- No crear la base en Render, ya existe.
-- Simplemente conecta a tu base con Go o pgAdmin.

DROP TABLE IF EXISTS Notificacion CASCADE;
DROP TABLE IF EXISTS Pago CASCADE;
DROP TABLE IF EXISTS Reserva CASCADE;
DROP TABLE IF EXISTS Habitacion CASCADE;
DROP TABLE IF EXISTS Usuario CASCADE;

-- Crear tipos ENUM
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tipo_usuario') THEN
        CREATE TYPE tipo_usuario AS ENUM ('cliente', 'admin');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'estado_habitacion') THEN
        CREATE TYPE estado_habitacion AS ENUM ('disponible', 'ocupada', 'mantenimiento');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'estado_reserva') THEN
        CREATE TYPE estado_reserva AS ENUM ('pendiente', 'confirmada', 'cancelada');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'metodo_pago') THEN
        CREATE TYPE metodo_pago AS ENUM ('tarjeta', 'transferencia', 'efectivo');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'estado_pago') THEN
        CREATE TYPE estado_pago AS ENUM ('pendiente', 'aprobado', 'rechazado');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tipo_notificacion') THEN
        CREATE TYPE tipo_notificacion AS ENUM ('confirmacion', 'recordatorio', 'cancelacion');
    END IF;
END $$;

-- Crear tablas
CREATE TABLE Usuario (
    id_cliente SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    telefono VARCHAR(20),
    password VARCHAR(255),
    typeC tipo_usuario DEFAULT 'cliente',
    fecha_registro DATE DEFAULT CURRENT_DATE
);

CREATE TABLE Habitacion (
    id_habitacion SERIAL PRIMARY KEY,
    numero VARCHAR(10) NOT NULL UNIQUE,
    tipo VARCHAR(50) NOT NULL,
    capacidad INT NOT NULL,
    precio DECIMAL(10,2) NOT NULL,
    estado estado_habitacion DEFAULT 'disponible'
);

CREATE TABLE Reserva (
    id_reserva SERIAL PRIMARY KEY,
    id_cliente INT NOT NULL REFERENCES Usuario(id_cliente),
    fecha_inicio DATE NOT NULL,
    fecha_fin DATE NOT NULL,
    estado estado_reserva DEFAULT 'pendiente',
    numero_habitacion VARCHAR(10),
    fecha_reserva TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Pago (
    id_pago SERIAL PRIMARY KEY,
    id_reserva INT NOT NULL REFERENCES Reserva(id_reserva),
    monto DECIMAL(10,2) NOT NULL,
    metodo metodo_pago NOT NULL,
    fecha_pago TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    estado estado_pago DEFAULT 'pendiente'
);

CREATE TABLE Notificacion (
    id_notificacion SERIAL PRIMARY KEY,
    id_reserva INT NOT NULL REFERENCES Reserva(id_reserva),
    mensaje VARCHAR(255) NOT NULL,
    fecha_envio TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tipo tipo_notificacion NOT NULL
);

-- Insertar datos de ejemplo
INSERT INTO Usuario (nombre, apellido, email, telefono) VALUES
('Juan', 'Pérez', 'juan.perez@example.com', '+56911111111'),
('María', 'González', 'maria.gonzalez@example.com', '+56922222222'),
('Pedro', 'Ramírez', 'pedro.ramirez@example.com', '+56933333333'),
('Laura', 'Torres', 'laura.torres@example.com', '+56944444444');

INSERT INTO Habitacion (numero, tipo, capacidad, precio, estado) VALUES
('101', 'simple', 1, 35000.00, 'disponible'),
('102', 'doble', 2, 55000.00, 'disponible'),
('201', 'suite', 3, 95000.00, 'disponible'),
('202', 'doble', 2, 60000.00, 'mantenimiento');




-- Consulta final
SELECT * FROM Usuario;

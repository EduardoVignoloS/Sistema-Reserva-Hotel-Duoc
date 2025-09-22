create database ReservaHotel;
use ReservaHotel;

DROP TABLE IF EXISTS Notificacion;
DROP TABLE IF EXISTS Pago;
DROP TABLE IF EXISTS Reserva;
DROP TABLE IF EXISTS Habitacion;
DROP TABLE IF EXISTS Cliente;

Commit;

CREATE TABLE Cliente (
    id_cliente INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    telefono VARCHAR(20),
    fecha_registro DATE DEFAULT (CURRENT_DATE)
);

CREATE TABLE Habitacion (
    id_habitacion INT AUTO_INCREMENT PRIMARY KEY,
    numero VARCHAR(10) NOT NULL UNIQUE,
    tipo VARCHAR(50) NOT NULL, -- simple, doble, suite
    capacidad INT NOT NULL,
    precio DECIMAL(10,2) NOT NULL,
    estado ENUM('disponible','ocupada','mantenimiento') DEFAULT 'disponible'
);

CREATE TABLE Reserva (
    id_reserva INT AUTO_INCREMENT PRIMARY KEY,
    id_cliente INT NOT NULL,
    id_habitacion INT NOT NULL,
    fecha_inicio DATE NOT NULL,
    fecha_fin DATE NOT NULL,
    estado ENUM('pendiente','confirmada','cancelada') DEFAULT 'pendiente',
    fecha_reserva TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_reserva_cliente FOREIGN KEY (id_cliente) REFERENCES Cliente(id_cliente),
    CONSTRAINT fk_reserva_habitacion FOREIGN KEY (id_habitacion) REFERENCES Habitacion(id_habitacion)
);

CREATE TABLE Pago (
    id_pago INT AUTO_INCREMENT PRIMARY KEY,
    id_reserva INT NOT NULL,
    monto DECIMAL(10,2) NOT NULL,
    metodo ENUM('tarjeta','transferencia','efectivo') NOT NULL,
    fecha_pago TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    estado ENUM('pendiente','aprobado','rechazado') DEFAULT 'pendiente',
    CONSTRAINT fk_pago_reserva FOREIGN KEY (id_reserva) REFERENCES Reserva(id_reserva)
);

CREATE TABLE Notificacion (
    id_notificacion INT AUTO_INCREMENT PRIMARY KEY,
    id_reserva INT NOT NULL,
    mensaje VARCHAR(255) NOT NULL,
    fecha_envio TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tipo ENUM('confirmacion','recordatorio','cancelacion') NOT NULL,
    CONSTRAINT fk_notificacion_reserva FOREIGN KEY (id_reserva) REFERENCES Reserva(id_reserva)
);

Commit;

INSERT INTO Cliente (nombre, apellido, email, telefono) VALUES
('Juan', 'Pérez', 'juan.perez@example.com', '+56911111111'),
('María', 'González', 'maria.gonzalez@example.com', '+56922222222'),
('Pedro', 'Ramírez', 'pedro.ramirez@example.com', '+56933333333'),
('Laura', 'Torres', 'laura.torres@example.com', '+56944444444');

INSERT INTO Habitacion (numero, tipo, capacidad, precio, estado) VALUES
('101', 'simple', 1, 35000.00, 'disponible'),
('102', 'doble', 2, 55000.00, 'disponible'),
('201', 'suite', 3, 95000.00, 'disponible'),
('202', 'doble', 2, 60000.00, 'mantenimiento');

INSERT INTO Reserva (id_cliente, id_habitacion, fecha_inicio, fecha_fin, estado) VALUES
(1, 1, '2025-10-01', '2025-10-05', 'confirmada'),
(2, 2, '2025-10-03', '2025-10-06', 'pendiente'),
(3, 3, '2025-10-10', '2025-10-15', 'cancelada'),
(4, 1, '2025-11-01', '2025-11-07', 'confirmada');

INSERT INTO Pago (id_reserva, monto, metodo, estado) VALUES
(1, 140000.00, 'tarjeta', 'aprobado'),
(2, 165000.00, 'transferencia', 'pendiente'),
(3, 475000.00, 'efectivo', 'rechazado'),
(4, 210000.00, 'tarjeta', 'aprobado');

INSERT INTO Notificacion (id_reserva, mensaje, tipo) VALUES
(1, 'Su reserva ha sido confirmada. ¡Gracias por elegirnos!', 'confirmacion'),
(1, 'Recordatorio: su check-in es el 01-10-2025.', 'recordatorio'),
(2, 'Su reserva está pendiente de pago.', 'confirmacion'),
(3, 'Lamentamos informar que su reserva fue cancelada.', 'cancelacion'),
(4, 'Su reserva fue confirmada con éxito.', 'confirmacion');

Commit;

Select * From Cliente;
Select * From Pago;
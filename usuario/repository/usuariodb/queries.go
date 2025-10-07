package usuariodb

var (
	query = `
	SELECT email FROM usuario WHERE email = :email`

	createUser = `
	INSERT INTO Usuario (
    nombre,
    apellido,
    email,
    password,
    telefono,
    typeC,
    fecha_registro
) VALUES (
    :nombre,
    :apellido,
    :email,
    :password,
    :telefono,
    :typeC,
    :fecha_registro
)
`
	login = `
	SELECT * FROM usuario WHERE email = :email`
)

package content

// Integrante represents a team member
type Integrante struct {
	Nombre  string
	Avatar  string
	Rol     string
	Detalle string
}

// Integrantes contains all the team members
var Integrantes = []Integrante{
	{
		Nombre:  "María De Los Angeles Martínez Mejías",
		Avatar:  "MM",
		Rol:     "Integrante Mujer",
		Detalle: "Apasionada por la innovación y el diseño, en constante búsqueda de aplicar sus conocimientos en proyectos reales dentro de la universidad.",
	},
	{
		Nombre:  "Felipe Andrés Maturana Araya",
		Avatar:  "FM",
		Rol:     "Representante y Directiva (Eje Industria)",
		Detalle: "Especialista analizando procesos y conectando el mundo del desarrollo de software con las necesidades inmediatas de la industria tecnológica.",
	},
	{
		Nombre:  "Pedro Javier Morales Nadal",
		Avatar:  "PM",
		Rol:     "Directiva (Eje Hardware)",
		Detalle: "Amante del diseño de hardware, microcontroladores y todas aquellas áreas de la ingeniería que trabajan en conexión directa con la electrónica.",
	},
	{
		Nombre:  "Enzo Tomás Rodríguez Pereira",
		Avatar:  "ER",
		Rol:     "Integrante",
		Detalle: "Desarrollador con entusiasmo y dedicación. Fanático de buscar herramientas nuevas y de escribir código que se adapte rápidamente.",
	},
	{
		Nombre:  "Edicson Luciano Ernesto Solar Salinas",
		Avatar:  "ES",
		Rol:     "Directiva (Eje Matemática)",
		Detalle: "Centrado en la teoría de números y en investigar cómo modelar problemas del mundo real a través de algoritmos altamente eficientes.",
	},
	{
		Nombre:  "Matías Alexander Vásquez Yelorm",
		Avatar:  "MV",
		Rol:     "Directiva (Telecomunicaciones)",
		Detalle: "Interesado en la gestión de redes, los protocolos de comunicación y las técnicas de conectividad de datos seguras y estables.",
	},
}

package content

import "html/template"

// BlogPost represents a single blog entry
type BlogPost struct {
	Slug    string
	Title   string
	Date    string
	Author  string
	Tags    []string
	Summary string
	Content template.HTML
}

// BlogPosts contains all the blog entries
var BlogPosts = []BlogPost{
	{
		Slug:    "introduccion-a-golang",
		Title:   "Introducción a Go: El lenguaje del futuro",
		Date:    "2026-03-10",
		Author:  "Equipo n!",
		Tags:    []string{"golang", "tutorial"},
		Summary: "Descubre por qué Go se ha convertido en uno de los lenguajes más populares para el desarrollo backend y sistemas distribuidos.",
		Content: `<p>Go (o Golang) es un lenguaje de programación creado por Google en 2009. Su diseño se centra en la simplicidad, la eficiencia y la concurrencia.</p>
		<h3>> Características principales</h3>
		<ul>
			<li>Compilación rápida a binarios nativos</li>
			<li>Goroutines para concurrencia ligera</li>
			<li>Sistema de tipos estático con inferencia</li>
			<li>Garbage collector eficiente</li>
			<li>Biblioteca estándar robusta</li>
		</ul>
		<p>En n! creemos que Go es una herramienta fundamental para cualquier desarrollador moderno. Su curva de aprendizaje es suave y su rendimiento es excepcional.</p>
		<h3>> ¿Por qué aprender Go?</h3>
		<p>Go es utilizado por empresas como Google, Uber, Twitch y Docker. Su ecosistema crece constantemente y la demanda de desarrolladores Go no deja de aumentar.</p>
		<pre><code>package main

import "fmt"

func main() {
    fmt.Println("¡Hola desde n!")
}</code></pre>
		<p>Si quieres comenzar tu viaje con Go, únete a nuestras sesiones semanales de programación.</p>`,
	},
	{
		Slug:    "matematicas-en-la-programacion",
		Title:   "Matemáticas en la Programación: Más allá de los números",
		Date:    "2026-03-05",
		Author:  "Equipo n!",
		Tags:    []string{"matemáticas", "algoritmos"},
		Summary: "Exploramos la conexión profunda entre las matemáticas y la programación, y cómo el pensamiento matemático mejora tu código.",
		Content: `<p>Las matemáticas y la programación comparten una raíz común: el pensamiento lógico y abstracto. En n! exploramos esta conexión todos los días.</p>
		<h3>> Álgebra lineal en Machine Learning</h3>
		<p>Los vectores, matrices y transformaciones lineales son la base de los algoritmos de inteligencia artificial. Sin álgebra lineal, no existiría el deep learning.</p>
		<h3>> Teoría de grafos en redes</h3>
		<p>Los algoritmos de Dijkstra, BFS y DFS son aplicaciones directas de la teoría de grafos que usamos diariamente en el desarrollo de software.</p>
		<h3>> Estadística en análisis de datos</h3>
		<p>Desde pruebas de hipótesis hasta regresión, la estadística es indispensable para cualquier científico de datos.</p>
		<p>En nuestra comunidad, organizamos talleres donde combinamos teoría matemática con implementación práctica en código.</p>`,
	},
	{
		Slug:    "comunidad-open-source",
		Title:   "Contribuir a Open Source: Guía para principiantes",
		Date:    "2026-02-28",
		Author:  "Equipo n!",
		Tags:    []string{"open-source", "github", "comunidad"},
		Summary: "Una guía paso a paso para hacer tu primera contribución a proyectos de código abierto desde nuestra comunidad universitaria.",
		Content: `<p>Contribuir a proyectos open source es una de las mejores formas de aprender, crecer como desarrollador y construir tu portafolio profesional.</p>
		<h3>> Paso 1: Encuentra un proyecto</h3>
		<p>Busca proyectos que te interesen en GitHub. Filtra por etiquetas como "good first issue" o "help wanted".</p>
		<h3>> Paso 2: Lee la documentación</h3>
		<p>Antes de contribuir, lee el README, el CONTRIBUTING.md y el código de conducta del proyecto.</p>
		<h3>> Paso 3: Fork y Clone</h3>
		<pre><code>git fork https://github.com/proyecto/ejemplo
git clone https://github.com/tu-usuario/ejemplo
git checkout -b mi-feature</code></pre>
		<h3>> Paso 4: Haz tu cambio</h3>
		<p>Implementa tu mejora, arregla el bug o actualiza la documentación. Asegúrate de seguir las convenciones del proyecto.</p>
		<h3>> Paso 5: Pull Request</h3>
		<p>Envía tu PR con una descripción clara de tus cambios. ¡Y celebra tu contribución!</p>
		<p>En n! organizamos sesiones mensuales de "Open Source Fridays" donde contribuimos juntos a proyectos relevantes.</p>`,
	},
}

// FindPost searches for a blog post by its slug
func FindPost(slug string) *BlogPost {
	for i := range BlogPosts {
		if BlogPosts[i].Slug == slug {
			return &BlogPosts[i]
		}
	}
	return nil
}

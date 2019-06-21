package repo

type Libro struct {
	ID     string `json:"id"`
	Autor  string `json:"autor"`
	Titulo string `json:"titulo"`
}

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

var libros map[string]Libro

func initRepo() {
	if libros == nil {
		libros = make(map[string]Libro)
		InsertarLibro(Libro{ID: "111", Titulo: "Alexandre Dumas", Autor: "Conde de Montecristo"})
		InsertarLibro(Libro{ID: "121", Titulo: "Patrick Suskind", Autor: "El Perfume"})
		InsertarLibro(Libro{ID: "112", Titulo: "Miguel de Cervantes", Autor: "El Quijote"})
		InsertarLibro(Libro{ID: "131", Titulo: "Fernando Trujillo", Autor: "La Biblia de los Caidos"})
		InsertarLibro(Libro{ID: "123", Titulo: "Mary Shelley", Autor: "Frankenstein"})
	}
}

func GetLibro(id string) Libro {
	initRepo()

	libro := libros[id]

	return libro
}

func InsertarLibro(libro Libro) {
	initRepo()

	libros[libro.ID] = libro
}

func AllLibros() []Libro {
	initRepo()

	var librosSlice []Libro

	for _, libro := range libros {
		librosSlice = append(librosSlice, libro)
	}

	return librosSlice
}

func DeleteLibro(id string) error {
	initRepo()

	if _, existe := libros[id]; !existe {
		return Error{
			Message: "El libro no existe en el repositorio",
		}
	}

	delete(libros, id)
	return nil
}

func UpdateLibro(id string, libro Libro) error {
	libro.ID = id // for avoiding user to change ID

	if _, existe := libros[id]; !existe {
		return Error{
			Message: "El libro no existe en el repositorio",
		}
	}

	libros[id] = libro

	return nil
}

package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

type taskList struct {
	tasks []*task
}

func (t *task) marcarCompleto() {
	t.completado = true
}

func (t *task) actualizarDescripcion(desc string) {
	t.descripcion = desc
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func (tl *taskList) agregarALista(tarea *task) {
	tl.tasks = append(tl.tasks, tarea)
}

func (tl *taskList) eliminarDeLista(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *taskList) imprimirLista() {
	for _, tarea := range tl.tasks {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Descripcion", tarea.descripcion)
	}
}

func (tl *taskList) imprimirListaCompletados() {
	for _, tarea := range tl.tasks {
		if tarea.completado {
			fmt.Println("Nombre", tarea.nombre)
			fmt.Println("Descripcion", tarea.descripcion)
		}
	}
}

func main() {
	t := &task{
		nombre:      "Completar el curso de Go",
		descripcion: "Pasar el examen",
	}

	t2 := &task{
		nombre:      "Completar el curso de Angular",
		descripcion: "Pasar el examen",
	}

	lista := &taskList{
		tasks: []*task{
			t,
		},
	}

	fmt.Printf("%+v \n", lista.tasks[0])
	lista.agregarALista(t2)
	lista.eliminarDeLista(1)

	lista.imprimirLista()
	lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)
	mapaTareas["Nico"] = lista
}

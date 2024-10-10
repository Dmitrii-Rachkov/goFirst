package composite

import "fmt"

/*
Концептуальный пример
Давайте попробуем понять паттерн Компоновщик, используя для примера файловую систему ОС. Внутри нее есть два типа
объектов: файлы и папки. В некоторых случаях они должны восприниматься как одно и то же. Здесь нам пригодится паттерн
проектирования Компоновщик.

Представьте, что вам нужно провести поиск по конкретному ключевому слову в вашей файловой системе. Такая операция
поиска применяется в равной степени и к файлам, и к папкам. В случае файла, она только проверит содержимое файла,
а в случае папки – обработает все файлы этой папки для нахождения ключевого слова.
*/

// component.go: Интерфейс компонента
type Component interface {
	search(string)
}

// folder.go: Компоновщик
type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

// file.go: Лист
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}

// main.go: Клиентский код
func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}

// output.txt: Результат выполнения
/*
Serching recursively for keyword rose in folder Folder2
Searching for keyword rose in file File2
Searching for keyword rose in file File3
Serching recursively for keyword rose in folder Folder1
Searching for keyword rose in file File1
*/

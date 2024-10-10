package prototype

import "fmt"

/*
Концептуальный пример
Давайте попробуем разобрать паттерн Прототип, используя для примера файловую систему ОС.
Файловая система является рекурсивной – папки содержат файлы и другие папки, которые, в свою очередь, могут
содержать файлы и папки, и так далее.

Каждый файл и папка могут быть представлены интерфейсом inode. Он имеет функцию clone.

Обе структуры файла и папки — file и folder — реализуют функции print и clone, поскольку они имеют тип inode.
Также, обратите внимание на функцию clone в file и folder. Функция clone в обеих случаях возвращает копию
соответствующего файла или папки. Во время клонирования мы добавляем ключевое слово «_clone» в поле имени.
*/

// inode.go: Интерфейс прототипа
type Inode interface {
	print(string)
	clone() Inode
}

// file.go: Конкретный прототип
type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}

// folder.go: Конкретный прототип
type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

// main.go: Клиентский код
func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}

//  output.txt: Результат выполнения

/*
Printing hierarchy for Folder2
  Folder2
    Folder1
        File1
    File2
    File3

Printing hierarchy for clone Folder
  Folder2_clone
    Folder1_clone
        File1_clone
    File2_clone
    File3_clone
*/

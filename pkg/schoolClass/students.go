package schoolClass

import "errors"

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
}

var studentID int

var Names = []string{
	"John",
	"Jack",
	"James",
	"Jill",
	"Jane",
	"Jenny",
	"Jenifer",
	"Jen",
	"Bob",
	"Bill",
	"Ben",
	"Barbara",
	"Barry",
}

var Surnames = []string{
	"Smith",
	"Johnson",
	"Williams",
	"Jones",
	"Brown",
	"Davis",
	"Miller",
	"Wilson",
	"Moore",
	"Taylor",
	"Anderson",
	"Thomas",
	"Jackson",
}
var Ages = []int{
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	25,
	26,
	27,
	28,
	29,
	30,
}

func GenerateStudents() []Student {
	var students []Student
	for i := range Ages {
		student := CreateStudent(Names[i], Surnames[i], Names[i]+"."+Surnames[i]+"@gmail.com", Ages[i])
		student.Id = i
		students = append(students, *student)
		studentID++
	}
	return students
}

func CreateStudent(name string, surname string, email string, age int) *Student {
	return &Student{Name: name, Surname: surname, Email: email, Age: age}
}

func GetStudentByID(id int) (*Student, error) {
	for _, student := range GenerateStudents() {
		if student.Id == id {
			return &student, nil
		}
	}
	return nil, errors.New("student not found")
}

package main

import "fmt"

type student struct {
	name            string
	subjectAndScore map[string]float64
}

func (stud student) Calculate_ave() float64 {
	score := 0.0
	average := 0.0
	for _, value := range stud.subjectAndScore {
		score += value
	}
	if len(stud.subjectAndScore) == 0 {
		average = 0.0
	} else {
		average = score / float64(len(stud.subjectAndScore))
	}
	return average
}
func Accept() student {
	fmt.Println("Whats your name?")
	var stud_name string
	var subject_no int
	fmt.Scan(&stud_name)
	person := student{name: stud_name, subjectAndScore: make(map[string]float64)}
	fmt.Printf("How many subjects do you take %v? ", stud_name)
	fmt.Scan(&subject_no)
	if subject_no <= 0 {
		fmt.Printf("That means you are not taking any course!\n")
	} else {
		fmt.Printf("Enter the subject(s) you take and the corresponding score you got per 100 (separated by space)\n")
		for i := 0; i < subject_no; i++ {
			var key string
			var value float64
			for {
				fmt.Scan(&key, &value)
				if value >= 0 || value <= 100 {
					break
				}
				fmt.Println("Incorrect input. correct it please!")

			}
			person.subjectAndScore[key] = value
		}
	}
	return person
}
func main() {
	p := Accept()
	av := p.Calculate_ave()
	fmt.Printf("Student name: %s\n", p.name)
	fmt.Println("Score per subject:")
	if len(p.subjectAndScore) == 0 {
		fmt.Println("None.")
	} else {
		for sub, score := range p.subjectAndScore {
			fmt.Printf(" %s: %v\n", sub, score)
		}
	}
	fmt.Printf("Average: %v ", av)
}

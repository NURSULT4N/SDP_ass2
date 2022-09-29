package main

import "fmt"

func removeAnElement(slice []Observer, val Observer) []Observer {
	for i, str := range slice {
		if str == val {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}

func removeAnElementStringType(slice []string, val string) []string {
	for i, str := range slice {
		if str == val {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}

type Observable interface {
	subscribe(Observable)
	unsubscribe(Observable)
	sendAll()
}

type Observer interface {
	handleEvent([]string)
}

func subscribe(jobSite *JobSite, observer Observer) {
	jobSite.subscribers = append(jobSite.subscribers, observer)
}

func unsubscribe(jobSite *JobSite, observer Observer) {
	jobSite.subscribers = removeAnElement(jobSite.subscribers, observer)
}

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

type Person struct {
	name string
}

func (p Person) handleEvent(vacancies []string) {
	fmt.Println()
	fmt.Print("Hello, dear ")
	fmt.Println(p.name)
	fmt.Println("Updated vacancies:")
	for _, v := range vacancies {
		fmt.Println(v)
	}
}

func sendAll(jobSite *JobSite) {
	for _, subscriber := range jobSite.subscribers {
		subscriber.handleEvent(jobSite.vacancies)
	}
}

func addVacancy(jobSite *JobSite, vacancy string) {
	jobSite.vacancies = append(jobSite.vacancies, vacancy)
	sendAll(jobSite)

}

func removeVacancy(jobSite *JobSite, vacancy string) {
	jobSite.vacancies = removeAnElementStringType(jobSite.vacancies, vacancy)
	sendAll(jobSite)
}

func main() {
	var jobSite JobSite
	//First Person
	firstPerson := Person{name: "Robert"}
	subscribe(&jobSite, firstPerson)
	//Job 1
	addVacancy(&jobSite, "*  FrontEnd Developer")

	//Second Person
	secondPerson := Person{name: "Martin"}
	subscribe(&jobSite, secondPerson)

	//Job 2
	addVacancy(&jobSite, "*  BackEnd Developer")

	//Unsubscribe and add Vacancy
	unsubscribe(&jobSite, firstPerson)
	addVacancy(&jobSite, "*  Full Stack Developer")

	//Remove Vacancy
	removeVacancy(&jobSite, "*  FrontEnd Developer")
}

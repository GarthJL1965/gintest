package main

import (
	"github.com/gjlanc65/gintest/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run("localhost:8080")
}

/*
POST data for postman

    {
        "id": 3,
        "LocalDate": "2022-09-11T17:10:00+10:00",
        "name": "GJL",
        "details": "Arrive Rome",
        "who": "DML&GJL"
    }

	<segment>
		<ID>1</ID>
		<LocalDate>2022-09-10T00:00:00+10:00</LocalDate>
		<Name>GJL</Name>
		<Details>Leave Sydney</Details>
		<Who>DML&amp;GJL</Who>
	</segment>

*/

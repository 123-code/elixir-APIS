package graphql
import (
	"github.com/graphql-go/graphql"
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
 )


 // entry point to our graphql server
 var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		"CrearPaciente":&graphql.Field{
			Type: PacienteType,
			"Nombre":&graphql.Field{
				Type: graphql.String,
			},
			"Apellido":&graphql.Field{
				Type: graphql.String,
			},
	
			"VPago" : &graphql.Field{
				Type:graphql.Int,
			},
	
			"VSemanales": &graphql.Field{
				Type:graphql.Int,
			},
	
			"Falto": &graphql.Field{
				Type:graphql.Boolean,
			},
		},
		},

 })

 var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:"rootMutation",
	Fields:graphql.Fields{

	},
 })
  

func main(){
	TypeFile, err := ioutil.ReadFile("./graphqltypes/type1.graphql")
	if err != nil {
		panic(err)
	}
   TypeString := string(TypeFile);
   
   
   var schema, err := graphql.ParseSchema(
    graphql.SchemaConfig{
        Query: rootQuery,
    },
    TypeString,
)
if err != nil {
    panic(err)
}
}
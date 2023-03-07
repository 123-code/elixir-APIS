package graph
import
(
	//"GraphQL/DB"
"GraphQLImp/graph/model"
"GraphQLImp/graph"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
pacientes []*model.Paciente
mutationResolver * graph.mutationResolver

}
package graph

import (
    "context"
    "gql/db"
    "gql/graph/model"
)

var dbConn, _ = db.Connect()

// CreatePaciente is the resolver for the createPaciente field.
func (r *mutationResolver) CreatePaciente(ctx context.Context, input model.PacienteInput) (*model.Paciente, error) {
    var paciente model.Paciente
    dbConn.Create(&paciente)
    return &paciente, nil
}

// UpdatePaciente is the resolver for the UpdatePaciente field.
func (r *mutationResolver) UpdatePaciente(ctx context.Context, id string, input model.UpdatePaciente) (*model.Paciente, error) {
    var paciente model.Paciente
    dbConn.First(&paciente, id)
    dbConn.Model(&paciente).Updates(&input)
    return &paciente, nil
}

// Pacientes is the resolver for the pacientes field.
func (r *queryResolver) Pacientes(ctx context.Context) ([]*model.Paciente, error) {
    var pacientes []*model.Paciente
    dbConn.Find(&pacientes)
    return pacientes, nil
}

// Paciente is the resolver for the paciente field.
func (r *queryResolver) Paciente(ctx context.Context, id string) (*model.Paciente, error) {
    var paciente model.Paciente
    dbConn.First(&paciente, id)
    return &paciente, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
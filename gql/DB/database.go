package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"gql/graph/model"
)

var dbConn *gorm.DB

func Connect() (*gorm.DB, error) {
    dsn := "postgres://zyjnpbwa:6kbWQ3f0S43kDBw2A6vkvmUAGlB-zhg5@kashin.db.elephantsql.com/zyjnpbwa"
    conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failded to connect to database")
    }

    conn.AutoMigrate(&model.Paciente{})
    conn.AutoMigrate(&model.Terapista{})
    conn.AutoMigrate(&model.UpdatePaciente{})
    conn.AutoMigrate(&model.PacienteInput{})

    return conn, err
}
type DB struct{}

func (d *DB) GetPaciente(id string) *model.Paciente {
	var paciente model.Paciente
	return &paciente
}

func (d *DB) GetAllPacientes() []*model.Paciente {
	var pacientes []*model.Paciente
	return pacientes
}

func (d *DB) CreatePaciente(input model.PacienteInput) *model.Paciente {
	var paciente model.Paciente
	paciente.Nombre = input.Nombre
	paciente.Apellido = input.Apellido
	paciente.Vsemana = input.Vsemana
	paciente.Paga = input.Paga

	result := dbConn.Create(&paciente)
	if result.Error != nil {
		return nil
	}

	return &paciente
}

func (d *DB) UpdatePaciente(id string, input model.UpdatePaciente) *model.Paciente {
	var paciente model.Paciente
	return &paciente
}
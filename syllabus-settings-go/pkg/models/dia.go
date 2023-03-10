package models

import (
	"gorm.io/gorm"
)

type Dia struct {
	gorm.Model
	DiaId     string `gorm:"index:idx_dia_id,unique"`
	DiaNome   string `json:"dia_nome"`
	DiaNumero int16  `gorm:"unique" json:"dia_numero"`
}

type DiaRequestModel struct {
	DiaNome   string `json:"dia_nome" binding:"required,max=50"`
	DiaNumero int16  `json:"dia_numero" binding:"required,gte=2,lte=7"`
}

type DiaResponseModel struct {
	DiaId     string `json:"dia_id"`
	DiaNome   string `json:"dia_nome"`
	DiaNumero int16  `json:"dia_numero"`
}

func (Dia) TableName() string {
	return "tb_dia"
}

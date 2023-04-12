package bella

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Matakuliahbel struct {
	ID         		primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	NamaMatakuliah	string             `bson:"namamatakuliah" json:"namamatakuliah"`
	KodeMatakuliah 	string             `bson:"kodemtkuliah" json:"kodemtkuliah"`
	Dosen       	string             `bson:"dosen" json:"dosen"`
	SKS   			string             `bson:"sks" json:"sks"`
}

type Jadwalkuliahbel struct {
	ID         	primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Hari		string             `bson:"hari" json:"hari"`
	JamMulai 	string             `bson:"jammulai" json:"jammulai"`
	JamSelesai  string             `bson:"jamselesai" json:"jamselesai"`
	Ruang   	string             `bson:"ruang" json:"ruang"`
}

type Kelasbel struct {
	ID         		primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Ruang			string             `bson:"ruang" json:"ruang"`
	KapasitasMhs 	string             `bson:"kapasitasmhs" json:"kapasitasmhs"`
}

type Dosenbel struct {
	ID         	primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	NamaDosen	string             `bson:"namadosen" json:"namadosen"`
	KodeDosen 	string             `bson:"kodedosen" json:"kodedosen"`
	Matakuliah 	string             `bson:"matakuliah" json:"matakuliah"`
}

type Mahasiswabel struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	NamaMhs	string             `bson:"namamhs" json:"namamhs"`
	Kelas	string             `bson:"kelas" json:"kelas"`
	Prodi   string             `bson:"prodi" json:"prodi"`
}
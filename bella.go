package bella

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMataKuliah(namamatakuliah string, kodematakuliah string, dosen string, sks string) (InsertedID interface{}) {
	var matakuliah Matakuliahbel
	matakuliah.NamaMatakuliah = namamatakuliah
	matakuliah.KodeMatakuliah = kodematakuliah
	matakuliah.Dosen = dosen
	matakuliah.SKS = sks

	return InsertOneDoc("DatabaseTugas3", "matakuliah", matakuliah)
}

func InsertJadwalKuliah(hari string, jammulai string, jamselesai string, ruang string) (InsertedID interface{}) {
	var jadwalkuliah Jadwalkuliahbel
	jadwalkuliah.Hari = hari
	jadwalkuliah.JamMulai = jammulai
	jadwalkuliah.JamSelesai = jamselesai
	jadwalkuliah.Ruang = ruang

	return InsertOneDoc("DatabaseTugas3", "jadwalkuliah", jadwalkuliah)
}

func InsertKelas(ruang string, kapasitasmhs string) (InsertedID interface{}) {
	var kelas Kelasbel
	kelas.Ruang = ruang
	kelas.KapasitasMhs = kapasitasmhs
	
	return InsertOneDoc("DatabaseTugas3", "kelas", kelas)
}

func InsertDosen(namadosen string, kodedosen string, matakuliah string) (InsertedID interface{}) {
	var dosen Dosenbel
	dosen.NamaDosen = namadosen
	dosen.KodeDosen = kodedosen
	dosen.Matakuliah = matakuliah

	return InsertOneDoc("DatabaseTugas3", "dosen", dosen)
}

func InsertMahasiswa(namamhs string, kelas string, prodi string) (InsertedID interface{}) {
	var mahasiswa Mahasiswabel
	mahasiswa.NamaMhs = namamhs
	mahasiswa.Kelas = kelas
	mahasiswa.Prodi = prodi

	return InsertOneDoc("DatabaseTugas3", "mahasiswa", mahasiswa)
}

func GetDataMatakuliah(kodematakuliah string) (data []Matakuliahbel) {
	user := MongoConnect("DatabaseTugas3").Collection("matakuliah")
	filter := bson.M{"kodemtkuliah": kodematakuliah}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatamatakuliah :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataJadwalkuliah(hari string) (data []Jadwalkuliahbel) {
	user := MongoConnect("DatabaseTugas3").Collection("jadwalkuliah")
	filter := bson.M{"hari": hari}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatajadwalkuliah :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataKelas(ruang string) (data []Kelasbel) {
	user := MongoConnect("DatabaseTugas3").Collection("kelas")
	filter := bson.M{"ruang": ruang}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatakelas :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataDosen(namadosen string) (data []Dosenbel) {
	user := MongoConnect("DatabaseTugas3").Collection("dosen")
	filter := bson.M{"namadosen": namadosen}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatadosen :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataMahasiswa(namamhs string) (data []Mahasiswabel) {
	user := MongoConnect("DatabaseTugas3").Collection("mahasiswa")
	filter := bson.M{"namamhs": namamhs}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatamahasiswa :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
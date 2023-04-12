package bella

import (
	"fmt"
	"testing"
)


func TestInsertMataKuliah(t *testing.T) {
	namamatakuliah := "Literasi Manusia"
	kodematakuliah := "TI41278"
	dosen := "NOVIANA RIZA"
	sks := "2"
	hasil := InsertMataKuliah(namamatakuliah, kodematakuliah, dosen, sks)
	fmt.Println(hasil)
}

func TestInsertJadwalKuliah(t *testing.T) {
	hari := "Rabu"
	jammulai := "10:00"
	jamselesai := "12:00"
	ruang := "310"
	hasil := InsertJadwalKuliah(hari, jammulai, jamselesai, ruang)
	fmt.Println(hasil)
}

func TestInsertKelas(t *testing.T) {
	ruang := "310"
	kapasitasmhs := "25"
	hasil := InsertKelas(ruang, kapasitasmhs)
	fmt.Println(hasil)
}

func TestInsertDosen(t *testing.T) {
	namadosen := "NOVIANA RIZA"
	kodedosen := "TI41278"
	matakuliah := "Literasi Manusia"
	hasil := InsertDosen(namadosen, kodedosen, matakuliah)
	fmt.Println(hasil)
}

func TestInsertMahasiswa(t *testing.T) {
	namamhs := "ARDIVA PUTRI TAVA PRAMESWARI"
	kelas := "2A"
	prodi := "D4 Teknik Informatika"
	hasil := InsertMahasiswa(namamhs, kelas, prodi)
	fmt.Println(hasil)
}

func TestGetDataMatakuliah(t *testing.T) {
	kodemtkuliah := "TI41278"
	dt := GetDataMatakuliah(kodemtkuliah)
	fmt.Println(dt)
}

func TestGetDataJadwalkuliah(t *testing.T) {
	hari := "Rabu"
	dt := GetDataJadwalkuliah(hari)
	fmt.Println(dt)
}

func TestGetDataKelas(t *testing.T) {
	ruang := "310"
	dt := GetDataKelas(ruang)
	fmt.Println(dt)
}

func TestGetDataDosen(t *testing.T) {
	namadosen := "NOVIANA RIZA"
	dt := GetDataDosen(namadosen)
	fmt.Println(dt)
}

func TestGetDataMahasiswa(t *testing.T) {
	namamhs := "ARDIVA PUTRI TAVA PRAMESWARI"
	dt := GetDataMahasiswa(namamhs)
	fmt.Println(dt)
}
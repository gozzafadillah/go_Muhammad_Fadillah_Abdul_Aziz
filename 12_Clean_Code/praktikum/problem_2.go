package main

// Root struct
type Kendaraan struct {
	kecepatanPerJam int
}

// Sub struct
type Mobil struct {
	kendaraan
}

// Mobil Start
func (m *Mobil) Berjalan() {
	m.tambahKecepatan(10)
}

// Menambahkan kecepatan
func (m *Mobil) tambahKecepatan(kecepatanBaru int) {
	output := m.kecepatanPerJam + kecepatanBaru
	m.kecepatanPerJam = output

}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLamban := Mobil{}
	mobilLamban.Berjalan()
}

package entities

type transfer struct {
	id					string
	user_id_pengirim	string
	user_id_penerima	string
	jumlah_transfer		int
	total_saldo			int
}
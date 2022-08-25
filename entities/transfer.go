package entities

type Transfer struct {
	Id					string
	User_id_pengirim	string
	User_id_penerima	string
	Jumlah_transfer		int
}
type HistoryTransfer struct {
	NamaPengirim string
	NamaPenerima string
	Nominal      int
	SisaSaldo    int
}
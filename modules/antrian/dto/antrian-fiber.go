package dto

type (
	StatusAntrianRequestV2 struct {
		KodePoli       string `json:"kodepoli" validate:"required"`
		TanggalPeriksa string `json:"tanggalperiksa" validate:"required"`
		KodeDokter     int    `json:"kodedokter" validate:"required"`
		JamPraktek     string `json:"jampraktek" validate:"required"`
	}

	GetSisaAntrianRequestV2 struct {
		Kodebooking string `json:"kodebooking" validate:"required"`
	}

	BatalAntreanRequestV2 struct {
		Kodebooking string `json:"kodebooking" validate:"required"`
		Keterangan  string `json:"keterangan" validate:"required"`
	}

	//

	CheckInRequestV2 struct {
		Kodebooking string `json:"kodebooking" validate:"required"`
		Waktu       int64  `json:"waktu" validate:"required"`
	}

	JadwalOperasiRequestV2 struct {
		Tanggalawal  string `json:"tanggalawal" validate:"required"`
		Tanggalakhir string `json:"tanggalakhir" validate:"required"`
	}

	JadwalOperasiPasienRequestV2 struct {
		Nopeserta string `json:"nopeserta" validate:"required"`
	}

	GetAntrianRequestV2 struct {
		Nomorkartu     string `json:"nomorkartu"`
		Nik            string `json:"nik" validate:"required"`
		Nohp           string `json:"nohp" validate:"required"`
		Kodepoli       string `json:"kodepoli" validate:"required"`
		Norm           string `json:"norm,omitempty"`
		Tanggalperiksa string `json:"tanggalperiksa" validate:"required"`
		Kodedokter     int    `json:"kodedokter" validate:"required"`
		Jampraktek     string `json:"jampraktek" validate:"required"`
		Jeniskunjungan int    `json:"jeniskunjungan" validate:"required"`
		Nomorreferensi string `json:"nomorreferensi" validate:"required"`
	}
)

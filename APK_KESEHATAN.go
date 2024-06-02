package main

import "fmt"

const NMAX int = 500
const DMAX int = 100
const chat int = 1000

type datadiri struct {
	nama, pw      		string
	tgl, bln, thn, no 	int
}

type pesan struct {
	nama, pesan, tag1, tag2, tag3 string
}

type forumapk [chat]pesan
type daftarpasien [NMAX]datadiri
type daftardokter [DMAX]datadiri

func main() {
	var opsimenu, opsidaftar, opsilogin,opsiforum, npas, ndok int
	var pwdokter, isipwd string
	var pas daftarpasien
	var dok daftardokter
	var logstate bool
	fmt.Print("Kata sandi autentikasi untuk dokter: ")
	fmt.Scan(&pwdokter)
	npas = -1
	ndok = -1
	for opsimenu != 4 {
		menu()
		fmt.Scan(&opsimenu)
		if opsimenu == 1 {
			fmt.Println(">> Daftar Sebagai:")
			fmt.Println("1. Dokter")
			fmt.Println("2. Pasien")
			fmt.Scan(&opsidaftar)
			if opsidaftar == 1 {
				if ndok <= DMAX-1 && logstate == false {
					DafDok(&ndok, &dok, isipwd, pwdokter)
				} else if ndok == DMAX-1 && logstate == false{
					fmt.Println("Jumlah dokter penuh!")
				} else {
					fmt.Println("Anda sudah log in!")
				}
			} else if opsidaftar == 2 {
				if npas <= NMAX-1 && logstate == false {
					DafPas(&npas, &pas)
				} else if logstate == false {
					fmt.Println("Daftar pasien penuh!")
				} else {
					fmt.Println("Anda sudah log in!")
				}
			}
		} else if opsimenu == 2 {
		if logstate == false {
			fmt.Println("Log in sebagai:")
			fmt.Println("1> Dokter")
			fmt.Println("2> Pasien") 
			fmt.Scan(&opsilogin)
			if opsilogin == 1 {
				loginDok(dok,isipw,pwdokter,ndok,&logstate)
			} else if opsilogin == 2 {
				loginPas(pas,npas,&logstate)
			}
		} else {
			fmt.Println("Anda sudah log in!")
		}
		} else if opsimenu == 3 {
			for opsiforum != 3 {
				menuforum()
				fmt.Scan(&opsiforum)
			}
		}
	}

}

func menu() {
	fmt.Println("<=======================>")
	fmt.Println("	     {WELCOME}        ")
	fmt.Println()
	fmt.Println("		 1> Daftar		  ")
	fmt.Println("	     2> Log in		  ")
	fmt.Println("		 3> Forum         ")
	fmt.Println("	     4> Exit		  ")
	fmt.Println()
	fmt.Println("<=======================>")
}

func menuforum() {
	fmt.Println("Selamat datang di forum! Untuk berkomentar dan bertanya, mohon mendaftar sebagai pasien dan log in.")
	fmt.Println("		{FORUM}        ")
	fmt.Println()
	fmt.Println("	   1> Cari")
	fmt.Println("	   2> Lihat Forum")
	fmt.Println("	   3> Kembali")
}

func DafDok(n *int, D *daftardokter, Gpw, Mpw string) {
	*n = *n + 1
	fmt.Scan(&Gpw)
	if Gpw == Mpw {
		fmt.Println("Nama: ")
		fmt.Scan(&D[*n].nama)
		fmt.Println("Password: ")
		fmt.Scan(&D[*n].pw)
		fmt.Println("Tanggal lahir (D/M/Y)")
		fmt.Scan(&D[*n].tgl, &D[*n].bln, &D[*n].thn)
		D[*n].no = *n
	} else if Gpw != Mpw {
		fmt.Println("Sandi Admin Salah!")
		DafDok(&*n, &*D, Gpw, Mpw)
	}
}

 func DafPas(n *int, P *daftarpasien) {
	*n = *n + 1
	fmt.Println("Nama: ")
	fmt.Scan(&P[*n].nama)
	fmt.Println("Password: ")
	fmt.Scan(&P[*n].pw)
	fmt.Println("Tanggal lahir (D/M/Y)")
	fmt.Scan(&P[*n].tgl, &P[*n].bln, &P[*n].thn)
	P[*n].no = *n
 }

 func loginDok(D daftardokter, Gpw,Mpw string, n int, ceklogin *bool) {
	var name, upw string
	var i int=0
	fmt.Scan(&Gpw)
	if Gpw == Mpw {
		fmt.Print("Nama: ")
		fmt.Scan(&name)
		fmt.Println()
		fmt.Print("Password: ")
		fmt.Scan(&upw)
		for i <= n {
			if name == D[i].nama && upw == D[i].pw {
				fmt.Println("Log in berhasil! Halo, Dokter ",D[i].nama)
				*ceklogin = true
			} else {
				fmt.Println("Nama atau Password salah!")
			}
		}
	}
 }

 func loginPas(P daftarpasien, n int, ceklogin *bool) {
	var name, upw, string
	fmt.Print("Nama: ")
	fmt.Scan(&name)
	fmt.Println()
	fmt.Print("Password: ")
	fmt.Scan(&upw)
	for i:=0; i <= n; i++ {
		if name == P[i].nama && upw == P[i].pw {
			fmt.Println("Log in berhasil!")
			*ceklogin = true
		}
	}
	if *ceklogin == false {
		fmt.Println("Username atau password salah.")
	}
}
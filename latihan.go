// package main

// import (
// 	"database/sql"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	_ "github.com/lib/pq"
// )

// type Pegawai struct {
// 	ID            int       `json:"id"`
// 	Nip           int       `json:"nip"`
// 	Nama_lengkap  string    `json:"nama_lengkap"`
// 	Jabatan       string    `json:"jabatan"`
// 	Jenis_kelamin string    `json:"jenis_kelamin"`
// 	Golongan      string    `json:"golongan"`
// 	Pangkat       string    `json:"pangkat"`
// 	Unit_kerja    string    `json:"unit_kerja"`
// 	Skpd          string    `json:"skpd"`
// 	CreatedAt     time.Time `json:"created_at"`
// }

// func initDB() (*sql.DB, error) {

// 	dns := "postgres://postgres.wjspbuobeznnunxlcvlj:Neilson4654%3F@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres?sslmode=require"

// 	db, err := sql.Open("postgres", dns)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		db.Close()
// 		return nil, err
// 	}

// 	return db, nil
// }

// func main() {
// 	db, err := initDB()

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	app := fiber.New()

// 	app.Get("/api/pegawai", func(c *fiber.Ctx) error {
// 		rows, err := db.Query(`
// 			SELECT
// 				id,
// 				nip,
// 				nama_lengkap,
// 				jabatan,
// 				jenis_kelamin,
// 				golongan,
// 				pangkat,
// 				unit_kerja,
// 				skpd,
// 				created_at
// 			FROM pegawai`,
// 		)

// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"message": "Failed to Query Database" + err.Error(),
// 			})
// 		}

// 		defer rows.Close()

// 		var pegawais []Pegawai
// 		for rows.Next() {
// 			var pegawai Pegawai
// 			err := rows.Scan(
// 				&pegawai.ID,
// 				&pegawai.Nip,
// 				&pegawai.Nama_lengkap,
// 				&pegawai.Jabatan,
// 				&pegawai.Jenis_kelamin,
// 				&pegawai.Golongan,
// 				&pegawai.Pangkat,
// 				&pegawai.Unit_kerja,
// 				&pegawai.Skpd,
// 				&pegawai.CreatedAt,
// 			)

// 			if err != nil {
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 					"message": "Failed to Scan Rows" + err.Error(),
// 				})
// 			}

// 			pegawais = append(pegawais, pegawai)

// 		}

// 		return c.JSON(fiber.Map{
// 			"Data": pegawais,
// 		})

// 	})

// 	app.Post("/api/pegawai", func(c *fiber.Ctx) error {
// 		var pegawai Pegawai
// 		err := c.BodyParser(&pegawai)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"message": "Failed to Parse" + err.Error(),
// 			})
// 		}

// 		stmt, err := db.Prepare(`
// 			INSERT INTO pegawai (
// 			    nip,
//                 nama_lengkap,
//                 jabatan,
//                 jenis_kelamin,
//                 golongan,
// 				pangkat,
//                 unit_kerja,
//                 skpd,
//                 created_at
// 			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id
// 		`)

// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"message": "Failed to Prepare DB" + err.Error(),
// 			})
// 		}

// 		defer stmt.Close()

// 		err = stmt.QueryRow(
// 			pegawai.Nip,
// 			pegawai.Nama_lengkap,
// 			pegawai.Jabatan,
// 			pegawai.Jenis_kelamin,
// 			pegawai.Golongan,
// 			pegawai.Pangkat,
// 			pegawai.Unit_kerja,
// 			pegawai.Skpd,
// 			pegawai.CreatedAt,
// 		).Scan(&pegawai.ID)

// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"Message": "Failed to Query Database" + err.Error(),
// 			})
// 		}

// 		return c.JSON(fiber.Map{
// 			"Message": "Data Berhasil di Tambahkan",
// 		})

// 	})

// 	app.Put("/api/pegawai/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")

// 		var pegawai Pegawai
// 		err := c.BodyParser(&pegawai)

// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"Message": "Failed to Parse" + err.Error(),
// 			})
// 		}

// 		stmt, err := db.Prepare(`
// 			UPDATE pegawai SET
// 			    nip=$1,
//                 nama_lengkap=$2,
//                 jabatan=$3,
//                 jenis_kelamin=$4,
//                 golongan=$5,
// 				pangkat=$6,
// 				unit_kerja=$7,
// 				skpd=$8
// 			WHERE id=$9 RETURNING id
// 		`)

// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"Message": "Failed to Prepare" + err.Error(),
// 			})
// 		}

// 		defer stmt.Close()

// 		_, err = stmt.Exec(
// 			pegawai.Nip,
// 			pegawai.Nama_lengkap,
// 			pegawai.Jabatan,
// 			pegawai.Jenis_kelamin,
// 			pegawai.Golongan,
// 			pegawai.Pangkat,
// 			pegawai.Unit_kerja,
// 			pegawai.Skpd,
// 			id,
// 		)

// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"Message": "Failed to Update " + err.Error(),
// 			})
// 		}

// 		return c.JSON(fiber.Map{
// 			"Message": "Data Berhasil di Update",
// 		})

// 	})

// 	app.Delete("/api/pegawai/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")

// 		stmt, err := db.Prepare(`DELETE FROM pegawai WHERE id =$1`)

// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"Message": "Failed to Prepare" + err.Error(),
// 			})
// 		}

// 		defer stmt.Close()

// 		_, err = stmt.Exec(id)

// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"Message": "Failed to Delete Database" + err.Error(),
// 			})
// 		}

// 		return c.JSON(fiber.Map{
// 			"Message": "Data Berhasil di Hapus",
// 		})

// 	})

// 	app.Listen(":8081")
// }

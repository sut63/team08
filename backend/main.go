package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sut63/team08/controllers"
	"github.com/sut63/team08/ent"
	"github.com/sut63/team08/ent/bloodtype"
	"github.com/sut63/team08/ent/gender"
	"github.com/sut63/team08/ent/prefix"
	"github.com/sut63/team08/ent/roomtype"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// --------------------------------- kao ----------------------------------------
type Nurses struct {
	Nurse []Nurse
}

type Nurse struct {
	Name     string
	Email    string
	Password string
	Tel      string
}

type Patients struct {
	Patient []Patient
}

type Patient struct {
	Name      string
	Age       int
	Weight    float64
	Height    float64
	Gender    int
	Prefix    int
	BloodType int
}

// --------------------------------- gop ----------------------------------------
type Rooms struct {
	Room []Room
}

type Room struct {
	Name     string
	Building int
	Floor    int
	RoomType int
}

type Roomtypes struct {
	Roomtype []Roomtype
}

type Roomtype struct {
	Name     string
	Bathroom int
	Toilet   int
	AreaSize float64
	Etc      string
}

// --------------------------------- mild ----------------------------------------
type SchemeTypes struct {
	SchemeType []SchemeType
}

type SchemeType struct {
	SchemeTypeName string
}

type Funds struct {
	Fund []Fund
}

type Fund struct {
	FundName string
}

type Certificates struct {
	Certificate []Certificate
}

type Certificate struct {
	CertificateName string
}

type Medicals struct {
	Medical []Medical
}

type Medical struct {
	MedicalName     string
	MedicalEmail    string
	MedicalPassword string
	MedicalTel      string
}

// --------------------------------- non ----------------------------------------
type Doctors struct {
	Doctor []Doctor
}
type Doctor struct {
	DoctorName     string
	DoctorEmail    string
	DoctorPassword string
	DoctorTel      string
}

type Diseases struct {
	Disease []Disease
}
type Disease struct {
	DiseaseName string
}

type Departments struct {
	Department []Department
}
type Department struct {
	DepartmentName string
}

// --------------------------------- bua ----------------------------------------
//Drugs struct
type Drugs struct {
	Drug []Drug
}

//Drug struct
type Drug struct {
	Name string
}

// --------------------------------- Pmay ----------------------------------------
type Examinationrooms struct {
	Examinationroom []Examinationroom
}
type Examinationroom struct {
	ExaminationroomName string
}

type Tools struct {
	Tool []Tool
}
type Tool struct {
	ToolName string
}

type Operatives struct {
	Operative []Operative
}
type Operative struct {
	OperativeName string
}

type Operativerecords struct {
	Operativerecord []Operativerecord
}
type Operativerecord struct {
	OperativerecordTime string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	// ---------------------------------  kao controller ----------------------------------------
	controllers.NewBloodtypeController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewNurseController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewPrefixController(v1, client)
	controllers.NewRentController(v1, client)
	controllers.NewRoomController(v1, client)
	controllers.NewRoomtypeController(v1, client)
	// --------------------------------- mild controller ----------------------------------------
	controllers.NewSchemeTypeController(v1, client)
	controllers.NewFundController(v1, client)
	controllers.NewCertificateController(v1, client)
	controllers.NewMedicalController(v1, client)
	controllers.NewCoveredPersonController(v1, client)
	// --------------------------------- non controller ----------------------------------------
	controllers.NewDoctorController(v1, client)
	controllers.NewDiseaseController(v1, client)
	controllers.NewDepartmentController(v1, client)
	controllers.NewDiagnoseController(v1, client)
	// --------------------------------- bua controller ----------------------------------------
	controllers.NewDrugController(v1, client)
	controllers.NewPrescriptionController(v1, client)
	// --------------------------------- Pmay ----------------------------------------
	controllers.NewExaminationroomController(v1, client)
	controllers.NewToolController(v1, client)
	controllers.NewOperativeController(v1, client)
	controllers.NewOperativerecordController(v1, client)
	// --------------------------------- kao Set Data ----------------------------------------
	// Set Nuses Data
	nurses := Nurses{
		Nurse: []Nurse{
			Nurse{"น.ส. สมร เย็นตา", "samon_123@hotmail.com", "samon123456", "097-2345678"},
			Nurse{"นาง ตาหวาน ตากลม", "tawan_eiei@gmail.com", "12345678", "067-7893452"},
		},
	}

	for _, n := range nurses.Nurse {
		client.Nurse.
			Create().
			SetNurseName(n.Name).
			SetNurseEmail(n.Email).
			SetNursePassword(n.Password).
			SetNurseTel(n.Tel).
			Save(context.Background())
	}

	// Set RoomTypes Data
	roomtypes := Roomtypes{
		Roomtype: []Roomtype{
			Roomtype{"Premier Royal Suite", 2, 1, 146.0, "Luxury interiors, Separate living room and patient room"},
			Roomtype{"VIP Suite", 1, 1, 71.0, "Luxury interiors, Separate living room and patient room"},
			Roomtype{"Single Room", 1, 0, 34.0, "Bathroom with shower"},
		},
	}

	for _, rt := range roomtypes.Roomtype {
		client.Roomtype.
			Create().
			SetName(rt.Name).
			SetBathroom(rt.Bathroom).
			SetToilets(rt.Toilet).
			SetAreasize(rt.AreaSize).
			SetEtc(rt.Etc).
			Save(context.Background())
	}

	// Set Rooms Data
	rooms := Rooms{
		Room: []Room{
			Room{"A111", 1, 1, 3},
			Room{"A112", 1, 1, 3},
			Room{"A113", 1, 1, 3},
			Room{"A114", 1, 1, 3},
			Room{"A115", 1, 1, 3},
			Room{"A116", 1, 1, 3},

			Room{"B121", 1, 2, 2},
			Room{"B122", 1, 2, 2},
			Room{"B123", 1, 2, 2},

			Room{"C131", 1, 3, 1},

			Room{"A211", 2, 1, 3},
			Room{"A212", 2, 1, 3},
			Room{"A213", 2, 1, 3},
			Room{"A214", 2, 1, 3},
			Room{"A215", 2, 1, 3},
			Room{"A216", 1, 1, 3},

			Room{"B221", 2, 2, 2},
			Room{"B222", 2, 2, 2},
			Room{"B223", 2, 2, 2},

			Room{"C231", 2, 3, 1},
		},
	}

	for _, r := range rooms.Room {

		rt, err := client.Roomtype.
			Query().
			Where(roomtype.IDEQ(int(r.RoomType))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Room.
			Create().
			SetName(r.Name).
			SetBuilding(r.Building).
			SetFloor(r.Floor).
			SetRoomtype(rt).
			Save(context.Background())
	}

	// Set Bloodtypes Data
	Bloodtypes := []string{"O", "A", "B", "AB"}
	for _, bt := range Bloodtypes {
		client.Bloodtype.
			Create().
			SetBTname(bt).
			Save(context.Background())
	}

	// Set Genders Data
	genders := []string{"Male", "Female"}
	for _, g := range genders {
		client.Gender.
			Create().
			SetGname(g).
			Save(context.Background())
	}

	// Set Prefixs Data
	prefixs := []string{"Mr.", "Mrs.", "Miss"}
	for _, pf := range prefixs {
		client.Prefix.
			Create().
			SetPname(pf).
			Save(context.Background())
	}

	// Set Patients Data
	patients := Patients{
		Patient: []Patient{
			Patient{"Dhetporn Krongsin", 21, 70, 183, 1, 1, 1},
			Patient{"Ploymanee Chuyat", 22, 51, 170, 2, 2, 3},
		},
	}

	for _, p := range patients.Patient {

		bt, err := client.Bloodtype.
			Query().
			Where(bloodtype.IDEQ(int(p.BloodType))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		g, err := client.Gender.
			Query().
			Where(gender.IDEQ(int(p.Gender))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		pe, err := client.Prefix.
			Query().
			Where(prefix.IDEQ(int(p.Prefix))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Patient.
			Create().
			SetPatientName(p.Name).
			SetPatientAge(p.Age).
			SetPatientWeight(p.Weight).
			SetPatientHeight(p.Height).
			SetGender(g).
			SetPrefix(pe).
			SetBloodtype(bt).
			Save(context.Background())
	}
	// --------------------------------- mild Set Data ----------------------------------------
	// Set SchemeType Data
	schemeTypes := SchemeTypes{
		SchemeType: []SchemeType{
			SchemeType{"สิทธิการรักษาพยาบาลของข้าราชกาล"},
			SchemeType{"สิทธิประกันสังคม"},
			SchemeType{"สิทธิหลักประกันสุขภาพ"},
		},
	}
	for _, st := range schemeTypes.SchemeType {
		client.SchemeType.
			Create().
			SetSchemeTypeName(st.SchemeTypeName).
			Save(context.Background())
	}
	// Set Fund Data
	funds := Funds{
		Fund: []Fund{
			Fund{"ด้านการรักษาพยาบาล"},
			Fund{"ด้านการส่งเสริมสุขภาพ และป้องกันโรค"},
			Fund{"ด้านการฟื้นฟูสุขภาพ / แพทย์แผนไทย"},
		},
	}
	for _, f := range funds.Fund {
		client.Fund.
			Create().
			SetFundName(f.FundName).
			Save(context.Background())
	}

	// Set Certificate Data
	certificates := Certificates{
		Certificate: []Certificate{
			Certificate{"รับใบรับรองแพทย์"},
			Certificate{"ไม่รับใบรับรองแพทย์บ"},
		},
	}
	for _, ce := range certificates.Certificate {
		client.Certificate.
			Create().
			SetCertificateName(ce.CertificateName).
			Save(context.Background())
	}
	// Set Medicals Data
	medicals := Medicals{
		Medical: []Medical{
			Medical{"สมศรี ภาคภูมิ", "somsee@hotmail.com", "123", "0989987653"},
			Medical{"มีใจ มะมี", "mejai@hotmail.com", "123", "0897563456"},
			Medical{"ถังแก๊ส เดินได้", "tuggas@hotmail.com", "123", "0893456324"},
		},
	}
	for _, mm := range medicals.Medical {
		client.Medical.
			Create().
			SetMedicalName(mm.MedicalName).
			SetMedicalEmail(mm.MedicalEmail).
			SetMedicalPassword(mm.MedicalPassword).
			SetMedicalTel(mm.MedicalTel).
			Save(context.Background())
	}

	// --------------------------------- non Set Data ----------------------------------------
	// Set Doctor Data
	doctors := Doctors{
		Doctor: []Doctor{
			Doctor{"Dr. James", "James@gmail.com", "James", "0889977665"},
			Doctor{"Dr. John", "John@gmail.com", "John", "0899177998"},
			Doctor{"Dr. Susan", "Susan@gmail.com", "Susan", "0812345671"},
		},
	}
	for _, d := range doctors.Doctor {
		client.Doctor.
			Create().
			SetDoctorName(d.DoctorName).
			SetDoctorEmail(d.DoctorEmail).
			SetDoctorPassword(d.DoctorPassword).
			SetDoctorTel(d.DoctorTel).
			Save(context.Background())
	}
	// Set Disease Data
	diseases := Diseases{
		Disease: []Disease{
			Disease{"แก้วหูอักเสบ"},
			Disease{"อัมพฤกษ์"},
			Disease{"โรคต้อหิน"},
			Disease{"โรคไส้ติ่งอักเสบ"},
			Disease{"โรคเส้นเลือดขอด"},
			Disease{"โรคทางเดินหายใจในเด็ก"},
		},
	}
	for _, di := range diseases.Disease {
		client.Disease.
			Create().
			SetDiseaseName(di.DiseaseName).
			Save(context.Background())
	}

	// Set Department Data
	departments := Departments{
		Department: []Department{
			Department{"แผนกกุมารเวชกรรม"},
			Department{"แผนกเวชศาสตร์"},
			Department{"แผนกจักษุ"},
			Department{"แผนกหู คอ จมูก"},
			Department{"แผนกศัลยกรรม"},
			Department{"แผนกจิตเวช"},
		},
	}
	for _, dep := range departments.Department {
		client.Department.
			Create().
			SetDepartmentName(dep.DepartmentName).
			Save(context.Background())
	}
	// --------------------------------- bua Set Data ----------------------------------------
	// Set Drugs Data
	drugs := Drugs{
		Drug: []Drug{
			Drug{"Paracetamol"},
			Drug{"ยาธาตุน้ำขาว"},
			Drug{"Aspirin"},
			Drug{"ยาธาตุน้ำแดง"},
			Drug{"ยาแก้ไอน้ำดำ"},
			Drug{"ENO"},
			Drug{"Decolgen"},
		},
	}

	for _, dr := range drugs.Drug {

		client.Drug.
			Create().
			SetDrugName(dr.Name).
			Save(context.Background())
	}
	// --------------------------------- Pmay Set Data ----------------------------------------
	// Set Tool Data
	tools := Tools{
		Tool: []Tool{
			Tool{"เข็ม"},
			Tool{"ถุงน้ำเกลือ"},
			Tool{"เครื่องมือทำคลอด"},
		},
	}
	for _, t := range tools.Tool {
		client.Tool.
			Create().
			SetToolName(t.ToolName).
			Save(context.Background())
	}

	// Set Operative Data
	operatives := Operatives{
		Operative: []Operative{
			Operative{"ผ่าตัด"},
			Operative{"เย็บแผล"},
			Operative{"ให้น้ำเกลือ"},
		},
	}
	for _, o := range operatives.Operative {
		client.Operative.
			Create().
			SetOperativeName(o.OperativeName).
			Save(context.Background())
	}

	// Set Examinationroom Data
	examinationrooms := Examinationrooms{
		Examinationroom: []Examinationroom{
			Examinationroom{"ห้องผ่าตัด"},
			Examinationroom{"ห้องเย็บแผล"},
			Examinationroom{"ห้องให้น้ำเกลือ"},
		},
	}
	for _, e := range examinationrooms.Examinationroom {
		client.Examinationroom.
			Create().
			SetExaminationroomName(e.ExaminationroomName).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}

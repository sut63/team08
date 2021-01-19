// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/sut63/team08/ent/bloodtype"
	"github.com/sut63/team08/ent/certificate"
	"github.com/sut63/team08/ent/department"
	"github.com/sut63/team08/ent/disease"
	"github.com/sut63/team08/ent/doctor"
	"github.com/sut63/team08/ent/drug"
	"github.com/sut63/team08/ent/examinationroom"
	"github.com/sut63/team08/ent/fund"
	"github.com/sut63/team08/ent/gender"
	"github.com/sut63/team08/ent/medical"
	"github.com/sut63/team08/ent/nurse"
	"github.com/sut63/team08/ent/operative"
	"github.com/sut63/team08/ent/operativerecord"
	"github.com/sut63/team08/ent/patient"
	"github.com/sut63/team08/ent/prefix"
	"github.com/sut63/team08/ent/prescription"
	"github.com/sut63/team08/ent/rent"
	"github.com/sut63/team08/ent/room"
	"github.com/sut63/team08/ent/roomtype"
	"github.com/sut63/team08/ent/schema"
	"github.com/sut63/team08/ent/schemetype"
	"github.com/sut63/team08/ent/tool"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	bloodtypeFields := schema.Bloodtype{}.Fields()
	_ = bloodtypeFields
	// bloodtypeDescBTname is the schema descriptor for BTname field.
	bloodtypeDescBTname := bloodtypeFields[0].Descriptor()
	// bloodtype.BTnameValidator is a validator for the "BTname" field. It is called by the builders before save.
	bloodtype.BTnameValidator = bloodtypeDescBTname.Validators[0].(func(string) error)
	certificateFields := schema.Certificate{}.Fields()
	_ = certificateFields
	// certificateDescCertificateName is the schema descriptor for Certificate_Name field.
	certificateDescCertificateName := certificateFields[0].Descriptor()
	// certificate.CertificateNameValidator is a validator for the "Certificate_Name" field. It is called by the builders before save.
	certificate.CertificateNameValidator = certificateDescCertificateName.Validators[0].(func(string) error)
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescDepartmentName is the schema descriptor for Department_Name field.
	departmentDescDepartmentName := departmentFields[0].Descriptor()
	// department.DepartmentNameValidator is a validator for the "Department_Name" field. It is called by the builders before save.
	department.DepartmentNameValidator = departmentDescDepartmentName.Validators[0].(func(string) error)
	diseaseFields := schema.Disease{}.Fields()
	_ = diseaseFields
	// diseaseDescDiseaseName is the schema descriptor for Disease_Name field.
	diseaseDescDiseaseName := diseaseFields[0].Descriptor()
	// disease.DiseaseNameValidator is a validator for the "Disease_Name" field. It is called by the builders before save.
	disease.DiseaseNameValidator = diseaseDescDiseaseName.Validators[0].(func(string) error)
	doctorFields := schema.Doctor{}.Fields()
	_ = doctorFields
	// doctorDescDoctorName is the schema descriptor for Doctor_Name field.
	doctorDescDoctorName := doctorFields[0].Descriptor()
	// doctor.DoctorNameValidator is a validator for the "Doctor_Name" field. It is called by the builders before save.
	doctor.DoctorNameValidator = doctorDescDoctorName.Validators[0].(func(string) error)
	// doctorDescDoctorPassword is the schema descriptor for Doctor_Password field.
	doctorDescDoctorPassword := doctorFields[1].Descriptor()
	// doctor.DoctorPasswordValidator is a validator for the "Doctor_Password" field. It is called by the builders before save.
	doctor.DoctorPasswordValidator = doctorDescDoctorPassword.Validators[0].(func(string) error)
	// doctorDescDoctorEmail is the schema descriptor for Doctor_Email field.
	doctorDescDoctorEmail := doctorFields[2].Descriptor()
	// doctor.DoctorEmailValidator is a validator for the "Doctor_Email" field. It is called by the builders before save.
	doctor.DoctorEmailValidator = doctorDescDoctorEmail.Validators[0].(func(string) error)
	// doctorDescDoctorTel is the schema descriptor for Doctor_tel field.
	doctorDescDoctorTel := doctorFields[3].Descriptor()
	// doctor.DoctorTelValidator is a validator for the "Doctor_tel" field. It is called by the builders before save.
	doctor.DoctorTelValidator = doctorDescDoctorTel.Validators[0].(func(string) error)
	drugFields := schema.Drug{}.Fields()
	_ = drugFields
	// drugDescDrugName is the schema descriptor for Drug_Name field.
	drugDescDrugName := drugFields[0].Descriptor()
	// drug.DrugNameValidator is a validator for the "Drug_Name" field. It is called by the builders before save.
	drug.DrugNameValidator = drugDescDrugName.Validators[0].(func(string) error)
	examinationroomFields := schema.Examinationroom{}.Fields()
	_ = examinationroomFields
	// examinationroomDescExaminationroomName is the schema descriptor for examinationroom_name field.
	examinationroomDescExaminationroomName := examinationroomFields[0].Descriptor()
	// examinationroom.ExaminationroomNameValidator is a validator for the "examinationroom_name" field. It is called by the builders before save.
	examinationroom.ExaminationroomNameValidator = examinationroomDescExaminationroomName.Validators[0].(func(string) error)
	fundFields := schema.Fund{}.Fields()
	_ = fundFields
	// fundDescFundName is the schema descriptor for Fund_Name field.
	fundDescFundName := fundFields[0].Descriptor()
	// fund.FundNameValidator is a validator for the "Fund_Name" field. It is called by the builders before save.
	fund.FundNameValidator = fundDescFundName.Validators[0].(func(string) error)
	genderFields := schema.Gender{}.Fields()
	_ = genderFields
	// genderDescGname is the schema descriptor for Gname field.
	genderDescGname := genderFields[0].Descriptor()
	// gender.GnameValidator is a validator for the "Gname" field. It is called by the builders before save.
	gender.GnameValidator = genderDescGname.Validators[0].(func(string) error)
	medicalFields := schema.Medical{}.Fields()
	_ = medicalFields
	// medicalDescMedicalName is the schema descriptor for Medical_Name field.
	medicalDescMedicalName := medicalFields[0].Descriptor()
	// medical.MedicalNameValidator is a validator for the "Medical_Name" field. It is called by the builders before save.
	medical.MedicalNameValidator = medicalDescMedicalName.Validators[0].(func(string) error)
	// medicalDescMedicalEmail is the schema descriptor for Medical_Email field.
	medicalDescMedicalEmail := medicalFields[1].Descriptor()
	// medical.MedicalEmailValidator is a validator for the "Medical_Email" field. It is called by the builders before save.
	medical.MedicalEmailValidator = medicalDescMedicalEmail.Validators[0].(func(string) error)
	// medicalDescMedicalPassword is the schema descriptor for Medical_Password field.
	medicalDescMedicalPassword := medicalFields[2].Descriptor()
	// medical.MedicalPasswordValidator is a validator for the "Medical_Password" field. It is called by the builders before save.
	medical.MedicalPasswordValidator = medicalDescMedicalPassword.Validators[0].(func(string) error)
	// medicalDescMedicalTel is the schema descriptor for Medical_Tel field.
	medicalDescMedicalTel := medicalFields[3].Descriptor()
	// medical.MedicalTelValidator is a validator for the "Medical_Tel" field. It is called by the builders before save.
	medical.MedicalTelValidator = medicalDescMedicalTel.Validators[0].(func(string) error)
	nurseFields := schema.Nurse{}.Fields()
	_ = nurseFields
	// nurseDescNurseName is the schema descriptor for nurse_name field.
	nurseDescNurseName := nurseFields[0].Descriptor()
	// nurse.NurseNameValidator is a validator for the "nurse_name" field. It is called by the builders before save.
	nurse.NurseNameValidator = nurseDescNurseName.Validators[0].(func(string) error)
	// nurseDescNurseEmail is the schema descriptor for nurse_email field.
	nurseDescNurseEmail := nurseFields[1].Descriptor()
	// nurse.NurseEmailValidator is a validator for the "nurse_email" field. It is called by the builders before save.
	nurse.NurseEmailValidator = nurseDescNurseEmail.Validators[0].(func(string) error)
	// nurseDescNursePassword is the schema descriptor for nurse_password field.
	nurseDescNursePassword := nurseFields[2].Descriptor()
	// nurse.NursePasswordValidator is a validator for the "nurse_password" field. It is called by the builders before save.
	nurse.NursePasswordValidator = nurseDescNursePassword.Validators[0].(func(string) error)
	// nurseDescNurseTel is the schema descriptor for nurse_tel field.
	nurseDescNurseTel := nurseFields[3].Descriptor()
	// nurse.NurseTelValidator is a validator for the "nurse_tel" field. It is called by the builders before save.
	nurse.NurseTelValidator = nurseDescNurseTel.Validators[0].(func(string) error)
	operativeFields := schema.Operative{}.Fields()
	_ = operativeFields
	// operativeDescOperativeName is the schema descriptor for operative_Name field.
	operativeDescOperativeName := operativeFields[0].Descriptor()
	// operative.OperativeNameValidator is a validator for the "operative_Name" field. It is called by the builders before save.
	operative.OperativeNameValidator = operativeDescOperativeName.Validators[0].(func(string) error)
	operativerecordFields := schema.Operativerecord{}.Fields()
	_ = operativerecordFields
	// operativerecordDescOperativeTime is the schema descriptor for OperativeTime field.
	operativerecordDescOperativeTime := operativerecordFields[0].Descriptor()
	// operativerecord.DefaultOperativeTime holds the default value on creation for the OperativeTime field.
	operativerecord.DefaultOperativeTime = operativerecordDescOperativeTime.Default.(func() time.Time)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescPatientName is the schema descriptor for Patient_name field.
	patientDescPatientName := patientFields[0].Descriptor()
	// patient.PatientNameValidator is a validator for the "Patient_name" field. It is called by the builders before save.
	patient.PatientNameValidator = patientDescPatientName.Validators[0].(func(string) error)
	// patientDescPatientAge is the schema descriptor for Patient_age field.
	patientDescPatientAge := patientFields[1].Descriptor()
	// patient.PatientAgeValidator is a validator for the "Patient_age" field. It is called by the builders before save.
	patient.PatientAgeValidator = patientDescPatientAge.Validators[0].(func(int) error)
	// patientDescPatientWeight is the schema descriptor for Patient_weight field.
	patientDescPatientWeight := patientFields[2].Descriptor()
	// patient.PatientWeightValidator is a validator for the "Patient_weight" field. It is called by the builders before save.
	patient.PatientWeightValidator = patientDescPatientWeight.Validators[0].(func(float64) error)
	// patientDescPatientHeight is the schema descriptor for Patient_height field.
	patientDescPatientHeight := patientFields[3].Descriptor()
	// patient.PatientHeightValidator is a validator for the "Patient_height" field. It is called by the builders before save.
	patient.PatientHeightValidator = patientDescPatientHeight.Validators[0].(func(float64) error)
	prefixFields := schema.Prefix{}.Fields()
	_ = prefixFields
	// prefixDescPname is the schema descriptor for Pname field.
	prefixDescPname := prefixFields[0].Descriptor()
	// prefix.PnameValidator is a validator for the "Pname" field. It is called by the builders before save.
	prefix.PnameValidator = prefixDescPname.Validators[0].(func(string) error)
	prescriptionFields := schema.Prescription{}.Fields()
	_ = prescriptionFields
	// prescriptionDescPrescripDateTime is the schema descriptor for Prescrip_DateTime field.
	prescriptionDescPrescripDateTime := prescriptionFields[1].Descriptor()
	// prescription.DefaultPrescripDateTime holds the default value on creation for the Prescrip_DateTime field.
	prescription.DefaultPrescripDateTime = prescriptionDescPrescripDateTime.Default.(func() time.Time)
	rentFields := schema.Rent{}.Fields()
	_ = rentFields
	// rentDescRentID is the schema descriptor for rent_id field.
	rentDescRentID := rentFields[0].Descriptor()
	// rent.RentIDValidator is a validator for the "rent_id" field. It is called by the builders before save.
	rent.RentIDValidator = rentDescRentID.Validators[0].(func(string) error)
	// rentDescKinTel is the schema descriptor for kin_tel field.
	rentDescKinTel := rentFields[1].Descriptor()
	// rent.KinTelValidator is a validator for the "kin_tel" field. It is called by the builders before save.
	rent.KinTelValidator = rentDescKinTel.Validators[0].(func(string) error)
	// rentDescKinName is the schema descriptor for kin_name field.
	rentDescKinName := rentFields[2].Descriptor()
	// rent.KinNameValidator is a validator for the "kin_name" field. It is called by the builders before save.
	rent.KinNameValidator = rentDescKinName.Validators[0].(func(string) error)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescName is the schema descriptor for name field.
	roomDescName := roomFields[0].Descriptor()
	// room.NameValidator is a validator for the "name" field. It is called by the builders before save.
	room.NameValidator = roomDescName.Validators[0].(func(string) error)
	// roomDescBuilding is the schema descriptor for building field.
	roomDescBuilding := roomFields[1].Descriptor()
	// room.BuildingValidator is a validator for the "building" field. It is called by the builders before save.
	room.BuildingValidator = roomDescBuilding.Validators[0].(func(int) error)
	// roomDescFloor is the schema descriptor for floor field.
	roomDescFloor := roomFields[2].Descriptor()
	// room.FloorValidator is a validator for the "floor" field. It is called by the builders before save.
	room.FloorValidator = roomDescFloor.Validators[0].(func(int) error)
	roomtypeFields := schema.Roomtype{}.Fields()
	_ = roomtypeFields
	// roomtypeDescName is the schema descriptor for name field.
	roomtypeDescName := roomtypeFields[0].Descriptor()
	// roomtype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	roomtype.NameValidator = roomtypeDescName.Validators[0].(func(string) error)
	// roomtypeDescAreasize is the schema descriptor for areasize field.
	roomtypeDescAreasize := roomtypeFields[3].Descriptor()
	// roomtype.AreasizeValidator is a validator for the "areasize" field. It is called by the builders before save.
	roomtype.AreasizeValidator = roomtypeDescAreasize.Validators[0].(func(float64) error)
	// roomtypeDescEtc is the schema descriptor for etc field.
	roomtypeDescEtc := roomtypeFields[4].Descriptor()
	// roomtype.EtcValidator is a validator for the "etc" field. It is called by the builders before save.
	roomtype.EtcValidator = roomtypeDescEtc.Validators[0].(func(string) error)
	schemetypeFields := schema.SchemeType{}.Fields()
	_ = schemetypeFields
	// schemetypeDescSchemeTypeName is the schema descriptor for SchemeType_Name field.
	schemetypeDescSchemeTypeName := schemetypeFields[0].Descriptor()
	// schemetype.SchemeTypeNameValidator is a validator for the "SchemeType_Name" field. It is called by the builders before save.
	schemetype.SchemeTypeNameValidator = schemetypeDescSchemeTypeName.Validators[0].(func(string) error)
	toolFields := schema.Tool{}.Fields()
	_ = toolFields
	// toolDescToolName is the schema descriptor for Tool_Name field.
	toolDescToolName := toolFields[0].Descriptor()
	// tool.ToolNameValidator is a validator for the "Tool_Name" field. It is called by the builders before save.
	tool.ToolNameValidator = toolDescToolName.Validators[0].(func(string) error)
}

import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page,pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2'; // alert
import NavigateBeforeIcon from '@material-ui/icons/NavigateBefore';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  Link,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntPatient } from '../../api/models/EntPatient'; // import interface Patient
import { EntDoctor } from '../../api/models/EntDoctor'; // import interface Doctor
import { EntDisease} from '../../api/models/EntDisease'; // import interface Disease
import { EntDepartment } from '../../api/models/EntDepartment'; // import interface Department

// name
import { Cookies } from 'react-cookie/cjs';//cookie
import { useCookies } from 'react-cookie/cjs';//cookie
const cookies = new Cookies();
const Name = cookies.get('Name');

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));


const Diagnose: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

 
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
  const [diseases, setDiseases] = React.useState<EntDisease[]>([]);
  const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
  

  const getPatient = async () => {
    const res = await http.listPatient({ limit: 2, offset: 0 });
    setPatients(res);
  };

  const getDoctor = async () => {
    const res = await http.listDoctor({ limit: 3, offset: 0 });
    setDoctors(res);
  };

  const getDisease = async () => {
    const res = await http.listDisease({ limit: 8, offset: 0 });
    setDiseases(res);
  };

  const getDepartment = async () => {
    const res = await http.listDepartment({ limit: 6, offset: 0 });
    setDepartments(res);
  };

    // Lifecycle Hooks
    useEffect(() => {
      getPatient();
      getDoctor();
      getDisease();
      getDepartment();
    }, []);

  const DiagnoseIDhandleChange = (event: any) => {
    setDiagnoseID(event.target.value as string);
  };
  const DiagnoseSymptomshandleChange = (event: any) => {
    setDiagnoseSymptoms(event.target.value as string);
  };
  const DiagnoseNotehandleChange = (event: any) => {
    setDiagnoseNote(event.target.value as string);
  };
  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientID(event.target.value as number);
  };
  const DoctorhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDoctorID(event.target.value as number);
  };
  const DiseasehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDiseaseID(event.target.value as number);
  };
  const DepartmenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDepartmentID(event.target.value as number);
  };

  const [patientID, setPatientID] = React.useState(Number);
  const [doctorID, setDoctorID] = React.useState(Number);
  const [diseaseID, setDiseaseID] = React.useState(Number);
  const [departmentID, setDepartmentID] = React.useState(Number);
  const [diagnoseid, setDiagnoseID] = React.useState(String);
  const [diagnosesymptoms, setDiagnoseSymptoms] = React.useState(String);
  const [diagnosenote, setDiagnoseNote] = React.useState(String);


  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });
  const diagnoses = {
    note: diagnosenote,
    disease: diseaseID,
    doctor: doctorID,
    patient: patientID,
    department: departmentID,
    symptoms: diagnosesymptoms,
    number: diagnoseid
  }
  // alert error
const aleartMessage = (icon: any, title: any) => {
  Toast.fire({
    icon: icon,
    title: title,
  });
}
// check error 
const checkCaseSaveError = (field: string) => {
  switch (field) {
    case 'Doctor not found':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาเลือกชื่อแพทย์");
      return;
    case 'Diagnose_ID':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณากรอก ID การวินิจฉัย ขึ้นต้นด้วย D และตัวเลข 4 ตัว");
      return;
    case 'Patient not found':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาป้อนชื่อผู้ป่วย");
      return;
    case 'Disease not found':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาเลือกโรคที่วินิจฉัย");
      return;
      case 'Diagnose_Symptoms':
        aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณากรอกอาการเพิ่มเติม หากไม่มีให้ใช้ -");
        return;
    case 'Department not found':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาเลือกแผนกที่เข้ารับการรักษา");
      return;
    case 'Diagnose_Note':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณากรอกหมายเหตุ หากไม่มีให้ใช้ -");
      return;
    default:
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ");
      return;
  };
}

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/diagnoses';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(diagnoses),
    };
    console.log(diagnoses); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
        
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name || data.error)
        }

      });
}

// name
function a11yProps(index: any) {
  return {
    id: `scrollable-force-tab-${index}`,
    'aria-controls': `scrollable-force-tabpanel-${index}`,
  };
}
const [cookies, setCookie, removeCookie] = useCookies(['cookiename']);

  function Logout() {
    removeCookie('ID', { path: '/' })
    removeCookie('Name', { path: '/' })
    removeCookie('Email', { path: '/' })
    window.location.href = "http://localhost:3000/";
  }

  return (
    <Page theme={pageTheme.service}>
     <Header style={HeaderCustom} title={`ระบบวินิจฉัยโรคผู้ป่วยใน`}>
     <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <div style={{ marginLeft: 1 }}>{Name}</div>
        <div style={{ marginLeft: 10 }}>
        <Link component={RouterLink} to="/">
             Logout
         </Link>
         </div>
      </Header>

      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>แพทย์ที่รักษา</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="doctor"
                  value={doctorID}
                  onChange={DoctorhandleChange}
                >
                  {doctors.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.doctorName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
            <div className={classes.paper}>รหัสการวินิจฉัย</div>
          </Grid>
          <Grid item xs={8}>
          <TextField
             name="diagnoseID"
             id="Diagnose_ID"
             label=""
             multiline
             defaultValue="Default Value"
             variant="outlined"
             onChange={DiagnoseIDhandleChange}
             value={diagnoseid} 
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
            />
          </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ชื่อผู้ป่วย</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="patient"
                  value={patientID}
                  onChange={PatienthandleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.patientName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>การวินิจฉัยโรค</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="disease"
                  value={diseaseID}
                  onChange={DiseasehandleChange}
                >
                  {diseases.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.diseaseName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
            <div className={classes.paper}>อาการเพิ่มเติม</div>
          </Grid>
          <Grid item xs={8}>
          <FormControl variant="outlined" className={classes.formControl}>
          <TextField
             name="symptoms"
             id="symptoms"
             variant="outlined"
             value={diagnosesymptoms} 
             onChange={DiagnoseSymptomshandleChange}
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
                
            />
            </FormControl>
          </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>แผนกการเข้ารักษา</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="department"
                  value={departmentID}
                  onChange={DepartmenthandleChange}
                >
                  {departments.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.departmentName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
            <div className={classes.paper}>หมายเหตุ</div>
          </Grid>
          <Grid item xs={8}>
          <TextField
             name="note"
             label=""
             multiline
             defaultValue="Default Value"
             variant="outlined"
             value={diagnosenote} 
             onChange={DiagnoseNotehandleChange}
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
                
            />
          </Grid>

            <Grid item xs={4}></Grid>
            <Grid item xs={8}>
            <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
                
              >
                SAVE
            </Button>
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            <Button
              variant="contained"
              size="large"
              startIcon={<NavigateBeforeIcon />}
              component={RouterLink} to="/homedoctor"
            >
              Back
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Diagnose; 
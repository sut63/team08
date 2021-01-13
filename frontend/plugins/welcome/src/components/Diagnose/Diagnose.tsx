import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page,pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Button,
  Link,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntPatient } from '../../api/models/EntPatient'; // import interface Patient
import { EntDoctor } from '../../api/models/EntDoctor'; // import interface Doctor
import { EntDisease} from '../../api/models/EntDisease'; // import interface Disease
import { EntDepartment } from '../../api/models/EntDepartment'; // import interface Department
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
    
  // alert setting
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

  const getPatient = async () => {
    const res = await http.listPatient({ limit: 8, offset: 0 });
    setPatients(res);
  };

  const getDoctor = async () => {
    const res = await http.listDoctor({ limit: 8, offset: 0 });
    setDoctors(res);
  };

  const getDisease = async () => {
    const res = await http.listDisease({ limit: 8, offset: 0 });
    setDiseases(res);
  };

  const getDepartment = async () => {
    const res = await http.listDepartment({ limit: 8, offset: 0 });
    setDepartments(res);
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



  // Lifecycle Hooks
  useEffect(() => {
    getPatient();
    getDoctor();
    getDisease();
    getDepartment();
  }, []);


  // function save data
  function save() {
    const diagnoses = {
      patient: patientID,
      doctor: doctorID,
      disease: diseaseID,
      department: departmentID,
    }
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
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }

      });
}
  
  return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`ระบบการวินิจฉัยโรคผู้ป่วยใน`}>
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <Link component={RouterLink} to="/">
             Logout
         </Link>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
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
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Diagnose;
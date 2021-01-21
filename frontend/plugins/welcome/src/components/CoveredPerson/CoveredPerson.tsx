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
import { EntSchemeType } from '../../api/models/EntSchemeType'; // import interface SchemeType
import { EntFund} from '../../api/models/EntFund'; // import interface Fund
import { EntCertificate } from '../../api/models/EntCertificate'; // import interface Certificate
import { EntMedical } from '../../api/models/EntMedical'; // import interface Medical

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


const CoveredPerson: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

 
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  const [schemeTypes, setSchemeTypes] = React.useState<EntSchemeType[]>([]);
  const [funds, setFunds] = React.useState<EntFund[]>([]);
  const [certificates, setCertificates] = React.useState<EntCertificate[]>([]);
  const [medicals, setMedicals] = React.useState<EntMedical[]>([]);
    


  const getPatient = async () => {
    const res = await http.listPatient({ limit: 2, offset: 0 });
    setPatients(res);
  };

  const getSchemeType = async () => {
    const res = await http.listSchemeType({ limit: 3, offset: 0 });
    setSchemeTypes(res);
  };

  const getFund = async () => {
    const res = await http.listFund({ limit: 3, offset: 0 });
    setFunds(res);
  };

  const getCertificate = async () => {
    const res = await http.listCertificate({ limit: 2, offset: 0 });
    setCertificates(res);
  };
  const getMedical = async () => {
    const res = await http.listMedical({ limit: 3, offset: 0 });
    setMedicals(res);
  };


  // Lifecycle Hooks
  useEffect(() => {
    getPatient();
    getSchemeType();
    getFund();
    getCertificate();
    getMedical();
  }, []);

  
  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientID(event.target.value as number);
  };
  const SchemeTypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSchemeTypeID(event.target.value as number);
  };
  const FundhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setFundID(event.target.value as number);
  };

  const CertificatehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setCertificateID(event.target.value as number);
  };
  const CoveredPersonNumberhandleChange = (event: any) => {
    setCoveredPersonNumber(event.target.value as string);
  };
  const CoveredPersonNotehandleChange = (event: any) => {
    setCoveredPersonNote(event.target.value as string);
  };
  const FundTitlehandleChange = (event: any) => {
    setFundTitle(event.target.value as string);
  };
  const MedicalhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setMedicalID(event.target.value as number);
  };

  const [patientID, setPatientID] = React.useState(Number);
  const [schemeTypeID, setSchemeTypeID] = React.useState(Number);
  const [fundID, setFundID] = React.useState(Number);
  const [certificateID, setCertificateID] = React.useState(Number);
  const [coveredPersonNumber, setCoveredPersonNumber] = React.useState(String);
  const [coveredPersonNote, setCoveredPersonNote] = React.useState(String);
  const [fundTitle, setFundTitle] = React.useState(String);
  const [medicalID, setMedicalID] = React.useState(Number);
 

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
    case 'CoveredPerson_Number':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: รหัสสิทธิการรักษาขึ้นต้นด้วย CP ตามด้วยตัวเลข 3 หลัก");
      return;
    case 'Patient not found':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาเลือกชื่อผู้ป่วย");
      return;
    case 'SchemeType not found':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาเลือกประเภทสิทธิรักษาพยาบาล");
      return;
    case 'Fund not found':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาเลือกกองทุน");
      return;
      case 'Fund_Title':
        aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณากรอกคำร้อง");
        return;
    case 'Certificate not found':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาเลือกใบรับรองแพทย์");
      return;
    case 'CoveredPerson_Note':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณากรอกหมายเหตุ ถ้าไม่มีใส่ ' - ' ");
      return;
    case 'Medical':
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาเลือก");
      return;  
    default:
      aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ");
      return;
  };
}

  // function save data
  function save() {


  const coveredpersons = {
    patient: patientID,
    fund: fundID,
    schemeType: schemeTypeID,
    certificate: certificateID,
    number: coveredPersonNumber,
    note: coveredPersonNote,
    fundTitle: fundTitle,
    medical: medicalID,
  }


    const apiUrl = 'http://localhost:8080/api/v1/coveredpersons';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(coveredpersons),
    };
    console.log(coveredpersons); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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
      <Header style={HeaderCustom} title={`ระบบสิทธิการรักษาพยาบาล`}>
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
            <div className={classes.paper}>เจ้าหน้าที่เวชระเบียน</div>
          </Grid>
          <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="medical"
                  value={medicalID}
                  onChange={MedicalhandleChange}
                >
                  {medicals.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.medicalName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
            <div className={classes.paper}>รหัสสิทธิการรักษา</div>
          </Grid>
          <Grid item xs={8}>
          <TextField
             name="number"
             label=""
             multiline
             defaultValue="Default Value"
             variant="outlined"
             value={coveredPersonNumber} 
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
                onChange={CoveredPersonNumberhandleChange}
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
              <div className={classes.paper}>ประเภทสิทธิการรักษาพยาบาล</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="schemeType"
                  value={schemeTypeID}
                  onChange={SchemeTypehandleChange}
                >
                  {schemeTypes.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.schemeTypeName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>กองทุน</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="fund"
                  value={fundID}
                  onChange={FundhandleChange}
                >
                  {funds.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.fundName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
            <div className={classes.paper}>คำร้องกองทุน</div>
          </Grid>
          <Grid item xs={8}>
          <TextField
             name="fundtitile"
             label=""
             multiline
             defaultValue="Default Value"
             variant="outlined"
             value={fundTitle} 
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
                onChange={FundTitlehandleChange}
            />
          </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ใบรับรองแพทย์</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="certificate"
                  value={certificateID}
                  onChange={CertificatehandleChange}
                >
                  {certificates.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.certificateName}
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
             name="fundtitile"
             label=""
             multiline
             defaultValue="Default Value"
             variant="outlined"
             value={coveredPersonNote} 
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
                onChange={CoveredPersonNotehandleChange}
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
              component={RouterLink} to="/homemedical"
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

export default CoveredPerson;
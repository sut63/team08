import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; 
import { Link as RouterLink } from 'react-router-dom';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2'; // alert
import NavigateBeforeIcon from '@material-ui/icons/NavigateBefore';
import {
  Container,
  Grid,
  FormControl,
  Select,
  MenuItem,
  TextField,
  Button,
  Link,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { 
  EntBloodtype,
  EntGender,
  EntPrefix,
  } from '../../api/models'; // import interface User

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


const Register: FC<{email:string}> = (email) => {
  
  
  
  
  const classes = useStyles();
  const http = new DefaultApi();

  

  const [bloodtypes, setBloodtypes] = React.useState<EntBloodtype[]>([]);
  const [prefixs, setPrefixs] = React.useState<EntPrefix[]>([]);
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  
  
  

  const getBloodtype = async () => {
    const res = await http.listBloodtype({ limit: 4, offset: 0 });
    setBloodtypes(res);
    console.log()
  };


  const getPrefix = async () => {
    const res = await http.listPrefix({ limit: 3, offset: 0 });
    setPrefixs(res);
  };

  const getGender = async () => {
    const res = await http.listGender({ limit: 2, offset: 0 });
    setGenders(res);
  };

  // Lifecycle Hooks
   // Lifecycle Hooks
   useEffect(() => {
    getBloodtype();
    getPrefix();
    getGender();
  }, []);

  // set data to object playlist_video
  
  
  const NamehandleChange = (event: any) => {
    setName(event.target.value as string);
  };
  
  const PrefixhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPredixID(event.target.value as number);
  };
  
  const GenderhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setGenderID(event.target.value as number);
  };
  
  const WeighthandleChange = (event: any) => {
    setWeight(event.target.value as string);
  };

  const HeighthandleChange = (event: any) => {
    setHeight(event.target.value as string);
  };

  const AgehandleChange = (event: any) => {
    setAge(event.target.value as string);
  };

  const BloodTypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBloodtypesID(event.target.value as number);
  };

  const DoctorhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDoctorID(event.target.value as number);
  };
  
  const [prefixID, setPredixID] = React.useState(Number);
  const [name, setName] = React.useState(String);
  const [genderID, setGenderID] = React.useState(Number);
  const [weight, setWeight] = React.useState(String);
  const [height, setHeight] = React.useState(String);
  const [age, setAge] = React.useState(String);
  const [bloodtypeID, setBloodtypesID] = React.useState(Number);
  const [doctorID, setDoctorID] = React.useState(Number);

  
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


  // function save data
  function save() {
    const patients = {
      age: age,
      height: height,
      name: name,
      weight: weight,
      prefix: prefixID,
      doctor: doctorID,
      gender: genderID,
      bloodtype: bloodtypeID,
    };
    
    
    console.log(patients)
   
    const apiUrl = 'http://localhost:8080/api/v1/patients';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(patients),
    };

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
      
      <Header style={HeaderCustom} title={`ระบบลงทะเบียนผู้ป่วยใน`}>
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
            <Grid item xs={3}>
              <div className={classes.paper}>คำนำหน้าชื่อ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <Select
                  id ="prefix_id"
                  name="prefix_id"
                  value={prefixID }
                  onChange={PrefixhandleChange}
                >
                  {prefixs.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.pname}
                      </MenuItem>
                    );
                  })}
                  
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อผู้ป่วย</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  
                  
                  value = {name || ''}
                  id ="patientName"
                  name="patientName"
                  variant="outlined" 
                  onChange={NamehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                  shrink: true,
                }}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เพศ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <Select
                  name="gender"
                  id="gender"
                  value={genderID}
                  onChange={GenderhandleChange}
                >
                  {genders.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.gname}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>น้ำหนัก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  name = "weight"
                  id="weight" 
                  value = {weight}
                  
                  variant="outlined"
                  onChange={WeighthandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                  shrink: true,
                }}
                />
                
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ส่วนสูง</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  name = "height"
                  id="height" 
                  value = {height}
                  variant="outlined"
                  onChange={HeighthandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                  shrink: true,
                }}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>อายุ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  name ="age"
                  id="age" 
                  value = {age}
                  variant="outlined"
                  onChange={AgehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                  shrink: true,
                }}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>กรุ๊ปเลือด</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <Select
                  name="bloodtype_id"
                  id="bloodtype_id"
                  value={bloodtypeID }
                  onChange={BloodTypehandleChange}
                >
                  {bloodtypes.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.bTname}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            
            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
            <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                Save
              </Button>
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              <Button
                variant="contained"
                color="amber"
                size="large"
                startIcon={<NavigateBeforeIcon />}
                component={RouterLink} to="/homenurse"
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

export default Register;
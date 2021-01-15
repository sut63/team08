import React, { FC, useEffect } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import LocalHospitalIcon from '@material-ui/icons/LocalHospital';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntNurse, EntDoctor, EntMedical } from '../../api/models'; // import interface User
import Swal from 'sweetalert2'; // alert
const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));


const Login: FC<{}> = () => {


  const classes = useStyles();
  const api = new DefaultApi();

  const [nurses, SetNurses] = React.useState<EntNurse[]>([]);
  const [doctors, SetDoctors] = React.useState<EntDoctor[]>([]);
  const [medicals, SetMedicals] = React.useState<EntMedical[]>([]);

  const getNurse = async () => {
    const res = await api.listNurse({ limit: 10, offset: 0 });
    SetNurses(res);
    console.log(res)
  }
  const getDoctor = async () => {
    const res = await api.listDoctor({ limit: 10, offset: 0 });
    SetDoctors(res);
    console.log(res)
  }
  const getMedicals = async () => {
    const res = await api.listMedical({ limit: 10, offset: 0 });
    SetMedicals(res);
    console.log(res)
  }



  const EmailhandleChange = (event: any) => {
    setEmail(event.target.value as string);
  }

  const PasswordhandleChange = (event: any) => {
    setPassword(event.target.value as string);
  }

  const [email, setEmail] = React.useState(String);
  const [password, setPassword] = React.useState(String);

  useEffect(() => {
    getNurse();
    getDoctor();
    getMedicals();
  }, []);

  var status = false
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
  function login() {
    console.log(email, password)
    nurses.filter(item => item.nurseEmail === email).map(item2 => {
      if (status == false) {
        if (item2.nursePassword == password) {
          window.location.href = "http://localhost:3000/homenurse";
          status = true
        }
      }
    })
    doctors.filter(item => item.doctorEmail === email).map(item2 => {
      if (status == false) {
        if (item2.doctorPassword == password) {
          window.location.href = "http://localhost:3000/homedoctor";
          status = true
        }
      }
    })
    medicals.filter(item => item.medicalEmail === email).map(item2 => {
      if (status == false) {
        if (item2.medicalPassword == password) {
          window.location.href = "http://localhost:3000/homemedical";
          status = true
        }
      }
    })
    if (status == false) {
      Toast.fire({
        icon: 'error',
        title: 'username หรือ password ไม่ถูกต้อง',
      });
    }
    status = false
  };


  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true" fontSize="large" />
          <LocalHospitalIcon aria-controls="fade-menu" aria-haspopup="true" fontSize="large" />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <TextField
          variant="outlined"
          margin="normal"
          required
          fullWidth
          id="email"
          label="Email Address"
          name="email"
          value={email}
          onChange={EmailhandleChange}
        />
        <TextField
          variant="outlined"
          margin="normal"
          required
          fullWidth
          name="password"
          label="Password"
          type="password"
          id="password"
          value={password}
          onChange={PasswordhandleChange}
        />
        <Button
          type="submit"
          fullWidth
          variant="contained"
          color="primary"
          onClick={login}


        >
          Sign In
          </Button>

      </div>

    </Container>
  );
};
export default Login;

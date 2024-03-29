import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save';
import { Link as RouterLink } from 'react-router-dom';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2'; // alert
import NavigateBeforeIcon from '@material-ui/icons/NavigateBefore';
import IconButton from '@material-ui/core/IconButton';
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
  EntRoom,
  EntPatient,
  EntRoomtype,
  EntNurse,
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


const Rent: FC<{ email: string }> = (email) => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [rooms, setRooms] = React.useState<EntRoom[]>([]);
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  const [roomtypes, setRoomtypes] = React.useState<EntRoomtype[]>([]);
  const [nurses, SetNurses] = React.useState<EntNurse[]>([]);
  const getNurse = async () => {
    const res = await http.listNurse({ limit: -1, offset: 0 });
    SetNurses(res);
    console.log(res)
  };

  const getRoom = async () => {
    const res = await http.listRoom({ limit: -1, offset: 0 });
    setRooms(res);
    console.log(res)
  };

  const getPatient = async () => {
    const res = await http.listPatient({ limit: 10, offset: 0 });
    setPatients(res);
  };

  const getRoomtype = async () => {
    const res = await http.listRoomtype({ limit: 3, offset: 0 });
    setRoomtypes(res);
  };
  // Lifecycle Hooks
  useEffect(() => {
    getRoom();
    getPatient();
    getRoomtype();
    getNurse();
  }, []);
  const RoomtypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRoomtypeID(event.target.value as number);
    console.log(roomtypeID)
  };
  const RoomhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRoomID(event.target.value as number);
    console.log()
  };
  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientID(event.target.value as number);
    console.log(roomtypeID)
  };
  const AddedhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAdded(event.target.value as string);
    console.log(roomtypeID)
  };
  const NursehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setNurseID(event.target.value as number);
    console.log(roomtypeID)
  };
  const RentIDhandleChange = (event: any) => {
    setRentID(event.target.value as string);
  };
  const KinNamehandleChange = (event: any) => {
    setKinName(event.target.value as string);
  };
  const KinTelhandleChange = (event: any) => {
    setKinTel(event.target.value as string);
  };
  const [roomtypeID, setRoomtypeID] = React.useState(Number);
  const [roomID, setRoomID] = React.useState(Number);
  const [patientID, setPatientID] = React.useState(Number);
  const [nurseID, setNurseID] = React.useState(Number);
  const [addeds, setAdded] = React.useState(String);
  const [rentID, setRentID] = React.useState(String);
  const [kinname, setKinName] = React.useState(String);
  const [kintel, setKinTel] = React.useState(String);

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

  const rents = {
    rentid: rentID,
    room: roomID,
    patient: patientID,
    kinname: kinname,
    kintel: kintel,
    nurse: nurseID,
    added: addeds + ":00+00:00",
  };
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
      case 'rent_id':
        aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: รหัสการจองขึ้นต้นด้วย R ตามด้วยตัวเลข 4 หลัก");
        return;
      case 'kin_tel':
        aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: รูปแบบของเบอร์โทรไม่ถูกต้อง");
        return;
      case 'kin_name':
        aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: รูปแบบชื่อไม่ถูกต้อง ตัวอักษรตัวแรกขึ้นต้นด้วยตัวพิมพ์ใหญ่");
        return;
      case "added time wrong":
        aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาป้อนวันที่ต้องการเข้าห้องพัก");
        return;
      case 'room not found':
        aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาป้อนห้องพักที่ต้องการเข้าห้องพัก");
        return;
      case 'patient not found':
        aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาป้อนชื่อผู้ป่วยที่ต้องการเข้าห้องพัก");
        return;
      case 'nurse not found':
        aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: กรุณาป้อนชื่อพยาบาลที่ประจำห้องพักนี้");
        return;
      default:
        aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ: ผู้ป่วยหรือห้องพักนี้ได้ทำการจองแล้ว ");
        return;
    };

  }
  // function save data
  function save() {
    console.log(rents)
    console.log(roomtypeID)
    const apiUrl = 'http://localhost:8080/api/v1/rents';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(rents),
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
          checkCaseSaveError(data.error.Name || data.error)
        }
      });
    rooms.map(item => {
      console.log(item.edges)
    })
    // function find room with roomtype

  }

  return (
    <Page theme={pageTheme.service}>

      <Header style={HeaderCustom} title={`ระบบจองห้องพักผู้ป่วย`}>
        <IconButton>
          <AccountCircleIcon />
        </IconButton>

        <div style={{ marginLeft: 10 }}> </div>
        <Link component={RouterLink} to="/">
          Logout
         </Link>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}>

            </Grid>

            <Grid item xs={3}></Grid>

            <Grid item xs={9}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>รหัสการจองห้องพัก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  name="rent_id"
                  id="rent_id"
                  value={rentID}
                  variant="outlined"
                  onChange={RentIDhandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />

              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ประเภทห้อง</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <Select
                  name="roomtype_id"
                  id="roomtype_id"
                  value={roomtypeID}
                  onChange={RoomtypehandleChange}
                >
                  {roomtypes.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>รายการห้องพัก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <Select
                  name="room_id"
                  id="room_id"
                  value={roomID}
                  onChange={RoomhandleChange}
                >
                  {rooms.filter(item => item.edges?.roomtype?.id === roomtypeID).map((item2: any) => {
                    return (

                      <MenuItem key={item2.id} value={item2.id}>
                        {item2.name}
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
                <Select
                  id="patient_id"
                  name="patient_id"

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
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อญาติผู้ป่วย</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  name="kin_name"
                  id="kin_name"
                  value={kinname}
                  variant="outlined"
                  onChange={KinNamehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />

              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เบอร์โทรญาติผู้ป่วย</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  name="kin_tel"
                  id="kin_tel"
                  value={kintel}
                  variant="outlined"
                  onChange={KinTelhandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />

              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>พยาบาล</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <Select
                  id="nurse_id"
                  name="nurse_id"
                  value={nurseID}
                  onChange={NursehandleChange}
                >
                  {nurses.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nurseName}
                      </MenuItem>
                    );
                  })}

                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เวลา</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกเวลา"
                  name="added"
                  type="datetime-local"
                  value={addeds || null} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={AddedhandleChange}
                />
              </form>
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

export default Rent;
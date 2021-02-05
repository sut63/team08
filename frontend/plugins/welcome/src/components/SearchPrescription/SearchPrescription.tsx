//style
import React, {  useState, useEffect } from 'react';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, TextField, Button,  Grid,Link } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Link as RouterLink } from 'react-router-dom';
//api
import { DefaultApi } from '../../api/apis';
//table
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
//entity
import { EntPrescription } from '../../api/models/EntPrescription';
//alert
import Swal from 'sweetalert2'
//icon
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import BackspaceTwoToneIcon from '@material-ui/icons/BackspaceTwoTone';
import AddCircleOutlineTwoToneIcon from '@material-ui/icons/AddCircleOutlineTwoTone';
// name
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useCookies } from 'react-cookie/cjs';//cookie
const cookies = new Cookies();
const Name = cookies.get('Name');
const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  root: {
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'center',
  },
  margin: {
    margin: theme.spacing(2),
  },
  insideLabel: {
   margin: theme.spacing(1),
  },
  button: {
   marginLeft: '40px',
  },
  textField: {
   width: 500 ,
   marginLeft:7,
   marginRight:-7,
  },
  select: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
  },
  paper: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(1),
  },
  pa: {
    marginTop: theme.spacing(2),
  },
  table: {
    minWidth: 650,
  },
  }),
);

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


export default function ComponentsTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  const [checkprescriptionnumber, setPrescriptionnumbers] = useState(false);
  const [prescription, setPrescription] = useState<EntPrescription[]>([])

  const [prescriptionnumber, setPrescriptionnumber] = useState(String);
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getPrescriptions = async () => {
      const res = await http.listPrescription({ offset: 0 });
      setLoading(false);
      setPrescription(res);
    };
    getPrescriptions();
  }, [loading]);

  const prescriptionnumberhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setPrescriptionnumbers(false);
    setPrescriptionnumber(event.target.value as string);

  };

  const deletePrescription = async (id: number) => {
    const res = await http.deletePrescription({ id: id });
    setLoading(true);
  };

  const checkresearch = async () => {
    var check = false;
    prescription.map(item => {
      if (prescriptionnumber != "") {
        if (item.prescripNumber?.includes(prescriptionnumber)) {
          setPrescriptionnumbers(true);
          alertMessage("success", "ค้นหาข้อมูลการสั่งยาสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ค้นหาข้อมูลการสั่งยาไม่สำเร็จ");
    }
    console.log(checkprescriptionnumber)
    if (prescriptionnumber == "") {
      alertMessage("info", "แสดงข้อมูลการสั่งยาทั้งหมดในระบบ");
    }
  };
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
      <Header style={HeaderCustom} title={`ระบบค้นหาข้อมูลการสั่งยาให้ผู้ป่วยใน`}>
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
        <ContentHeader title="">

        <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/pre" 
          variant="contained"
          color="primary"
          startIcon={<AddCircleOutlineTwoToneIcon/>}
          > 
          เพิ่มข้อมูล 
          </Button>
         
          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/homenurse" 
          variant="contained"
          color="default"
          startIcon={<BackspaceTwoToneIcon/>}
          > 
          ย้อนกลับ 
          </Button>
        </ContentHeader>
        
        <div className={classes.root}>
        <form noValidate autoComplete="on">
          <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
            size="small"
          >
            <TextField
            style={{ width: 250 ,marginLeft:7,marginRight:-7,marginTop:5}}
            className={classes.textField}
              id="prescriptionnumber"
              label="ค้นหาข้อมูลการสั่งยา"
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={prescriptionnumber}
              onChange={prescriptionnumberhandlehange}
            />
            </FormControl>

            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              <Button  
              onClick={() => {
                checkresearch();
                setSearch(true);
              }}
              variant="contained" 
              color="secondary" 
              startIcon={<SearchTwoToneIcon/>}
              > 
              ค้นหาข้อมูล 
              </Button>
            </form>
            </div>
        
        <Grid container justify="center">
          <Grid item xs={12} md={12}>
            <Paper>
              {search ? (
                <div>
                  {  checkprescriptionnumber ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                            <TableCell align="center">ลำดับที่</TableCell>
                            <TableCell align="center">รหัสการสั่งยา</TableCell>
                            <TableCell align="center">ผู้ป่วย</TableCell>
                            <TableCell align="center">แพทย์ผู้สั่งยา</TableCell>
                            <TableCell align="center">ยาที่สั่ง</TableCell>
                            <TableCell align="center">วันเวลาที่สั่งยา</TableCell>
                            <TableCell align="center">พยาบาลผู้บันทึก</TableCell>
                            <TableCell align="center">วิธีการใช้ยา</TableCell>
                            <TableCell align="center">หมายเหตุ</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {prescription.filter((filter: any) => filter.prescripNumber.includes(prescriptionnumber)).map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.prescripNumber}</TableCell>
                              <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                              <TableCell align="center">{item.edges?.doctor?.doctorName}</TableCell>
                              <TableCell align="center">{item.edges?.drug?.drugName}</TableCell>
                              <TableCell align="center">{moment(item.prescripDateTime).format('DD/MM/YYYY hh:mm')}</TableCell>
                              <TableCell align="center">{item.edges?.nurse?.nurseName}</TableCell>
                              <TableCell align="center">{item.prescripIssue}</TableCell>
                              <TableCell align="center">{item.prescripNote}</TableCell>
                              <Button
                                onClick={() => {
                                   if (item.id === undefined) {
                                    return;
                                    }
                                   deletePrescription(item.id);
                                    }}
                                   style={{ marginLeft: 10 }}
                                  variant="contained"
                                  color="secondary"
                                 >
                                  Delete
                                  </Button>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : prescriptionnumber == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                              <TableCell align="center">ลำดับที่</TableCell>
                            <TableCell align="center">รหัสการสั่งยา</TableCell>
                            <TableCell align="center">ผู้ป่วย</TableCell>
                            <TableCell align="center">แพทย์ผู้สั่งยา</TableCell>
                            <TableCell align="center">ยาที่สั่ง</TableCell>
                            <TableCell align="center">วันเวลาที่สั่งยา</TableCell>
                            <TableCell align="center">พยาบาลผู้บันทึก</TableCell>
                            <TableCell align="center">วิธีการใช้ยา</TableCell>
                            <TableCell align="center">หมายเหตุ</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {prescription.map((item: any) => (
                                <TableRow key={item.id}>
                                  <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.prescripNumber}</TableCell>
                              <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                              <TableCell align="center">{item.edges?.doctor?.doctorName}</TableCell>
                              <TableCell align="center">{item.edges?.drug?.drugName}</TableCell>
                              <TableCell align="center">{moment(item.prescripDateTime).format('DD/MM/YYYY hh:mm')}</TableCell>
                              <TableCell align="center">{item.edges?.nurse?.nurseName}</TableCell>
                              <TableCell align="center">{item.prescripIssue}</TableCell>
                              <TableCell align="center">{item.prescripNote}</TableCell>
                              <Button
                                onClick={() => {
                                   if (item.id === undefined) {
                                    return;
                                    }
                                   deletePrescription(item.id);
                                    }}
                                   style={{ marginLeft: 10 }}
                                    variant="contained"
                                  color="secondary"
                                 >
                                  Delete
                                </Button>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>

                      </div>
                    ) : null}
                </div>
              ) : null}
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}
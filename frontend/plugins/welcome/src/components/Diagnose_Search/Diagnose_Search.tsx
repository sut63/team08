//style
import React, {  useState, useEffect } from 'react';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, TextField, Button, Grid, Link } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
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
import { EntDiagnose } from '../../api/models/EntDiagnose';
//alert
import Swal from 'sweetalert2'
//icon
import AddCircleOutlineTwoToneIcon from '@material-ui/icons/AddCircleOutlineTwoTone';
import BackspaceTwoToneIcon from '@material-ui/icons/BackspaceTwoTone';
import CancelTwoToneIcon from '@material-ui/icons/CancelTwoTone';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import DeleteTwoToneIcon from '@material-ui/icons/DeleteTwoTone';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useCookies } from 'react-cookie/cjs';//cookie
import { Link as RouterLink } from 'react-router-dom';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';

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

  const [checkpatientname, setPatientnames] = useState(false);
  const [diagnose, setDiagnose] = useState<EntDiagnose[]>([])

  const [patientname, setPatientname] = useState(String);
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getDiagnoses = async () => {
      const res = await http.listDiagnose({ offset: 0 });
      setLoading(false);
      setDiagnose(res);
    };
    getDiagnoses();
  }, [loading]);

  const patientnamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setPatientnames(false);
    setPatientname(event.target.value as string);

  };

  const cleardata = () => {
    setPatientname("");
    setSearch(false);
    setPatientnames(false);
    setSearch(false);
  }

  const deleteDiagnose = async (id: number) => {
    const res = await http.deleteDiagnose({ id: id });
    setLoading(true);
  };

  const checkresearch = async () => {
    var check = false;
    diagnose.map(item => {
      if (patientname != "") {
        if (item.edges?.patient?.patientName?.includes(patientname)) {
          setPatientnames(true);
          alertMessage("success", "ค้นหาประวัติการวินิจฉัยสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบประวัติการวินิจฉัย");
    }
    console.log(checkpatientname)
    if (patientname == "") {
      alertMessage("info", "แสดงข้อมูลการวินิจฉัยโรคผู้ป่วยในทั้งหมดในระบบ");
    }
  };

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
      <Header style={HeaderCustom} title={`ระบบค้นหาประวัติการวินิจฉัยโรคผู้ป่วยใน`} >
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
        <ContentHeader title=" ">

        <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/dia" 
          variant="contained"
          color="primary"
          startIcon={<AddCircleOutlineTwoToneIcon/>}
          > 
          เพิ่มข้อมูล 
          </Button>

          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/homedoctor" 
          variant="contained"
          color="default"
          startIcon={<BackspaceTwoToneIcon/>}
          > 
          ย้อนกลับ 
          </Button>
        </ContentHeader>

        <div className={classes.root}>
        <form noValidate autoComplete="off">
          <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
            size="small"
          >
           
            <TextField
            style={{ width: 250 ,marginLeft:7,marginRight:-7,marginTop:5}}
            className={classes.textField}
              id="patientname"
              label = "ค้นหาชื่อผู้ป่วย..."
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={patientname}
              onChange={patientnamehandlehange}
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
                  {  checkpatientname ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                                <TableCell align="center">รหัสการวินิจฉัย</TableCell>
                                <TableCell align="center">แพทย์ที่รักษา</TableCell>
                                <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                                <TableCell align="center">โรค</TableCell>
                                <TableCell align="center">อาการเพิ่มเติม</TableCell>
                                <TableCell align="center">แผนกการเข้ารักษา</TableCell>
                                <TableCell align="center">หมายเหตุ</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {diagnose.filter((filter: any) => filter.edges?.patient?.patientName.includes(patientname)).map((item: any) => (
                            <TableRow key={item.id}>
                            <TableCell align="center">{item.diagnoseID}</TableCell>                          
                            <TableCell align="center">{item.edges?.doctor?.doctorName}</TableCell>
                            <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                            <TableCell align="center">{item.edges?.disease?.diseaseName}</TableCell>
                            <TableCell align="center">{item.diagnoseSymptoms}</TableCell>
                            <TableCell align="center">{item.edges?.department?.departmentName}</TableCell>
                            <TableCell align="center">{item.diagnoseNote}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : patientname == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                                <TableCell align="center">รหัสการวินิจฉัย</TableCell>
                                <TableCell align="center">แพทย์ที่รักษา</TableCell>
                                <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                                <TableCell align="center">โรค</TableCell>
                                <TableCell align="center">อาการเพิ่มเติม</TableCell>
                                <TableCell align="center">แผนกการเข้ารักษา</TableCell>
                                <TableCell align="center">หมายเหตุ</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {diagnose.map((item: any) => (
                                <TableRow key={item.id}>
                                <TableCell align="center">{item.diagnoseID}</TableCell>                          
                                <TableCell align="center">{item.edges?.doctor?.doctorName}</TableCell>
                                <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                                <TableCell align="center">{item.edges?.disease?.diseaseName}</TableCell>
                                <TableCell align="center">{item.diagnoseSymptoms}</TableCell>
                                <TableCell align="center">{item.edges?.department?.departmentName}</TableCell>
                                <TableCell align="center">{item.diagnoseNote}</TableCell>
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
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
//entity
import { EntCoveredPerson } from '../../api/models/EntCoveredPerson';
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

  const [checkpatientname, setPatientnames] = useState(false);
  const [coveredperson, setCoveredPerson] = useState<EntCoveredPerson[]>([])

  const [patientname, setPatientname] = useState(String);
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getCoveredPersons = async () => {
      const res = await http.listCoveredperson({ offset: 0 });
      setLoading(false);
      setCoveredPerson(res);
    };
    getCoveredPersons();
  }, [loading]);

  const patientnamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setPatientnames(false);
    setPatientname(event.target.value as string);

  };

  const deleteCoveredperson = async (id: number) => {
    const res = await http.deleteCoveredperson({ id: id });
    setLoading(true);
  };

  const checkresearch = async () => {
    var check = false;
    coveredperson.map(item => {
      if (patientname != "") {
        if (item.edges?.patient?.patientName?.includes(patientname)) {
          setPatientnames(true);
          alertMessage("success", "ค้นหาข้อมูลสิทธิการรักษาพยาบาลสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ค้นหาข้อมูลสิทธิการรักษาพยาบาลไม่สำเร็จ");
    }
    console.log(checkpatientname)
    if (patientname == "") {
      alertMessage("info", "แสดงข้อมูลสิทธิการรักษาพยาบาลทั้งหมดในระบบ");
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
        <ContentHeader title="">

          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/covered" 
          variant="contained"
          color="primary"
          startIcon={<AddCircleOutlineTwoToneIcon/>}
          > 
          เพิ่มข้อมูล 
          </Button>

        <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/homemedical" 
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
              id="patientname"
              label="ค้นหาสิทธิการรักษา..."
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
                          <TableCell align="center">ลำดับที่</TableCell>
                          <TableCell align="center">รหัสสิทธิการรักษา</TableCell>
                          <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                          <TableCell align="center">ประเภทสิทธิรักษาพยาบาล</TableCell>
                          <TableCell align="center">กองทุน</TableCell>
                          <TableCell align="center">คำร้องกองทุน</TableCell>
                          <TableCell align="center">ใบรับรองแพทย์</TableCell>
                          <TableCell align="center">หมายเหตุ</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {coveredperson.filter((filter: any) => filter.edges?.patient?.patientName.includes(patientname)).map((item: any) => (
                            <TableRow key={item.id}>
                             <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.coveredPersonNumber}</TableCell>
                              <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                              <TableCell align="center">{item.edges?.schemeType?.schemeTypeName}</TableCell>
                              <TableCell align="center">{item.edges?.fund?.fundName}</TableCell>
                              <TableCell align="center">{item.fundTitle}</TableCell>
                              <TableCell align="center">{item.edges?.certificate?.certificateName}</TableCell>
                              <TableCell align="center">{item.coveredPersonNote}</TableCell>
                              <Button
                                onClick={() => {
                                   if (item.id === undefined) {
                                    return;
                                    }
                                   deleteCoveredperson(item.id);
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
                    : patientname == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                              <TableCell align="center">ลำดับที่</TableCell>
                          <TableCell align="center">รหัสสิทธิการรักษา</TableCell>
                          <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                          <TableCell align="center">ประเภทสิทธิรักษาพยาบาล</TableCell>
                          <TableCell align="center">กองทุน</TableCell>
                          <TableCell align="center">คำร้องกองทุน</TableCell>
                          <TableCell align="center">ใบรับรองแพทย์</TableCell>
                          <TableCell align="center">หมายเหตุ</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {coveredperson.map((item: any) => (
                                <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.coveredPersonNumber}</TableCell>
                              <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                              <TableCell align="center">{item.edges?.schemeType?.schemeTypeName}</TableCell>
                              <TableCell align="center">{item.edges?.fund?.fundName}</TableCell>
                              <TableCell align="center">{item.fundTitle}</TableCell>
                              <TableCell align="center">{item.edges?.certificate?.certificateName}</TableCell>
                              <TableCell align="center">{item.coveredPersonNote}</TableCell>
                              <Button
                                onClick={() => {
                                   if (item.id === undefined) {
                                    return;
                                    }
                                   deleteCoveredperson(item.id);
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
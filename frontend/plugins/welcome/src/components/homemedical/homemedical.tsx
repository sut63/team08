import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import ApartmentIcon from '@material-ui/icons/Apartment';
import LoupeIcon from '@material-ui/icons/Loupe';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import HowToRegIcon from '@material-ui/icons/HowToReg';
import Typography from '@material-ui/core/Typography';
import {
  Container,
  Grid,
  Table,
} from '@material-ui/core';
import Link from '@material-ui/core/Link';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command

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
    width: 1000,
  },
}));

  
      function a11yProps(index: any) {
        return {
          id: `scrollable-force-tab-${index}`,
          'aria-controls': `scrollable-force-tabpanel-${index}`,
        };
      }

      export default function ScrollableTabsButtonForce() {
        const classes = useStyles();
        const [value, setValue] = React.useState(0);
      
        const handleChange = (event: React.ChangeEvent<{}>, newValue: number) => {
          setValue(newValue);
        };

  return (

    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`เจ้าหน้าที่เวชระเบียน`}>
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
  
        <Link component={RouterLink} to="/">
             Logout
         </Link>
      </Header>

    
    
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={1}>
            <Grid item xs={12}>
   <Table >
     <tr><td></td>
       <td  align="center">
           <br></br>
         <ApartmentIcon color="primary" style={{ fontSize: 150 }} />
     <Typography component="h1" variant="h5">
     SE63 Hospital
     </Typography><br></br>
     </td>
     </tr>

     <tr><td></td><td>
         <hr ></hr></td>
     </tr>
     <Grid item xs={5}>
              <div className={classes.paper}> </div>
            </Grid>
     <tr><br></br></tr>
     <tr><td></td>
     <td>
     <AppBar position="static" color="default">
        <Tabs
          value={value}
          onChange={handleChange}
          variant="scrollable"
          scrollButtons="on"
          indicatorColor="primary"
          textColor="primary"
          aria-label="scrollable force tabs example"
        >
           <Tab label="ระบบบันทึกสิทธิการรักษาพยาบาล" icon={<LoupeIcon />} {...a11yProps(0)}  component={RouterLink} to="/covered"/>
           <Tab label="ระบบค้นหาสิทธิการรักษาพยาบาล" icon={<HowToRegIcon />} {...a11yProps(4)}  component={RouterLink} to="/"/>
        </Tabs>
      </AppBar>
     </td>
     </tr>

     </Table>
     </Grid> 
           
           
            </Grid>
        </Container>
      </Content>
   
    </Page>
  );
  }
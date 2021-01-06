import React, { useEffect, useState } from 'react';
import { makeStyles, createStyles, Theme, useTheme } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import { Avatar, Button, Container } from '@material-ui/core';

import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';

import { Link as RouterLink } from 'react-router-dom';
import Typography from '@material-ui/core/Typography';

import { DefaultApi } from '../../api/apis';
import { EntPatient } from '../../api/models/EntPatient'
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
      paddingTop: 40,
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    details: {
      display: 'flex',
      flexDirection: 'column',
    },
    content: {

    },
    cover: {
      width: 150,
      display: 'flex',
      justify: 'flex-end',
      // paddingLeft: 60,
    },
    controls: {
      display: 'flex',
      alignItems: 'center',
      paddingLeft: theme.spacing(1),
      paddingBottom: theme.spacing(1),
    },
    RR: {
      display: 'flex',
      width: 400,
      height: 350,
      paddingRight: 0
    },
    ddd: {
      paddingRight: 70,
    },
    large: {
      width: theme.spacing(15),
      height: theme.spacing(15),
    },
  }),
);

export default function FullWidthGrid() {
  const [loading, setLoading] = useState(true);
  const classes = useStyles();

  const api = new DefaultApi();
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  const getPatients = async () => {
    const res = await api.listPatient({ limit: 10, offset: 0 });
    setLoading(false);
    setPatients(res);
  };
  // Lifecycle Hooks
  useEffect(() => {
    getPatients();
  }, [loading]);
  console.log(patients)
  return (
    <div className={classes.root}>
      <Container fixed >
        <Button variant="contained" color="secondary" component={RouterLink} to="/home">
          ผู้ป่วยรายใหม่
        </Button>
        
        <Grid container spacing={2}>
          {patients.map(key => {
            return (
              <Grid item className={classes.ddd}>
                <Card className={classes.RR}>
                  <Avatar variant="square" src="https://sv1.picz.in.th/images/2020/10/25/b3DaWW.png" className={classes.large} />
                  <div className={classes.details}>
                    <CardContent className={classes.content}>
                      <Typography component="h5" variant="h5">
                        ข้อมูลผู้ป่วย
                    </Typography>
                      <p>
                        เลขบัตรประชาชน : {key.cardId}
                      </p>
                      <p>
                        ชื่อ : {key.edges?.Title?.titleType}  {key.firstName}  {key.lastName}
                      </p>
                      <p>
                        เพศ : {key.edges?.Gender?.genderType}
                      </p>
                      <p>
                        อายุ : {key.age} ปี
                      </p>
                      <p>
                        วันเกิด : {key.birthday.substring(0, 10)}
                      </p>

                      <p>
                        อาชีพ : {key.edges?.Job?.jobName}
                      </p>
                      <p>
                        ที่อยู่ : {key.address}
                      </p>
                      <p>
                        อาการแพ้ยา : {key.allergic}
                      </p>
                    </CardContent>
                  </div>
                </Card>
              </Grid>
            )
          })}
        </Grid>
      </Container>
    </div>
  );
}
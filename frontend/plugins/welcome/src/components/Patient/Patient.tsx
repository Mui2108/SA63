import React, { FC, useEffect, useState } from 'react';
import { EntGender } from '../../api/models/EntGender';
import { EntTitle } from '../../api/models/EntTitle';
import { EntJob } from '../../api/models/EntJob';
import { Link as RouterLink } from 'react-router-dom';
import Swal from 'sweetalert2';

import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Select from '@material-ui/core/Select';
import FormControl from '@material-ui/core/FormControl';

import { DefaultApi } from '../../api/apis';

import SaveIcon from '@material-ui/icons/Save';
import BackIcon from '@material-ui/icons/ArrowBack';
import Box from '@material-ui/core/Box';
import {
  InputLabel,
  MenuItem,
} from '@material-ui/core';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      '& > *': {
        margin: theme.spacing(1),
        justifyContent: 'center',
        top: 75,
      },
      paddingTop: 95,
    },
    RootDiv: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    formControl: {
      width: 150,
    },
    textField: {
      width: 296,
    },
    textBox: {
      width: 200,
    },
    headTitle: {
      paddingLeft: 90,
    },
    idcard: {
      width: 368,
    },
    job: {
      display: '',
      paddingTop: 200,
    },
    formControl2: {
      width: 200,
    },
    textField2: {
      width: 200,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    textBox2: {
      width: 365,
    },
    button: {
      marginLeft: 375,
    },
    button2: {
      marginLeft: 0,
    },
  }),

);

interface patient {
  Card_id: string;
  First_name: string;
  Last_name: string;
  Allergic: string;
  Address: string;
  Age: number;
  Birthday: string;
  Gender: number;
  Title: number;
  Job: number;

}

const Patient: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [titles, setTitles] = React.useState<EntTitle[]>([]);
  const [jobs, setJobs] = React.useState<EntJob[]>([]);


  const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const [patient, setPatient] = React.useState<
    Partial<patient>>({});

  //console.log(patient)

  const getGenders = async () => {
    const res = await api.listGender({ limit: 10, offset: 0 });
    setGenders(res);
  };
  const getTitles = async () => {
    const res = await api.listTitle({ limit: 10, offset: 0 });
    setTitles(res);
  };
  const getJobs = async () => {
    const res = await api.listJob({ limit: 10, offset: 0 });
    setJobs(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getGenders();
    getTitles();
    getJobs();
  }, []);
  // set data to object pat
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Patient;
    const { value } = event.target;
    setPatient({ ...patient, [name]: value });
  };
  const handleChange2 = (
    event: React.ChangeEvent<{ name: string; value: number }>,
  ) => {
    const name = event.target.name as keyof typeof Patient;
    const { value } = event.target;
    setPatient({ ...patient, [name]: +value });
  };
  function clear() {
    setPatient({});
  }
  function checkNull() {
    if (patient.Age === null || patient.Allergic === null || patient.Birthday === null || patient.Card_id === null || patient.First_name === null || patient.Gender === null || patient.Job == null || patient.Last_name === null || patient.Title === null) {
      Toast.fire({
        icon: 'warning',
        title: 'โปรดระบุข้อมูลให้ครบ',
      });
    } else {
      save()
    }
  }

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/patients';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(patient),
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
          clear()
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });

        }
      });

  };

  return (
    <div className={classes.RootDiv}>
      <Box
        bgcolor="grey.700"
        color="white"
        width={350}
        height={90}
        position="absolute"
        top={90}
        left="ceter"
        zIndex="tooltip"

      >
        <h1 className={classes.headTitle}>ประวัติผู้ป่วย</h1>
      </Box>
      <form className={classes.root} noValidate autoComplete="off">
        <TextField
          className={classes.idcard}
          label="เลขบัตประชาชน"
          variant="outlined"
          name="Card_id"
          value={patient.Card_id || ''}
          onChange={handleChange}
        ></TextField>
        <FormControl variant="outlined" className={classes.formControl2}>
          <InputLabel id="name">อาชีพ</InputLabel>
          <Select
            label="อาชีพ"
            name="Job"
            onChange={handleChange}
            value={patient.Job || ''}
          >
            {jobs.map(item => {
              return (
                <MenuItem key={item.id} value={item.id}>
                  {item.jobName}
                </MenuItem>);
            })}
          </Select>
        </FormControl>
        <br />
        <FormControl variant="outlined" className={classes.formControl}>
          <InputLabel id="name">คำนำหน้า</InputLabel>
          <Select
            label="คำนำหน้า"
            name="Title"
            onChange={handleChange}
            value={patient.Title || ''}
          >
            {titles.map(item => {
              return (
                <MenuItem key={item.id} value={item.id}>
                  {item.titleType}
                </MenuItem>);
            })}

          </Select>
        </FormControl>

        <TextField
          className={classes.textBox}
          name="First_name"
          label="ชื่อ"
          variant="outlined"
          value={patient.First_name || ''}
          onChange={handleChange}

        ></TextField>

        <TextField
          className={classes.textBox}
          name="Last_name"
          label="นามสกุล"
          variant="outlined"
          value={patient.Last_name || ''}
          onChange={handleChange}
        ></TextField>
        <br />

        <FormControl variant="outlined" className={classes.formControl}>
          <InputLabel id="name">เพศ</InputLabel>
          <Select
            label="เพศ"
            name="Gender"
            value={patient.Gender || ''}
            onChange={handleChange}
          >
            {genders.map(item => {
              return (
                <MenuItem key={item.id} value={item.id}>
                  {item.genderType}
                </MenuItem>);
            })}

          </Select>
        </FormControl>
        <TextField
          className={classes.textBox}
          label="อายุ"
          variant="outlined"
          name="Age"
          type="number"
          value={patient.Age || ''}
          onChange={handleChange2}
        ></TextField>

        <TextField
          name="Birthday"
          variant="outlined"
          label="วันเกิด"
          type="date"
          defaultValue="2017-05-24"
          className={classes.textField2}
          value={patient.Birthday || ''}
          onChange={handleChange}
          InputLabelProps={{
            shrink: true,
          }}
        ></TextField>
        <br />
        <TextField
          className={classes.textBox}
          name="Address"
          value={patient.Address || ''}
          label="ที่อยู่"
          variant="outlined"
          onChange={handleChange}
        ></TextField>
        <TextField
          className={classes.textBox2}
          name="Allergic"
          label="ประวัติการแพ้ยา"
          multiline
          value={patient.Allergic || ''}
          rows={4}
          variant="outlined"
          onChange={handleChange}
        />
        <br />
        <Button
          variant="contained"
          color="primary"
          size="large"
          className={classes.button}
          startIcon={<SaveIcon />}
          onClick={checkNull}
        >
          บันทึก
        </Button>
        <Button
          variant="contained"
          size="large"
          className={classes.button2}
          startIcon={<BackIcon />}
          component={RouterLink} to="/"

        >
          กลับ
        </Button>
      </form>
    </div>
  );
};

export default Patient;

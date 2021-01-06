import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Patient from './components/Patient'
import TPatient from './components/TablePatient';
import Login from './components/Login';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/home', WelcomePage);
    router.registerRoute('/patient', Patient)
    router.registerRoute('/tpatient', TPatient)
   
  },
});

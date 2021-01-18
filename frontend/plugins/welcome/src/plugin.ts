import { createPlugin } from '@backstage/core';
import Regeister from './components/Patient_Register';
import Rent from './components/Rent_Room';
import SignIn from './components/SignIn'
import Covered from './components/CoveredPerson'
import homedoctor from './components/homedoctor'
import homenurse from './components/homenurse'
import homemedical from './components/homemedical'
import Prescription from './components/Prescription'
import Diagnose from './components/Diagnose'
import Operativerecord from './components/Operativerecord'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
    router.registerRoute('/homedoctor', homedoctor);
    router.registerRoute('/homenurse', homenurse);
    router.registerRoute('/homemedical', homemedical);
    router.registerRoute('/reg', Regeister);
    router.registerRoute('/rent', Rent);
    router.registerRoute('/covered', Covered);
    router.registerRoute('/pre', Prescription);
    router.registerRoute('/dia', Diagnose);
    router.registerRoute('/opera', Operativerecord);
    
    
  },
});

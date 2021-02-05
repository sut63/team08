import { createPlugin } from '@backstage/core';
import Regeister from './components/Patient_Register';
import Rent from './components/Rent_Room';
import SignIn from './components/SignIn'
import Covered from './components/CoveredPerson'
import Coveredsearch from './components/CoveredPersonsearch'
import homedoctor from './components/homedoctor'
import homenurse from './components/homenurse'
import homemedical from './components/homemedical'
import Prescription from './components/Prescription'
import SearchPrescription from './components/SearchPrescription';
import Diagnose from './components/Diagnose'
import Operativerecord from './components/Operativerecord'
import searchRent from './components/Rent_Roomsearch'

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
    router.registerRoute('/coveredsearch', Coveredsearch);
    router.registerRoute('/pre', Prescription);
    router.registerRoute('/searchpre', SearchPrescription);
    router.registerRoute('/dia', Diagnose);
    router.registerRoute('/opera', Operativerecord);
    router.registerRoute('/rentsearch', searchRent);
    
  },
});

import {createApp} from 'vue';
import App from './App.vue';
import {LoadingBar, Notify, Quasar} from 'quasar';
import '@quasar/extras/material-icons/material-icons.css';
import 'quasar/src/css/index.sass';
import router from './router';

const app = createApp(App);
app
  .use(Quasar, {
    plugins: { LoadingBar, Notify }, // import Quasar plugins and add here
    config: { loadingBar: {} },
  })
  .use(router)
app.mount('#app');

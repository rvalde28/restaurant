import { createApp } from 'vue';
import App from './App.vue';
import { provideEventBus } from './EventBus'; // Adjust path as per your project structure
// import '../src/assets/main.css';
import './index.css'
import router from './router';
import { globalState } from './globalstate.js';

const app = createApp(App);

// Use provideEventBus to provide EventBus to the entire Vue app
app.use(provideEventBus);

app.use(router);

app.config.globalProperties.$globalState = globalState;

app.mount('#app');
